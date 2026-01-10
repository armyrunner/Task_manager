-- Users Table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

-- Task Table
CREATE TABLE IF NOT EXISTS initial_tasks(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    task_description TEXT,
    due_date TEXT,
    start_date TEXT,
    finish_date TEXT,
    status TEXT,
    notes TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Completed Task Table				
CREATE TABLE IF NOT EXISTS completed_tasks(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    task_id INTEGER,
    task_description TEXT,
    due_date TEXT,
    start_date TEXT,
    finish_date TEXT,
    status TEXT,
    notes TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Refresh Tokens Table (for JWT)
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    token TEXT UNIQUE NOT NULL,
    expires_at TEXT NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Column migrations are handled in db.go runMigrations()
-- ALTER TABLE initial_tasks ADD COLUMN category TEXT;
-- ALTER TABLE completed_tasks ADD COLUMN category TEXT;
-- ALTER TABLE initial_tasks ADD COLUMN user_id INTEGER;
-- ALTER TABLE completed_tasks ADD COLUMN user_id INTEGER;