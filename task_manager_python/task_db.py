import sqlite3

def dict_factory(cursor,row):
    d = {}
    for idx, col in enumerate(cursor.description):
        d[col[0]] = row[idx]
    return d

class Task_DB():

    def __init__(self,db_filename):
        self.sql_file = "./schemas.sql"
        with open(self.sql_file,'r') as f:
            sql_script = f.read()

        self.conn = sqlite3.connect(db_filename)
        self.conn.row_factory = dict_factory

        self.cursor = self.conn.cursor()
        self.cursor.executescript(sql_script)

    def close_db(self):
        if self.conn:
            self.conn.close()

    def add_task(self,task_description,due_date,start_date=None,finish_date=None,status='Not Started',notes=None):
        data = [task_description,due_date,start_date,finish_date,status,notes]
        self.cursor.execute("INSERT INTO initial_tasks(task_description,due_date,start_date,finish_date,status,notes) Values(?,?,?,?,?,?)",data)
        self.conn.commit() 
        return self.cursor.lastrowid       

    def get_all_tasks(self):
        self.cursor.execute("SELECT * FROM initial_tasks")
        return self.cursor.fetchall()
    
    def get_one_task(self,task_id):
        self.cursor.execute("SELECT * FROM initial_tasks WHERE id = ?", (task_id,))
        return self.cursor.fetchone()
        
    def get_task_id(self,task_description):
        self.cursor.execute("SELECT id FROM initial_tasks WHERE task_description = ?",(task_description,))
        return self.cursor.fetchone()
    
    def update_tasks(self,task_description,due_date,start_date,finish_date,status,notes,task_id):
        data = [task_description,due_date,start_date,finish_date,status,notes,task_id]
        self.cursor.execute("UPDATE initial_tasks SET task_description = ?,due_date = ?,start_date = ?,finish_date = ?,status = ?,notes = ? WHERE id = ?",data)
        self.conn.commit()

    def delete_task(self, task_id):
        try:
            self.cursor.execute("DELETE FROM initial_tasks WHERE id = ?", (task_id,))
            self.conn.commit()
            return True  # success
        except sqlite3.Error as e:
            print(f"SQLite error during delete: {e}")
            return False

    def get_completed_tasks(self):
        self.cursor.execute("SELECT * FROM completed_tasks")
        return self.cursor.fetchall()

    def completed_tasks(self,task_info,task_id):
        data = [task_id,task_info['task_description'],task_info['due_date'],task_info['start_date'],task_info['finish_date'],task_info['status'],task_info['notes']]
        self.cursor.execute("INSERT INTO completed_tasks(task_id,task_description,due_date,start_date,finish_date,status,notes) Values(?,?,?,?,?,?,?)",data)
        self.conn.commit()
        return self.cursor.lastrowid