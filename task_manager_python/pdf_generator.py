from reportlab.lib.pagesizes import letter,landscape
from reportlab.pdfgen import canvas
from reportlab.platypus import SimpleDocTemplate,Table,TableStyle
from reportlab.lib import colors



class PDF_Generator():

    def __init__(self,tasks):
        self.tasks = tasks
        self.width, self.height = landscape(letter)
        
        #set variables for the different function calls
        self.fontName_title = 'Times-Roman'
        self.fontName = 'Helvetica-Bold'
        self.row_fontName = 'Helvetica'
        self.fontSize_title = 20
        self.col_fontSize = 16
        self.fontSize_data = 12

    def create_title_table(self,title):

        title_table = Table([[title]])

        styleConfig = TableStyle([
                    ('TEXTCOLOR',(0,0),(-1,0),colors.black), # the -1 is the end of the row no matter how long the row is
                    ('ALIGN',(0,0),(-1,-1),'CENTER'), # (-1,-1) basically the whole table will be center aligned
                    ('FONTNAME', (0,0),(-1,0),self.fontName_title),
                    ('FONTSIZE',(0,0),(-1,0),self.fontSize_title),
                    ('BOTTOMPADDING',(0,0),(-1,-1),20)
                    ])
        title_table.setStyle(styleConfig) #set the style of the Title for the report

        return title_table
    
    def create_section_title(self,title):

        section_title_table = Table([[title]])

        styleConfig = TableStyle([
                    ('TEXTCOLOR',(0,0),(-1,0),colors.black), # the -1 is the end of the row no matter how long the row is
                    ('ALIGN',(0,0),(-1,-1),'CENTER'), # (-1,-1) basically the whole table will be center aligned
                    ('FONTNAME', (0,0),(-1,0),self.fontName_title),
                    ('FONTSIZE',(0,0),(-1,0),self.fontSize_data),
                    ('BOTTOMPADDING',(0,0),(-1,-1),10)
                    ])
        section_title_table.setStyle(styleConfig) #set the style of the Title for the report

        return section_title_table
    
    def create_task_table(self,section_title):

        task_table_data = [["ID","Task Description","Due Date","Start Date","Finish Date","Status"]]
        task_table_data.extend([[task['task_id'],task['task_description'],for task in self.tasks]]) 
        # adding alternating colors for the table for fancier table look.
        # adding boarder to the table as well.

        rowNumb = len() #Data from Task Manager to determine how many rows
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


        return[]

    def create_pdf(self): 
        
        current_tasks = [task for task in self.tasks if task.status != "Completed" or task.status != 'completed']
        completed_tasks = [task for task in self.tasks if task.status == "Completed" or task.status == 'completed']

        #creating the table for the document
        pdf_doc = SimpleDocTemplate(self.filename,pagesizes=landscape(letter))
    
        
        
        pdf_doc.build([current_tasks]])



