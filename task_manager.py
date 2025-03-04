import pickle
from prettytable import PrettyTable
# from reportlab.lib.pagesizes import letter
# from reportlab.pdfgen import canvas
from task import Task

class Task_Manager:

    # init funcion
    def __init__(self,filename ='task.pkl'):
        self.filename = filename
        self.tasks = self.load_tasks()
        

    # load task function
    def load_tasks(self):
        try:
            with open(self.filename,'rb') as f:
                return pickle.load(f)
        except (FileExistsError, EOFError ):
            return []
    
    # save task function
    def save_task(self):
        with open(self.filename,'wb') as f:
            pickle.dump(self.tasks,f)
    
    # add task function
    def add_task(self,task_description=None,due_date = None,status = None):
        task_id = len(self.tasks) + 1
        tasks = Task(task_id,task_description,due_date,status)
        self.tasks.append(tasks)
        self.save_task()

    # update task function
    def update_task(self,task_id,task_description = None,due_date = None,status = None):
        for tasks in self.tasks:
            if task_id == tasks.task_id:
                Task(task_id,task_description,due_date,status)
                self.save_task()
                return True
            return False

    # delete task function
    def delete_task(self,task_id):
        self.tasks = [task for task in self.tasks if task.task_id != task_id]
        self.save_tasks()

    def display_tasks(self):
        table = PrettyTable
        table.field_names = ["Task Id","Description","Due Date","Status"]

        for tasks in self.tasks:
            tasks = [tasks.task_id,tasks.task_description,tasks.due_date,tasks.status]

        print(table)            

    # generate pdf function
