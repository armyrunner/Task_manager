package services

import (
	"io"
	"strings"

	"github.com/armyrunner/task_manager/models"
	"github.com/jung-kurt/gofpdf"
)

func setupPDF() *gofpdf.Fpdf{
	pdf := gofpdf.New("L", "mm", "Letter", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	return pdf
}

func addTitle(pdf *gofpdf.Fpdf,title string){
	
	// Title
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, title)
	pdf.Ln(12)

}

func addCategory(pdf *gofpdf.Fpdf,categoryName string){
	
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, " *** "+strings.ToUpper(categoryName)+" TASKS *** ")
	pdf.Ln(12)

}

func addNoDataMessage(pdf *gofpdf.Fpdf, message string){

	pdf.SetFont("Arial", "", 14)
	pdf.Cell(40, 10, message)
	pdf.Ln(12)
}


func GenerateReport(reportType, categoryName string, tasks []models.Task,w io.Writer) error {

	pdf := setupPDF()

	switch reportType {
	case "initial":
		addTitle(pdf,"Initial Tasks Report!")
	case "completed":
		addTitle(pdf,"Completed Tasks Report!")
	case "category":
		addCategory(pdf, categoryName)
	case "full":
		addTitle(pdf, "Full Tasks Report!")
	}
	
	if len(tasks) == 0 {
		addNoDataMessage(pdf, "There are no tasks to display")
	} else{
		addTaskTable(pdf,tasks)
	}

	return pdf.Output(w)
}

func addTaskTable(pdf *gofpdf.Fpdf, tasks []models.Task) {
	// Column widths
	const (
		colDesc    = 75.0
		colDue     = 30.0
		colStart   = 30.0
		colFinish  = 30.0
		colStatus  = 30.0
		colNotes   = 70.0
		lineHeight = 5.0
	)

	// Helper function to add table headers
	addTableHeaders := func() {
		pdf.SetFont("Arial", "B", 12)
		pdf.SetFillColor(200, 200, 200)
		pdf.CellFormat(colDesc, 10, "Task", "1", 0, "C", true, 0, "")
		pdf.CellFormat(colDue, 10, "Due Date", "1", 0, "C", true, 0, "")
		pdf.CellFormat(colStart, 10, "Start Date", "1", 0, "C", true, 0, "")
		pdf.CellFormat(colFinish, 10, "Finish Date", "1", 0, "C", true, 0, "")
		pdf.CellFormat(colStatus, 10, "Status", "1", 0, "C", true, 0, "")
		pdf.CellFormat(colNotes, 10, "Notes", "1", 1, "C", true, 0, "")
		pdf.SetFont("Arial", "", 10)
	}

	// Calculate how many lines a text needs for a given width
	calcLines := func(text string, width float64) int {
		if text == "" {
			return 1
		}
		lines := pdf.SplitText(text, width)
		if len(lines) == 0 {
			return 1
		}
		return len(lines)
	}

	// Add initial headers on first page
	addTableHeaders()

	for _, task := range tasks {
		// Calculate row height based on content
		descLines := calcLines(task.Description, colDesc-2)
		notesLines := calcLines(task.Notes, colNotes-2)
		maxLines := descLines
		if notesLines > maxLines {
			maxLines = notesLines
		}
		rowHeight := float64(maxLines) * lineHeight
		if rowHeight < 10 {
			rowHeight = 10
		}

		// Check if we need a new page
		if pdf.GetY()+rowHeight > 190 {
			pdf.AddPage()
			addTableHeaders()
		}

		// Save starting position
		x := pdf.GetX()
		y := pdf.GetY()

		// Description (MultiCell for wrapping)
		pdf.MultiCell(colDesc, lineHeight, task.Description, "1", "L", false)
		
		// Move to next column position
		pdf.SetXY(x+colDesc, y)
		pdf.CellFormat(colDue, rowHeight, task.DueDate, "1", 0, "C", false, 0, "")
		pdf.CellFormat(colStart, rowHeight, task.StartDate, "1", 0, "C", false, 0, "")
		pdf.CellFormat(colFinish, rowHeight, task.FinishDate, "1", 0, "C", false, 0, "")
		pdf.CellFormat(colStatus, rowHeight, task.Status, "1", 0, "C", false, 0, "")
		
		// Notes (MultiCell for wrapping)
		pdf.SetXY(x+colDesc+colDue+colStart+colFinish+colStatus, y)
		pdf.MultiCell(colNotes, lineHeight, task.Notes, "1", "L", false)

		// Move to next row
		pdf.SetY(y + rowHeight)
	}
}
