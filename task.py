
class Task:

    def __init__(self,task_id,task_description,due_date,status='Not Started'):
        self.task_id = task_id
        self.task_description = task_description
        self.due_date = due_date
        self.status = status
        
    def __rper__(self):
        return f"Task Number {self.task_id}, Task{self.task_description}, Due on{self.due_date}, Status{self.status} "


