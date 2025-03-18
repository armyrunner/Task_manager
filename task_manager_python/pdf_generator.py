from reportlab.lib.pagesizes import letter,landscape
from reportlab.pdfgen import canvas
from reportlab.platypus import SimpleDocTemplate,Table,TableStyle
from reportlab.lib import colors



class PDF_Generator():

    def __init__(self,tasks):
        self.tasks = tasks
        self.width, self.height = landscape(letter)
        self.filename = self.generate_filename()
        self.canvas = canvas.Canvas(self.filename, pagesize=letter)
        #set variables for the different function calls
        self.task_title = self.generate_task_title()
        self.fontName_title = 'Times-Roman'
        self.col_fontName = 'Helvetica-Bold'
        self.row_fontName = 'Helvetica'
        self.fontSize_title = 20
        self.col_fontSize = 16
        self.data_fontSize = 12

    
    def generate_filename(self):
        
        current_tasks = [task for task in self.tasks if task.status != "Completed" or task.status != 'completed']
        completed_tasks = [task for task in self.tasks if task.status == "Completed" or task.status == 'completed']

        if current_tasks:
            self.filename = 'cur_task_list.pdf'
        elif completed_tasks:
            self.filename = 'comp_task_list.pdf'
        else:
            self.filename = 'task_report.pdf'

    def generate_task_title(self):
        
        current_tasks = [task for task in self.tasks if task.status != "Completed" or task.status != 'completed']
        completed_tasks = [task for task in self.tasks if task.status == "Completed" or task.status == 'completed']

        if current_tasks:
            self.task_title = 'Current Task List'
        elif completed_tasks:
            self.task_title = 'Completed Task List'
        else:
            self.task_title = 'Task Report List'


    def create_pdf(self): 

        current_tasks = [task for task in self.tasks if task.status != "Completed" or task.status != 'completed']
        completed_tasks = [task for task in self.tasks if task.status == "Completed" or task.status == 'completed']
               
        #find the middle of the page to center the title of the report
        text_width = self.canvas.stringWidth(self.task_title,self.fontName_title,self.fontSize_title)
        x_pos = (self.width - text_width) / 2
        y_pos = self.height - 40

        #Drawing the string on the document to see Title of Document
        self.canvas.setFont(self.fontName_title,self.fontSize_title)
        self.canvas.drawString(x_pos,y_pos,self.task_title)

        #creating the table for the document
        pdf_doc = SimpleDocTemplate(
                self.filename,
                pagesizes=landscape(letter)
                )
   
        # adding alternating colors for the table for fancier table look.
        # adding boarder to the table as well.
        rowNumb = len(current_tasks) #Data from Task Manager to determine how many rows
        for i in range(1,rowNumb):
            if i % 2 == 0:
                bc = colors.burlywood
            else:
                bc = colors.beige

        styleConfig = TableStyle([
            ('BACKGROUND',(0,0),(3,0),colors.green),
            ('TEXTCOLOR',(0,0),(-1,0),colors.whitesmoke), # the -1 is the end of the row no matter how long the row is
            ('ALIGN',(0,0),(-1,-1),'CENTER'), # (-1,-1) basically the whole table will be center aligned
            ('FONTNAME', (0,0),(-1,0),self.col_fontName),
            ('FONTNAME',(0,1),(-1,-1),self.row_fontName),
            ('FONTSIZE',(0,0),(-1,0),14),
            ('BOTTOMPADDING',(0,0),(-1,0),12),
            ('BACKGROUND',(0,1),(-1,-1),bc),
            ('GRID',(0,1),(-1,-1),2,colors.black),
            ])
        current_tasks.setStyle(styleConfig) # set table with solid color for all table, with color for column names
        
        pdf_doc.build([current_tasks]])
        self.canvas.save()


