from task_manager import Task_Manager
from task_db import Task_DB
from simple_term_menu import TerminalMenu
import sys

class Main_Menu():

    def __init__(self):
        self.TM = Task_Manager("tasks.db")
        self.TDB = Task_DB("task.db")
        pass

    def menu_items(self):
        options = {
            "New Task": self.createTask,
            "Update Tasks": self.updateTasks,
            "Delete Tasks": self.deleteTask,
            "View Tasks": self.displayTasks,
            "Exit": self.exitProgram
        }
        self.dict_menu_setup(options)

    def menu_setup(self,options):
        term_menu = TerminalMenu(options)
        menu_index = term_menu.show()
        menu_selection = options[menu_index]
    
        return menu_selection
    
    def dict_menu_setup(self,options):
        menu_selection = self.menu_setup(list(options.keys()))
        get_options = options.get(menu_selection)
        get_options()

    def get_user_string(self,prompt):
        text = input(prompt).strip()
        return text
    
    def createTask(self):
        # def add_task(self,task_description="",due_date = "",start_date="",finish_date="",status ='Not Started')
        print("Create your task by answering some basic questions. if you don't know the answer you can leave it blank.")
        description = self.get_user_string("What is the Task assigned? ")
        due_date = self.get_user_string("When does the Task need to be completed? ")
        self.TM.add_task(description,due_date)
        
        
    def updateTasks(self):
        self.TM.load_tasks()

    def deleteTask(self):
        self.TM.load_tasks()
        self.displayTasks()
        task_id = int(self.get_user_string("What Task do you want to delete? "))
        self.TM.delete_task(task_id)
        self.displayTasks()

    def displayTasks(self):
        self.TM.display_tasks()
        pass

    def exitProgram(self):
        self.TDB.close_db()
        sys.exit(0)
        
