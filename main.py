           
from reportlab.lib.pagesizes import letter,landscape
from reportlab.pdfgen import canvas
           
    
      
def create_pdf(filename): 
         
    c = canvas.Canvas(filename,pagesize=landscape(letter))
    width, height = landscape(letter)

    

    text = 'Current Task List'
    font_name = 'Times-Roman'
    font_size = 20
    
    text_width = c.stringWidth(text,font_name,font_size)
    x_pos = (width - text_width) / 2
    y_pos = height - 40

    #Title of Document
    c.setFont(font_name,font_size)
    c.drawString(x_pos,y_pos,text)
    c.save()

def main():

    create_pdf('test.pdf')

if __name__ == "__main__":
    main()
