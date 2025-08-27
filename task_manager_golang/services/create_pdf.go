package services

import (
	"github.com/armyrunner/task_manager/models"
	"github.com/jung-kurt/gofpdf"
)

func PDF_Initial_Tasks(task_incomplete []models.Task, task_completed []models.Task, file_name string) error {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Title
	pdf.Cell(40, 10, "Task Report")
	pdf.Ln(12)

	// Current Tasks
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Current Tasks")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	if len(task_incomplete) == 0 {
		pdf.Cell(40, 10, "No incomplete tasks")
	} else {
		addTaskTable(pdf, task_incomplete)
	}

	// Add page break before Completed Tasks
	pdf.AddPage()

	// Completed Tasks
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Completed Tasks")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	if len(task_completed) == 0 {
		pdf.Cell(40, 10, "No completed tasks")
		pdf.Ln(12)
	} else {
		addTaskTable(pdf, task_completed)
	}
	return pdf.OutputFileAndClose(file_name)
}

func addTaskTable(pdf *gofpdf.Fpdf, task []models.Task) {
	//Set Header Style
	pdf.SetFont("Arial", "B", 12)
	pdf.SetFillColor(200, 200, 200)
	pdf.CellFormat(75, 10, "Task", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 10, "Due Date", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 10, "Start Date", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 10, "Finish Date", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 10, "Status", "1", 0, "C", true, 0, "")
	pdf.CellFormat(70, 10, "Notes", "1", 1, "C", true, 0, "")

	//Set Row Style
	pdf.SetFont("Arial", "", 12)
	for i, task := range task {
		// Check if we need a page break (every 15 rows)
		if i > 0 && i%15 == 0 {
			pdf.AddPage()
			// Re-add headers on new page
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

		pdf.CellFormat(75, 10, task.Description, "1", 0, "L", false, 0, "")
		pdf.CellFormat(30, 10, task.DueDate, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, task.StartDate, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, task.FinishDate, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, task.Status, "1", 0, "C", false, 0, "")
		pdf.CellFormat(70, 10, task.Notes, "1", 1, "L", false, 0, "")
	}
}
