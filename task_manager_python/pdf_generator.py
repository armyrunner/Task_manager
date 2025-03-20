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
        self.col_fontName = 'Helvetica-Bold'
        self.row_fontName = 'Helvetica'
        self.fontSize_title = 20
        self.col_fontSize = 16
        self.data_fontSize = 12


'''
To make the code even cleaner and more maintainable, we can break down the logic into smaller functions, eliminate redundant parts, and improve readability. Below is the refactored code with these improvements:

### Further Refactored Code:

```python
from reportlab.lib.pagesizes import letter
from reportlab.lib import colors
from reportlab.platypus import SimpleDocTemplate, Table, TableStyle

def generate_pdf(self, filename="tasks_report.pdf"):
    """Generate a PDF with current tasks and completed tasks separately."""
    doc = SimpleDocTemplate(filename, pagesize=letter)
    elements = []

    # Prepare task data
    current_tasks = [task for task in self.tasks if task.status != "Completed"]
    completed_tasks = [task for task in self.tasks if task.status == "Completed"]

    # Add title
    elements.append(self.create_title_table("Task Manager Report"))

    # Add current tasks table if available
    if current_tasks:
        elements.append(self.create_task_table("Current Tasks", current_tasks))

    # Add completed tasks table if available
    if completed_tasks:
        elements.append(self.create_task_table("Completed Tasks", completed_tasks))

    # Build the document
    doc.build(elements)

def create_title_table(self, title):
    """Create a table for the title of the document."""
    title_table = Table([[title]])
    title_table.setStyle(TableStyle([
        ('ALIGN', (0, 0), (-1, -1), 'CENTER'),
        ('FONTNAME', (0, 0), (-1, -1), 'Helvetica-Bold'),
        ('FONTSIZE', (0, 0), (-1, -1), 16),
        ('BOTTOMPADDING', (0, 0), (-1, -1), 20)
    ]))
    return title_table

def create_task_table(self, section_title, tasks):
    """Create a table for the list of tasks."""
    # Header for task table
    task_data = [["ID", "Description", "Due Date", "Status"]]
    task_data.extend([[task.task_id, task.description, task.due_date, task.status] for task in tasks])

    task_table = Table(task_data)
    task_table.setStyle(TableStyle([
        ('BACKGROUND', (0, 0), (-1, 0), colors.grey),
        ('TEXTCOLOR', (0, 0), (-1, 0), colors.whitesmoke),
        ('ALIGN', (0, 0), (-1, -1), 'CENTER'),
        ('FONTNAME', (0, 0), (-1, -1), 'Helvetica'),
        ('FONTSIZE', (0, 0), (-1, -1), 10),
        ('BOTTOMPADDING', (0, 0), (-1, -1), 5),
        ('GRID', (0, 0), (-1, -1), 0.5, colors.black)
    ]))

    # Add the section title above the table
    section_title_table = Table([[section_title]])
    section_title_table.setStyle(TableStyle([
        ('ALIGN', (0, 0), (-1, -1), 'CENTER'),
        ('FONTNAME', (0, 0), (-1, -1), 'Helvetica-Bold'),
        ('FONTSIZE', (0, 0), (-1, -1), 12),
        ('BOTTOMPADDING', (0, 0), (-1, -1), 10)
    ]))

    # Return both the section title and table in a list
    return [section_title_table, task_table]
```

### Key Changes:
1. **Modularized Functions:**
   - `create_title_table`: Generates a table for the document title.
   - `create_task_table`: Creates a table for any given section (either current or completed tasks).
   
2. **Simplified Logic:**
   - The task data generation (`task_data`) for both sections is now done in a concise one-liner using list comprehension.
   
3. **Reduced Redundancy:**
   - Both tables (current and completed tasks) now follow the same format and are handled by a single function (`create_task_table`). The only difference is the section title, which is passed as a parameter.

4. **Cleaner Data Population:**
   - Instead of appending each task row one by one, we generate the entire data set for the table in a more compact manner using list comprehension.

5. **Structure and Formatting:**
   - The title and each task table are formatted in a clean, readable way using `TableStyle`.

### Benefits of the Refactor:
- **Maintainability:** The code is now broken down into smaller, reusable functions, making it easier to maintain and extend.
- **Readability:** The logic is more declarative, so it is easier to understand.
- **DRY Principle:** The `create_task_table` function handles both current and completed tasks, eliminating duplication.
- **Scalability:** Adding additional sections in the future (e.g., overdue tasks) would be easy by just calling the `create_task_table` function with different parameters.

Let me know if you need further adjustments!

'''

  
    
    def create_pdf(self): 
        
 
        current_tasks = [task for task in self.tasks if task.status != "Completed" or task.status != 'completed']
        completed_tasks = [task for task in self.tasks if task.status == "Completed" or task.status == 'completed']
        

        if current_tasks:
            self.canvas = canvas.Canvas('cur_task_report.pdf', pagesize=letter)
            self.task_title = 'Current Task List'
                 
            #find the middle of the page to center the title of the report
            text_width = self.canvas.stringWidth(self.task_title,self.fontName_title,self.fontSize_title)
            x_pos = (self.width - text_width) / 2
            y_pos = self.height - 40

            #Drawing the string on the document to see Title of Document
            self.canvas.setFont(self.fontName_title,self.fontSize_title)
            self.canvas.drawString(x_pos,y_pos,self.task_title)

            #creating the table for the document
            pdf_doc = SimpleDocTemplate(self.filename,pagesizes=landscape(letter))
       
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


