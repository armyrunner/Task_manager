from reportlab.lib.pagesizes import letter,landscape
from reportlab.pdfgen import canvas
from reportlab.platypus import SimpleDocTemplate,Table,TableStyle
from reportlab.lib import colors
from task_manager import Task_Manager

class PDF_Generator():

    def __init__(self,filename):
        self.filename = filename
        self.tasks = task_manager.Task_manager()
        self.canvas = canvas.Canvas(filename,pagesize=landscape(letter))
        self.width, self.height = landscape(letter)
        
        #set variables for the different function calls
        self.curr_task_title = 'Current Task List'
        self.comp_task_title = 'Completed Task List'
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
        cur_task_table = Table() #TODO need data for the table

        elems = []
        elems.append(cur_task_table) # puts all data onto the document
        
        style = TableStyle([
            ('BACKGROUND',(0,0),(3,0),colors.green),
            ('TEXTCOLOR',(0,0),(-1,0),colors.whitesmoke), # the -1 is the end of the row no matter how long the row is
            ('ALIGN',(0,0),(-1,-1),'CENTER') # (-1,-1) basically the whole table will be center aligned
            ('FONTNAME', (0,0),(-1,0),' Courier-bold),
            ('FONTSIZE',(0,0),(-1,0),14),
            ('BOTTOMPADDING',(0,0),(-1,0),12),
            ('BACKGROUND',(0,1),(-1,-1),colors.beige),

            ])
        table.setStyle(style) # set table with solid color for all table, with color for column names
        
        # adding alternating colors for the table for fancier table look.
        # adding boarder to the table as well.
        rowNumb = len() # TODO Need data to get the length of rows
        for i in range(1,rowNumb):
            if i % 2 == 0:
                bc = colors.burlywood
            else:
                bc = colors.beige
        
        
        alt_color_table = TableStyle([
            ('BACKGROUND',(0,1),(-1,-1),bc),
            ('BOX',(0,0),(-1,-1),2,colors.black),
            ('GRID',(0,1),(-1,-1),2,colors.black),
            ])
        table.setStyle(alt_color_table)

        c.save()


