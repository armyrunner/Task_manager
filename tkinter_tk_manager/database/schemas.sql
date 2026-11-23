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

-- Column migrations are handled in db.go runMigrations()
-- ALTER TABLE initial_tasks ADD COLUMN category TEXT;
-- ALTER TABLE completed_tasks ADD COLUMN category TEXT;
