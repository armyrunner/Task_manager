-- Task Table
CREATE TABLE IF NOT EXISTS initial_tasks(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_description TEXT,
    due_date TEXT,
    start_date TEXT,
    finish_date TEXT,
    status TEXT,
    notes TEXT
    );

-- Completed Task Table				
CREATE TABLE IF NOT EXISTS completed_tasks(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id INTEGER,
    task_description TEXT,
    due_date TEXT,
    start_date TEXT,
    finish_date TEXT,
    status TEXT,
    notes TEXT
    );