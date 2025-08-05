import unittest
import os
import sqlite3
import task_db
from task_manager import Task_Manager


DB_FILENAME = "test_tasks.db"

def dict_factory(cursor,row):
    d = {}
    for idx, col in enumerate(cursor.description):
        d[col[0]] = row[idx]
    return d

class TestTaskManager(unittest.TestCase):

    def setUp(self):
        """Set up the test environment."""
        # We use a temporary file for testing to avoid messing with actual data files
        self.sql_file = "./schemas.sql"

        with open(self.sql_file,'r') as f:
            sql_script = f.read()

        self.conn = sqlite3.connect(DB_FILENAME)
        self.conn.row_factory = dict_factory

        self.cursor = self.conn.cursor()
        self.cursor.executescript(sql_script)
        self.conn.commit()

        self.manager = Task_Manager(DB_FILENAME)
        self.taskDB = task_db.Task_DB(DB_FILENAME)

    def tearDown(self):
        """Clean up after each test."""
        self.cursor.execute("DELETE FROM initial_tasks")
        self.cursor.execute("DELETE FROM completed_tasks")
        self.conn.commit()
        self.conn.close()
      
    def test_01_add_task(self):
        """Test adding a task."""
        self.manager.add_task("Test Task 1", "2025-03-10","2025-03-01","2025-03-10","pending")
        self.manager.tasks = self.manager.load_tasks()
        self.assertEqual(len(self.manager.tasks), 1)

    def test_02_update_task(self):
        """ Test updating tasks """
        self.manager.add_task("Test Task 2", due_date="2025-03-10",status="pending")
        self.manager.add_task("Test Task 3", due_date="2025-03-10",start_date="2025-03-01",status="Started")
        self.manager.add_task("Test Task 4", due_date="2025-03-10",status="Started")

        self.manager.display_tasks()
        
        # task_2 = self.manager.update_task(2,"Do Math Homework","2025-03-10","2025-03-07","2025-03-07","Need to do problems 5-10")
        # task_3 = self.manager.update_task(3,"Oil Change on Car","2025-03-10",None,None,"Need To Start")

        task_2 = self.manager.update_task(2,task_description="Do Math Homework",start_date="2025-03-10",finish_date=None,status="Need to do problems 5-10")
        task_3 = self.manager.update_task(3,task_description="Oil Change on Car",start_date="2025-03-10",finish_date=None,status="Need To Start")

        self.manager.display_tasks()

        self.assertEqual(task_2,'Update Successful')
        self.assertEqual(task_3,'Update Successful')

    def test_03_delete_task(self):
        self.manager.add_task("Test Task 1", "2025-03-10",None,None,"pending")
        self.manager.add_task("Test Task 2", "2025-03-10",None,None,"pending")
        self.manager.add_task("Test Task 3", "2025-03-10",None,None,"Started")

        task_id = self.taskDB.get_task_id("Test Task 2")
        self.manager.delete_task(task_id["id"])
        self.manager.display_tasks()
        self.assertEqual(len(self.manager.tasks),2)


    def test_04_display_task(self):
        self.manager.add_task("Test Task 1", "2025-03-10","pending")
        self.manager.add_task("Test Task 2", "2025-03-10","Started")
        self.manager.add_task("Test Task 3", "2025-03-10","Started")
        self.manager.add_task("Test Task 4", "2025-03-10","pending")
        self.manager.add_task("Test Task 5", "2025-03-10")
        self.manager.add_task("Test Task 6", "2025-03-10")
        
        self.manager.display_tasks()
        self.assertEqual(len(self.manager.tasks),6) 

 
    def test_05_create_pdf_task_report(self):
        self.manager.add_task("Complete Project Report","2025-03-20","2025-03-10","2025-03-15","Will be starting the project soon")
        self.manager.add_task("Fix Bugs","2025-03-22","2025-03-12","2025-03-18","In Progress")
        self.manager.add_task("Submit Final Version","2025-03-25","2025-03-15","2025-03-20","Completed")
        self.manager.add_task("Complete Project Report","2025-03-20","2025-03-10","2025-03-15","Will be starting the project soon")
        self.manager.add_task("Fix Bugs","2025-03-22","2025-03-12","2025-03-18","In Progress")
        self.manager.add_task("Submit Final Version","2025-03-25","2025-03-15","2025-03-20","Completed")
        self.manager.add_task("Complete Project Report","2025-03-20","2025-03-10","2025-03-15","Will be starting the project soon")
        self.manager.add_task("Fix Bugs","2025-03-22","2025-03-12","2025-03-18","In Progress")
        self.manager.add_task("Submit Final Version","2025-03-25","2025-03-15","2025-03-20","Completed")
        self.manager.display_tasks()
    
        self.manager.generate_pdf()

        assert os.path.exists(self.manager.task_report)



    if __name__ == '__main__':
        unittest.main()

