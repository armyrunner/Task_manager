# from pdf_generator import PDF_Generator
from task_manager import Task_Manager
from menu import Main_Menu
import os
import sys


TM = Task_Manager()
MM = Main_Menu()

def main():
  os.system("clear")
  MM.menu_items()
 
if __name__ == "__main__":
    main()
