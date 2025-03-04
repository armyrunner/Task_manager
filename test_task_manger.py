import unittest
import os
from task_manager import Task_Manager

class TestTaskManager(unittest.TestCase):

    def setUp(self):
        """Set up the test environment."""
        # We use a temporary file for testing to avoid messing with actual data files
        self.test_filename = "test_tasks.pkl"
        self.manager = Task_Manager(self.test_filename)

    # def tearDown(self):
    #     """Clean up after each test."""
    #     if os.path.exists(self.test_filename):
    #         os.remove(self.test_filename)

    def test_add_task(self):
        """Test adding a task."""
        self.manager.add_task("Test Task 1", "2025-03-10","pending")
        self.assertEqual(len(self.manager.tasks), 1)
        self.assertEqual(self.manager.tasks[0].task_description,"Test Task 1")
        self.assertEqual(self.manager.tasks[0].due_date,"2025-03-10")
        self.assertEqual(self.manager.tasks[0].status,"pending")



