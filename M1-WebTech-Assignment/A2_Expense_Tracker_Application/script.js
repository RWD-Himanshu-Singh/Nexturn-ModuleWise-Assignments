const expenseForm = document.getElementById('expense-form');
const expenseTable = document.getElementById('expense-table');
const categorySummary = document.getElementById('category-summary');
let expenses = JSON.parse(localStorage.getItem('expenses')) || [];

// Add expense
expenseForm.addEventListener('submit', (e) => {
    e.preventDefault();

    // Get values from the form
    const amount = parseFloat(document.getElementById('amount').value);
    const description = document.getElementById('description').value;
    const category = document.getElementById('category').value;

    if (!amount || !description || !category) {
        alert('Please fill out all fields!');
        return;
    }

    // Create an expense object
    const expense = { amount, description, category };
    expenses.push(expense);

    // Save to localStorage
    localStorage.setItem('expenses', JSON.stringify(expenses));

    // Update UI
    renderExpenses();
    renderSummary();

    // Reset form
    expenseForm.reset();
});

// Render expenses in the table
function renderExpenses() {
    expenseTable.innerHTML = '';
    expenses.forEach((expense, index) => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${expense.amount}</td>
            <td>${expense.description}</td>
            <td>${expense.category}</td>
            <td><button onclick="deleteExpense(${index})">Delete</button></td>
        `;
        expenseTable.appendChild(row);
    });
}

// Delete an expense
function deleteExpense(index) {
    expenses.splice(index, 1);
    localStorage.setItem('expenses', JSON.stringify(expenses));
    renderExpenses();
    renderSummary();
}

// Render category summary
function renderSummary() {
    const summary = {};
    expenses.forEach((expense) => {
        summary[expense.category] = (summary[expense.category] || 0) + expense.amount;
    });

    categorySummary.innerHTML = '';
    for (const [category, total] of Object.entries(summary)) {
        const li = document.createElement('li');
        li.textContent = `${category}: ${total}`;
        categorySummary.appendChild(li);
    }
}

// Initial render
renderExpenses();
renderSummary();
