
from reportlab.lib.pagesizes import letter
from reportlab.pdfgen import canvas
from reportlab.lib.units import inch
from reportlab.platypus import SimpleDocTemplate
from reportlab.lib.enums import TA_CENTER,TA_JUSTIFY

class PDF_Generator():

    def __init__(self,tasks,fileame="current_task.pdf"):
        self.tasks =  tasks
        self.canvas = canvas.Canvase(filename,pagesize=letter)
        

    def create_pdf(self):
        
        width, height = letter
            
        completed_tasks = [tasks for task in self.tasks if task.status == 'Completed']
        current_tasks = [tasks for task in self.tasks if task.status != 'Completed']

        #Title of Document
        self.canvas.setFont('Times-New-Roman',20)
        self.canvas.drawString(72,heith - 72,'Current Task List')
        self.canvas.save()
        

