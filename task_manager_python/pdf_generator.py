from reportlab.lib.pagesizes import letter,landscape
from reportlab.pdfgen import canvas
from reportlab.platypus import SimpleDocTemplate,Table,TableStyle
from reportlab.lib import colors

class PDF_Generator():

    def __init__(self,tasks,filename,task_title):
        self.filename = filename
        self.tasks = tasks
        self.canvas = canvas.Canvas(filename,pagesize=landscape(letter))
        self.width, self.height = landscape(letter)
        
        #set variables for the different function calls
        self.task_title = task_title
        self.fontName_title = 'Times-Roman'
        self.col_fontName = 'Helvetica-Bold'
        self.row_fontName = 'Helvetica'
        self.fontSize_title = 20
        self.col_fontSize = 16
        self.data_fontSize = 12

    def create_pdf(self): 
             
               
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
        task_data = []
        for task in self.tasks:
            task_data.append([task.task_id,task.task_description,task.due_date,task.start_date,task.finish_date,task.status])

        task_table = Table([["Task ID","Task Description",
                             "Due Date","Start Date","Finish Date","Status"]]+ task_data) #Getting Data from Task_Manager

        # adding alternating colors for the table for fancier table look.
        # adding boarder to the table as well.
        rowNumb = len(task_data) #Data from Task Manager to determine how many rows
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
        task_table.setStyle(styleConfig) # set table with solid color for all table, with color for column names
        
        pdf_doc.build([task_table])
        self.canvas.save()


