from task_manager import Task_Manager
from task_db import Task_DB
from simple_term_menu import TerminalMenu
import sys
import os

class Main_Menu():

    def __init__(self):
        self.TM = Task_Manager("task.db")
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

    def newTask_next_choice(self):
        options ={
            "Add Tasks":self.createTask,
            "View Tasks":self.displayTasks,
            "Main Menu":self.menu_items
        }
        self.dict_menu_setup(options)

    def viewTask_next_choice(self):
        options ={
            "Delete Tasks": self.deleteTask,
            "Print Tasks":self.printPDF,
            "Main Menu":self.menu_items
        }
        self.dict_menu_setup(options)

    def updateTask_next_choice(self):
        options ={
            "Update Tasks": self.updateTasks,
            "Print Tasks":self.printPDF,
            "Main Menu":self.menu_items
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
        while True:
            os.system("clear")
            print("Create your task by answering some basic questions. if you don't know the answer you can leave it blank.")
            description = self.get_user_string("What is the Task assigned? ")
            due_date = self.get_user_string("When does the Task need to be completed? ")
            self.TM.add_task(description,due_date)
            
            self.newTask_next_choice()
        
        
    def updateTasks(self):
        #update_task(self,task_id,task_description =None,due_date = None,start_date=None,finish_date=None,status ='Not Started',notes=None):
        while True:
            os.system("clear")
            self.TM.display_tasks()
            task_id = int(self.get_user_string("Which Task needs to be updated? "))
            print("Create your task by answering some basic questions. if you don't know the answer you can leave it blank.")
            description = self.get_user_string("What is the Task assigned? ")
            due_date = self.get_user_string("When does the Task need to be completed? ")
            start_date = self.get_user_string("When did you start the task? ")
            finish_date = self.get_user_string("When did you finish? ")
            status = self.get_user_string("What is the status of the Task?(Started,In-Progress,Complete/Completed)")
            notes = self.get_user_string("Any Notes for this Task(Keep it Brief)? ")
            self.TM.update_task(task_id,description,due_date,start_date,finish_date,status,notes)
            
            self.newTask_next_choice()


    def deleteTask(self):
        os.system("clear")
        self.TM.load_tasks()
        self.displayTasks()
        task_id = int(self.get_user_string("What Task do you want to delete? "))
        self.TM.delete_task(task_id)
        self.displayTasks()

    def displayTasks(self):
        os.system("clear")
        self.TM.display_tasks()
        self.viewTask_next_choice()

    def printPDF(self):
        self.TM.generate_pdf()
        os.system()
        self.menu_items()

    def exitProgram(self):
        os.system()
        self.TDB.close_db()
        sys.exit(0)
        
