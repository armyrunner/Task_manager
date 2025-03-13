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
        self.curr_task_title = task_title
        self.fontName_title = 'Times-Roman'
        self.col_fontName = 'Courier-bold'
        self.row_fontName = 'Courier'
        self.fontSize_title = 20
        self.col_fontSize = 16
        self.data_fontSize = 12

    def create_pdf(filename): 
             
               
        #find the middle of the page to center the title of the report
        text_width = c.stringWidth(text,font_name,font_size)
        x_pos = (width - text_width) / 2
        y_pos = height - 40

        #Drawing the string on the document to see Title of Document
        c.setFont(font_name,font_size)
        c.drawString(x_pos,y_pos,text)

        #creating the table for the document
        pdf_doc = SimpleDocTemplate(
                filename,
                pagesizes=landscape(letter)
                )
        task_table = Table(self.tasks) #Getting Data from Task_Manager

        elems = []
        elems.append(task_table) # puts all data onto the document
        
        # adding alternating colors for the table for fancier table look.
        # adding boarder to the table as well.
        rowNumb = len(task_table) #Data from Task Manager to determine how many rows
        for i in range(1,rowNumb):
            if i % 2 == 0:
                bc = colors.burlywood
            else:
                bc = colors.beige

        styleConfig = TableStyle([
            ('BACKGROUND',(0,0),(3,0),colors.green),
            ('TEXTCOLOR',(0,0),(-1,0),colors.whitesmoke), # the -1 is the end of the row no matter how long the row is
            ('ALIGN',(0,0),(-1,-1),'CENTER') # (-1,-1) basically the whole table will be center aligned
            ('FONTNAME', (0,0),(-1,0),' Courier-bold),
            ('FONTSIZE',(0,0),(-1,0),14),
            ('BOTTOMPADDING',(0,0),(-1,0),12),
            ('BACKGROUND',(0,1),(-1,-1),bc),
            ('GRID',(0,1),(-1,-1),2,colors.black),
            ])
        table.setStyle(styleConfig) # set table with solid color for all table, with color for column names
        
        c.save()


