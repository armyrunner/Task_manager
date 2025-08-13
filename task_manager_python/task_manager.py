
from prettytable import PrettyTable
from task import Task
from pdf_generator import PDF_Generator
from task_db import *

class Task_Manager:

    # init funcion
    def __init__(self,db_filename): 
        self.tdb = Task_DB(db_filename)
        self.task_report = "task_report.pdf"
        self.tasks = self.load_tasks()
    
        # load task function
    def load_tasks(self):
        all_tasks = self.tdb.get_all_tasks()
        tasks = []
        for row in all_tasks:
            tasks.append(Task(
                task_id=row['id'],
                task_description=row['task_description'],
                due_date=row['due_date'],
                start_date=row['start_date'],
                finish_date=row['finish_date'],
                status=row['status'],
                notes=row['notes']
            ))
        return tasks
    
    # save task function
    def save_tasks(self):
        # Might apply later for larger program. reserved for bulk saving, more fore GUI
        pass
    
    # add task function
    def add_task(self,task_description="",due_date = "",start_date=None,finish_date=None,status ='Not Started',notes=None):
        task_id = self.tdb.add_task(task_description,due_date,start_date,finish_date,status,notes)
        self.tasks = self.load_tasks()
        return task_id     


    # update task function
    def update_task(self,task_id,task_description =None,due_date = None,start_date=None,finish_date=None,status ='Not Started',notes=None):
        exists = self.tdb.get_one_task(task_id)

        if not exists:
            return None
        
        task_info = {
            'task_description': task_description if task_description not in [None,''] else exists['task_description'],
            'due_date': due_date if due_date not in [None,'']  else exists['due_date'],
            'start_date': start_date if start_date not in [None,'']  else exists['start_date'],
            'finish_date': finish_date if finish_date not in [None,'']  else exists['finish_date'],
            'status': status if status not in [None,'']  else exists['status'],
            'notes':notes if notes not in [None,'']  else exists['notes']
        }
        
        if task_info['status'].lower() in ['complete', 'completed']:
            self.tdb.completed_tasks(task_info)
            self.tdb.delete_task(task_id)
        else:
            self.tdb.update_tasks(
                task_info['task_description'],
                task_info['due_date'],
                task_info['start_date'],
                task_info['finish_date'],
                task_info['status'],
                task_info['notes'],
                task_id
            )
            
            msg = "Update Successful"
        
        self.tasks = self.load_tasks()
        return msg

    # delete task function
    def delete_task(self,task_id):
        removed = self.tdb.delete_task(task_id)
        if removed:
            self.tasks = [task for task in self.tasks if task.task_id != task_id]

    def display_tasks(self):
        table = PrettyTable()
        table.field_names = ["Task Id","Description","Due Date","Start Date","Finish Date","Status","Notes"]

        for task in self.tasks:
            table.add_row([task.task_id,task.task_description,task.due_date,task.start_date,task.finish_date,task.status,task.notes])

        print(table)          


    def generate_pdf(self):
        pdf = PDF_Generator(self.tasks)
        pdf.create_pdf(self.task_report)
        self.tdb.close_db()




