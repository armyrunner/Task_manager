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

    def add_task(self,description,due_date,start_date=None,finish_date=None,status='Not Started'):
        data = [description,due_date,start_date,finish_date,status]
        self.cursor.execute("INSERT INTO initial_tasks(description,due_date,start_date,finish_date,status) Values(?,?,?,?,?)",data)
        self.conn.commit() 
        return self.cursor.lastrowid       

    def get_all_tasks(self):
        self.cursor.execute("SELECT * FROM initial_tasks")
        return self.cursor.fetchall()
    
    def update_tasks(self,description,due_date,start_date,finish_date,status,task_id):
        data = [description,due_date,start_date,finish_date,status,task_id]
        self.cursor.execute("UPDATE initial_tasks SET description = ?,due_date = ?,start_date = ?,finish_date = ?,status = ? WHERE id = ?",data)
        self.conn.commit()

    def delete_task(self,task_id):
        self.cursor.execute("DELETE FROM initial_tasks WHERE id = ?",task_id)
        self.conn.commit()
        
