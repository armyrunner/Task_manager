
from prettytable import PrettyTable
from task import Task
from pdf_generator import PDF_Generator
from task_db import Task_DB

import os

class Task_Manager:

    # init funcion
    def __init__(self): 
        self.tasks = self.load_tasks()
        self.task_report = "task_report.pdf"
        self.tdb = Task_DB()
    
        # load task function
    def load_tasks(self):
        all_tasks = self.tdb.get_all_tasks()
        tasks = []
        for row in all_tasks:
            task_id,descritption,due_date,start_date,finsih_date,status = row
            tasks.append(Task(task_id,descritption,due_date,start_date,finsih_date,status))
        return tasks
    
    # save task function
    def save_tasks(self):
        # Might apply later for larger program. reserved for bulk saving, more fore GUI
        pass
    
    # add task function
    def add_task(self,task_description="",due_date = "",start_date="",finish_date="",status ='Not Started'):
        task_id = self.tdb.add_task(task_description,due_date,start_date,finish_date,status)
        tasks = Task(task_id,task_description,due_date,start_date,finish_date,status)
        self.tasks.append(tasks)        


    # update task function
    def update_task(self,task_id,task_description = "",due_date = "",start_date="",finish_date="",status ='Not Started'):
        for task in self.tasks:
            if task.task_id == task_id:
                task_description = task_description or task.task_description
                due_date = due_date or task.due_date
                start_date = start_date or task.start_date
                finish_date = finish_date or task.finish_date
                status = status or task.status
        
            success = self.tdb.update_tasks(task_id,task_description,due_date,start_date,finish_date,status)

            if success:
                task.task_description = task_description
                task.due_date = due_date
                task.start_date = start_date
                task.finish_date = finish_date
                task.status = status
                return True
            return False
        return False

    # delete task function
    def delete_task(self,task_id):
        removed = self.tdb.delete_task(task_id)
        if removed:
            self.tasks = [task for task in self.tasks if task.task_id != task_id]

    def display_tasks(self):
        table = PrettyTable()
        table.field_names = ["Task Id","Description","Due Date","Start Date","Finish Date","Status"]

        for task in self.tasks:
            table.add_row([task.task_id,task.task_description,task.due_date,task.start_date,task.finish_date,task.status])

        print(table)          


    def generate_pdf(self):
        pdf = PDF_Generator(self.tasks)
        pdf.create_pdf(self.task_report)


