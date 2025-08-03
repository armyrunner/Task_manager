import sqlite3

def dict_factory(cursor,row):
    d = {}
    for idx, col in enumerate(cursor.description):
        d[col[0]] = row[idx]
    return d

class Task_DB():

    def __init__(self):
        self.sql_file = "./schemas.sql"
        with open(self.sql_file,'r') as f:
            sql_script = f.read()

        self.conn = sqlite3.connect("task.db")
        self.conn.row_factory = dict_factory

        self.cursor = self.conn.cursor()
        self.cursor.executescript(sql_script)

    def add_task(self,description,due_date,start_date,finish_date,status):
        date = [description,due_date,start_date,finish_date,status]
        self.cursor.execute("INSERT INTO initial_tasks(description,due_date,start_date,finish_date,status) Values(?,?,?,?,?)")
        self.conn.commit()        




        
