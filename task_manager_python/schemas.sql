-- Task Table
CREATE TABLE IF NOT EXISTS initial_tasks(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT,
    due_date TEXT,
    start_date TEXT,
    finsih_date TEXT,
    status TEXT
    );

-- Completed Task Table				
CREATE TABLE IF NOT EXISTS completed_tasks(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id INTEGER NOT NULL,
    description TEXT,
    due_date TEXT,
    start_date TEXT,
    finsih_date TEXT,
    status TEXT,
    FOREIGN KEY(task_id) REFERENCES Tasks(id)
    );