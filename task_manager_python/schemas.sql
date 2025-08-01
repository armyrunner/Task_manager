-- Task Table
CREATE TABLE IF NOT EXISTS Tasks(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT,
    start_date TEXT,
    due_date TEXT,
    finsih_date TEXT,
    status TEXT
    );

-- Completed Task Table				
CREATE TABLE IF NOT EXISTS CompletedTasks(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id INTEGER NOT NULL,
    description TEXT,
    start_date TEXT,
    due_date TEXT,
    finsih_date TEXT,
    status TEXT,
    FOREIGN KEY(task_id) REFERENCES Tasks(id)
    );