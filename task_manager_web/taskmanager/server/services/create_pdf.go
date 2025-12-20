package services

import (
	"github.com/armyrunner/task_manager/models"
	"github.com/jung-kurt/gofpdf"
	"strings"
)

func PDF_Initial_Tasks(task_incomplete map[string][]models.Task, task_completed []models.Task, file_name string) error {
	pdf := gofpdf.New("L", "mm", "Letter", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Title
	pdf.Cell(40, 10, "Task Report")
	pdf.Ln(12)
	

	// Current Tasks
	// pdf.SetFont("Arial", "B", 14)
	// pdf.Cell(40, 10, "Current Tasks")
	// pdf.Ln(10)

	if len(task_incomplete) == 0 {
		pdf.Cell(40, 10, "There are no current tasks to display")
	} else {
		for _, category := range models.CategoryOrder {	
			pdf.SetFont("Arial", "B", 14)
			pdf.Cell(40, 10, " *** "+strings.ToUpper(category)+" TASKS *** ")
			pdf.Ln(12)
			if len(task_incomplete[category]) == 0 {
				pdf.Ln(12)
				pdf.Cell(40, 10, "There are no "+strings.ToLower(category)+" tasks to display")
				pdf.Ln(12)

			} else {
				addTaskTable(pdf, task_incomplete[category])
				pdf.AddPage()
			}
		}
	}

	// Add some space before Completed Tasks
	pdf.AddPage()
	//pdf.Ln(10) // Add some spacing instead of a full page break

	// Completed Tasks
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Completed Tasks")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	if len(task_completed) == 0 {
		pdf.Cell(40, 10, "There are no completed tasks to display")
		pdf.Ln(12)
	} else {
		addTaskTable(pdf, task_completed)
	}
	return pdf.OutputFileAndClose(file_name)
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
