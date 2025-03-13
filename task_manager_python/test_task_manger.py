import unittest
import os
from task_manager import Task_Manager

class TestTaskManager(unittest.TestCase):

    def setUp(self):
        """Set up the test environment."""
        # We use a temporary file for testing to avoid messing with actual data files
        self.test_filename = "test_tasks.pkl"
        self.manager = Task_Manager(self.test_filename)

    def tearDown(self):
         """Clean up after each test."""
         if os.path.exists(self.test_filename):
             os.remove(self.test_filename)

        
    def test_01_add_task(self):
        """Test adding a task."""
        self.manager.add_task("Test Task 1", "2025-03-10","2025-03-01","2025-03-10","pending")
        self.assertEqual(len(self.manager.tasks), 1)
        self.assertEqual(self.manager.tasks[0].task_description,"Test Task 1")
        self.assertEqual(self.manager.tasks[0].due_date,"2025-03-10")
        self.assertEqual(self.manager.tasks[0].start_date,"2025-03-01")
        self.assertEqual(self.manager.tasks[0].finish_date,"2025-03-10")
        self.assertEqual(self.manager.tasks[0].status,"pending")


    def test_02_update_task(self):
        """ Test updating tasks """
        self.manager.add_task("Test Task 1", "2025-03-10","pending")
        self.manager.add_task("Test Task 2", "2025-03-10","2025-03-01","Started")
        self.manager.add_task("Test Task 3", "2025-03-10","Started")
        
        task_2 = self.manager.update_task(2,"Do Math Homework","2025-03-07","2025-03-07","Need to do problems 5-10")
        task_3 = self.manager.update_task(3,"Oil Change on Car","2025-03-020","Need To Start")

        self.assertTrue(task_2,'Task 3 Updated')
        self.assertTrue(task_3,'Task 4 Updated')

    def test_03_delete_task(self):
        self.manager.add_task("Test Task 1", "2025-03-10","pending")
        self.manager.add_task("Test Task 2", "2025-03-10","Started")
        self.manager.add_task("Test Task 3", "2025-03-10","Started")

        task_2 = self.manager.delete_task(2)
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
        self.manager.add_taskk("Test Task 1", "2025-03-10","2025-03-01","2025-03-10","Completed")
        self.manager.add_task("Test Task 2", "2025-03-10","2025-03-01","2025-03-10","completed")
        self.manager.add_task("Test Task 3", "2025-03-10","pending")
        self.manager.add_task("Test Task 4", "2025-03-10","pending")
        self.manager.add_task("Test Task 5", "2025-03-10")
        self.manager.add_task("Test Task 6", "2025-03-10")
        
        pdfCreated = self.manager.generate_pdf()

        self.assertTrue(pdfCreated,'The pdf was created')
    

    if __name__ == '__main__':
        unittest.main()

