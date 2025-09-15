<script>
  import { onMount } from 'svelte';
  import { GetTasks, AddTask, UpdateTask, DeleteTask } from '../wailsjs/go/main/App';

  let tasks = [];
  let newTask = '';
  let priority = 'medium';
  let dueDate = '';

  onMount(async () => {
  const fetchedTasks = await GetTasks();
  tasks = Array.isArray(fetchedTasks) ? fetchedTasks : [];
});


  async function addTask() {

    const task = {
      text: newTask,
      completed: false,
      priority: priority,
      due_date: dueDate,
    };

    const id = await AddTask(task);
    task.id = id; 

    tasks = [...tasks, task]; 

    newTask = '';
    priority = 'medium';
    dueDate = '';
  }

  async function toggleTaskCompletion(index) {
    tasks[index].completed = !tasks[index].completed;
    await UpdateTask(tasks[index]);
    tasks = [...tasks];
  }

  async function deleteTask(index) {
    const id = tasks[index].id;
    await DeleteTask(id);
    tasks.splice(index, 1);
    tasks = [...tasks];
  }
</script>

<div>
  <h1>To-Do List</h1>

  <input
    type="text"
    placeholder="Enter a new task"
    bind:value={newTask}
  />
  <input type="date" bind:value={dueDate} />
  <select bind:value={priority}>
    <option value="low">Low</option>
    <option value="medium">Medium</option>
    <option value="high">High</option>
  </select>
  <button on:click={addTask}>Add Task</button>

  <div>
    {#each tasks as task, index}
      <div class="task {task.completed ? 'completed' : ''}">
        <div>
          <input
            type="checkbox"
            checked={task.completed}
            on:change={() => toggleTaskCompletion(index)}
          />
          <span class="priority-{task.priority}">
            {task.text}
            {#if task.due_date}
              (due: {task.due_date})
            {/if}
          </span>
        </div>
        <button on:click={() => deleteTask(index)}>Delete</button>
      </div>
    {/each}
  </div>
</div>