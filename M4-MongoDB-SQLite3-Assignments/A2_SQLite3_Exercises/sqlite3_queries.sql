
-- Q1. List names of employees hired after January 1, 2021
SELECT Name FROM Employees WHERE HireDate > '2021-01-01';

-- Q2. Calculate average salary of employees in each department
SELECT DepartmentID, AVG(Salary) AS AverageSalary FROM Employees GROUP BY DepartmentID;

-- Q3. Find the department name with the highest total salary
SELECT d.DepartmentName FROM Departments d
JOIN Employees e ON d.DepartmentID = e.DepartmentID
GROUP BY d.DepartmentID
ORDER BY SUM(e.Salary) DESC
LIMIT 1;

-- Q4. List all departments with no employees assigned
SELECT DepartmentName FROM Departments d
LEFT JOIN Employees e ON d.DepartmentID = e.DepartmentID
WHERE e.EmployeeID IS NULL;

-- Q5. Fetch all employee details with their department names
SELECT e.*, d.DepartmentName FROM Employees e
JOIN Departments d ON e.DepartmentID = d.DepartmentID;
