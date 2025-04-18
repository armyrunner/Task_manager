import pickle
import os

from prettytable import PrettyTable
from task import Task
from pdf_generator import PDF_Generator

class Task_Manager:

    # init funcion
    def __init__(self,filename ='task.pkl'):
        self.filename = filename
        self.tasks = self.load_tasks()
        self.task_report = "task_report.pdf"
    
        # load task function
    def load_tasks(self):
        if not os.path.exists(self.filename):
            with open(self.filename,'wb') as f:
                pickle.dump([],f)
        try:
            with open(self.filename,'rb') as f:
                return pickle.load(f)
        except (EOFError):
            return []
    
    # save task function
    def save_tasks(self):
        with open(self.filename,'wb') as f:
            pickle.dump(self.tasks,f)
    
    # add task function
    def add_task(self,task_description="",due_date = "",start_date="",finish_date="",status ='Not Started'):
        task_id = len(self.tasks) + 1
        tasks = Task(task_id,task_description,due_date,start_date,finish_date,status)
        self.tasks.append(tasks)        
        self.save_tasks()

    # update task function
    def update_task(self,task_id,task_description = "",due_date = "",start_date="",finish_date="",status ='Not Started'):
        for task in self.tasks:
            if task_id == task.task_id:
                if task_description:
                    task.task_description = task_description
                if due_date:
                    task.due_date = due_date
                if start_date:
                    task.start_date = start_date
                if finish_date:
                    task.finish_date = finish_date
                if status:
                    task.status = status
                self.save_tasks()
                return True
        return False

    # delete task function
    def delete_task(self,task_id):
        self.tasks = [task for task in self.tasks if task.task_id != task_id]
        self.save_tasks()

    def display_tasks(self):
        table = PrettyTable()
        table.field_names = ["Task Id","Description","Due Date","Start Date","Finish Date","Status"]

        for task in self.tasks:
            table.add_row([task.task_id,task.task_description,task.due_date,task.start_date,task.finish_date,task.status])

        print(table)          


    def generate_pdf(self):
        pdf = PDF_Generator(self.tasks)
        pdf.create_pdf(self.task_report)


