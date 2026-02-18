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

func addTaskTable(pdf *gofpdf.Fpdf, task []models.Task) {
	// Helper function to add table headers
	addTableHeaders := func() {
		pdf.SetFont("Arial", "B", 12)
		pdf.SetFillColor(200, 200, 200)
		pdf.CellFormat(75, 10, "Task", "1", 0, "C", true, 0, "")
		pdf.CellFormat(30, 10, "Due Date", "1", 0, "C", true, 0, "")
		pdf.CellFormat(30, 10, "Start Date", "1", 0, "C", true, 0, "")
		pdf.CellFormat(30, 10, "Finish Date", "1", 0, "C", true, 0, "")
		pdf.CellFormat(30, 10, "Status", "1", 0, "C", true, 0, "")
		pdf.CellFormat(70, 10, "Notes", "1", 1, "C", true, 0, "")
		pdf.SetFont("Arial", "", 12)
	}

	// Add initial headers on first page
	addTableHeaders()

	// Add task rows with automatic page breaks and header repetition
	for i, task := range task {
		// Check if we need a page break before adding the next row
		// Letter page height is ~279mm, with margins we have ~250mm usable space
		// Each row is 10mm high, so we can fit about 24-25 rows per page
		// But we need to account for headers (10mm) and some margin
		// So we'll break every 23 rows to be safe
		if i > 0 && i%15 == 0 {
			pdf.AddPage()
			addTableHeaders() // Re-add headers on new page
		}

		// Truncate long text to prevent overflow
		description := task.Description
		//if len(description) > 25 {
		//	description = description[:22] + "..."
		//}

		notes := task.Notes
		if len(notes) > 50 {
			notes = notes[:45] + "..."
		}

		pdf.CellFormat(75, 10, description, "1", 0, "L", false, 0, "")
		pdf.CellFormat(30, 10, task.DueDate, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, task.StartDate, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, task.FinishDate, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, task.Status, "1", 0, "C", false, 0, "")
		pdf.CellFormat(70, 10, notes, "1", 1, "L", false, 0, "")
	}
}
