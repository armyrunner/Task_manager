from reportlab.lib.pagesizes import letter,landscape
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
    
    def create_task_table(self, section_title, tasks):
        task_table_data = [["ID", "Task Description", "Due Date", "Start Date", "Finish Date", "Status"]]
        # Add task data
        task_table_data.extend([[task.task_id, task.task_description, task.due_date, task.start_date, task.finish_date, task.status] for task in tasks])
        
        # Create the table object
        task_table = Table(task_table_data)

        # Add alternating row colors for styling
        rowNumb = len(task_table_data)
        for i in range(1, rowNumb):
            if i % 2 == 0:
                bc = colors.lightgrey
            else:
                bc = colors.whitesmoke

        styleConfig_tasks = TableStyle([
            ('BACKGROUND', (0, 0), (-1, 0), colors.grey),
            ('TEXTCOLOR', (0, 0), (-1, 0), colors.whitesmoke),
            ('ALIGN', (0, 0), (-1, -1), 'CENTER'),
            ('FONTNAME', (0, 0), (-1, 0), self.row_fontName),
            ('FONTSIZE', (0, 0), (-1, 0), 14),
            ('BOTTOMPADDING', (0, 0), (-1, 0), 12),
            ('BACKGROUND', (0, 1), (-1, -1), bc),
            ('GRID', (0, 1), (-1, -1), 2, colors.black),
        ])

        task_table.setStyle(styleConfig_tasks)

        # Create section title
        section_title_table = Table([[section_title]])

        section_style_config = TableStyle([
            ('ALIGN', (0, 0), (-1, -1), 'CENTER'),
            ('FONTNAME', (0, 0), (-1, 0), self.row_fontName),
            ('FONTSIZE', (0, 0), (-1, 0), 14),
            ('BOTTOMPADDING', (0, 0), (-1, 0), 12),
        ])
        section_title_table.setStyle(section_style_config)

        # Return the section title and table as flowables
        return section_title_table, task_table

    def create_pdf(self,filename='task_report.pdf'): 

        current_tasks = []
        completed_tasks = []

        for task in self.tasks:
            if task.status.upper() != "COMPLETED":
                # Adding current task or tasks that are not completed
                current_tasks.append(task) 
            else:
                # Adding task that are Completed
                completed_tasks.append(task)

        #creating the table for the document
        pdf_doc = SimpleDocTemplate(filename,pagesizes=landscape(letter))
    	
        elements = []

        elements.append(self.create_title_table("Task Manager Report"))
        if current_tasks:
            current_section, current_table =self.create_task_table("Current Tasks",current_tasks)
            elements.append(current_section)
            elements.append(current_table)

        if completed_tasks:
            completed_section, completed_table = self.create_task_table("Completed Tasks",completed_tasks)
            elements.append(completed_section)
            elements.append(completed_table)

        pdf_doc.build(elements)



