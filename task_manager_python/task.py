
class Task:

    def __init__(self,task_id,task_description,due_date= None,start_date=None,finish_date=None,status='Not Started',notes=None):
        self.task_id = task_id
        self.task_description = task_description
        self.due_date = due_date
        self.start_date = start_date
        self.finish_date = finish_date
        self.status = status
        self.notes = notes
        
    def __rper__(self):
        return f"Task Number {self.task_id}, Task{self.task_description}, Due on{self.due_date}, Status{self.status},Notes{self.notes} "


