// Drag & Drop Handlers
let draggedItem = null;

function handleDragStart(event) {
  draggedItem = this;
  setTimeout(() => (this.style.display = "none"), 0);
}

// Allow dropping
function handleDragOver(event) {
  event.preventDefault();
}

function handleDrop(event) {
  event.preventDefault();

  // Checking if the drop target is within the task list
  if (this.parentElement && this.parentElement.id === "taskList") {
    this.parentElement.insertBefore(draggedItem, this);
  } else if (this.id === "taskList") {
    this.appendChild(draggedItem);
  }

  draggedItem.style.display = "flex";
  draggedItem = null;

  // Saving new tasks to localStorage
  saveTasks();
}

// Ensuring that the dragged item always reappears even if dropped outside
function handleDragEnd() {
  if (draggedItem) {
    draggedItem.style.display = "flex";
    draggedItem = null;
  }
}

// Create task element with drag & drop and click listeners
function createTaskElement(taskText, isCompleted) {
  let newTask = document.createElement("li");
  newTask.setAttribute("draggable", "true");

  newTask.innerHTML = `
        <span class="task-name">${taskText}</span>
        <button class="update-btn">Update</button>
        <button class="done-btn">Done</button>
        <button class="remove-btn" onclick="removeTask(this)">Remove</button>
    `;

  if (isCompleted) {
    newTask.classList.add("completed");
    newTask.querySelector(".done-btn").textContent = "Undo"; // Change "Done" to "Undo" for completed tasks
  }

  // Toggle completed state on Done button click
  newTask.querySelector(".done-btn").addEventListener("click", () => {
    newTask.classList.toggle("completed");
    let doneButton = newTask.querySelector(".done-btn");
    if (newTask.classList.contains("completed")) {
      doneButton.textContent = "Undo";
      newTask.querySelector(".task-name").style.textDecoration = "line-through";
    } else {
      doneButton.textContent = "Done";
      newTask.querySelector(".task-name").style.textDecoration = "none";
    }
    saveTasks();
    updatePendingCount();
  });

  // Makes task text editable on Update button click
  newTask.querySelector(".update-btn").addEventListener("click", () => {
    let taskName = newTask.querySelector(".task-name");
    taskName.setAttribute("contenteditable", "true");
    taskName.focus();
    newTask.querySelector(".update-btn").textContent = "Save";

    taskName.addEventListener("blur", () => saveEditedTask(newTask));
    newTask
      .querySelector(".update-btn")
      .addEventListener("click", () => saveEditedTask(newTask));
  });

  // Attaching drag & drop handlers
  newTask.addEventListener("dragstart", handleDragStart);
  newTask.addEventListener("dragover", handleDragOver);
  newTask.addEventListener("drop", handleDrop);

  return newTask;
}

// Saving the edited task
function saveEditedTask(taskElement) {
  let taskName = taskElement.querySelector(".task-name");
  taskName.removeAttribute("contenteditable");
  taskElement.querySelector(".update-btn").textContent = "Update";
  saveTasks();
}

// Adding a new task
function addTask() {
  let taskInput = document.getElementById("taskInput");
  let taskText = taskInput.value.trim();
  if (taskText) {
    let taskList = document.getElementById("taskList");
    let newTask = createTaskElement(taskText, false);
    taskList.appendChild(newTask);
    taskInput.value = "";
    saveTasks();
    updatePendingCount();
  }
}

// Removing a task
function removeTask(taskButton) {
  let taskToRemove = taskButton.parentElement;
  taskToRemove.remove();
  saveTasks();
  updatePendingCount();
}

// Saving tasks to localStorage
function saveTasks() {
  let taskList = document.getElementById("taskList");
  let tasks = [];
  taskList.querySelectorAll("li").forEach((task) => {
    tasks.push({
      text: task.querySelector(".task-name").textContent,
      isCompleted: task.classList.contains("completed"),
    });
  });
  localStorage.setItem("tasks", JSON.stringify(tasks));
}

// Loading tasks from localStorage
function loadTasks() {
  let tasks = JSON.parse(localStorage.getItem("tasks")) || [];
  let taskList = document.getElementById("taskList");
  tasks.forEach((task) => {
    let taskElement = createTaskElement(task.text, task.isCompleted);
    taskList.appendChild(taskElement);
  });
  updatePendingCount();
}

// Updating the count of pending tasks
function updatePendingCount() {
  let taskList = document.getElementById("taskList");
  let pendingCount = taskList.querySelectorAll("li:not(.completed)").length;
  document.getElementById(
    "pendingCount"
  ).textContent = `Remaining tasks : ${pendingCount}`;
}

// Initialize app
document.addEventListener("DOMContentLoaded", function () {
  loadTasks();
  document.getElementById("addTaskBtn").addEventListener("click", addTask);
  document.addEventListener("dragend", handleDragEnd);
});
