
from reportlab.lib.pagesizes import letter
from reportlab.pdfgen import canvas
from reportlab.lib.units import inch
from reportlab.platypus import SimpleDocTemplate
from reportlab.lib.enums import TA_CENTER,TA_JUSTIFY

class PDF_Generator():

    def __init__(self,filename):
        self.pdf_filename = filename


    def create_pdf(self,task_id,task_description,due_date,start_date,finish_date,status):
        
        task_document 

