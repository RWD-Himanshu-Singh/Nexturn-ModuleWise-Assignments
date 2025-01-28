
CREATE TABLE Employees (
    EmployeeID INTEGER PRIMARY KEY,
    Name TEXT NOT NULL,
    HireDate DATE NOT NULL,
    Salary REAL NOT NULL,
    DepartmentID INTEGER,
    FOREIGN KEY (DepartmentID) REFERENCES Departments(DepartmentID)
);

CREATE TABLE Departments (
    DepartmentID INTEGER PRIMARY KEY,
    DepartmentName TEXT NOT NULL
);

-- Sample data for testing
INSERT INTO Departments VALUES (1, 'HR'), (2, 'Engineering'), (3, 'Marketing');
INSERT INTO Employees VALUES 
(1, 'Alice', '2022-03-01', 70000, 2),
(2, 'Bob', '2020-06-15', 50000, 1),
(3, 'Charlie', '2023-01-20', 60000, 3),
(4, 'Diana', '2021-12-01', 80000, 2);
