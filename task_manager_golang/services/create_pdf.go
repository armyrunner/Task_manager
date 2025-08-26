package services

import (
	"fmt"

	"github.com/armyrunner/task_manager/models"
	"github.com/jung-kurt/gofpdf"
)



func PDF_Initial_Tasks(task_incomplete []models.Task, task_completed []models.Task, file_name string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	
	// Title
	pdf.Cell(40, 10, "Task Report")
	pdf.Ln(12)

	// Incomplete Tasks
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Incomplete Tasks")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	if len(task_incomplete) == 0 {
		pdf.Cell(40, 10, "No incomplete tasks")
	} else {
		for _, task := range task_incomplete {
			pdf.Cell(40, 10, fmt.Sprintf("%d. %s %s %s %s %s %s", task.ID, task.Description, task.DueDate, task.StartDate, task.FinishDate, task.Status, task.Notes))
			pdf.Ln(12)
		}
	}

	// Completed Tasks
	pdf.Ln(6)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Completed Tasks")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	if len(task_completed) == 0 {
		pdf.Cell(40, 10, "No completed tasks")
		pdf.Ln(12)
	} else {
		for _, task := range task_completed {
			pdf.Cell(40, 10, fmt.Sprintf("%d. %s %s %s %s %s %s", task.ID, task.Description, task.DueDate, task.StartDate, task.FinishDate, task.Status, task.Notes))
			pdf.Ln(12)
		}
	}
	return pdf.OutputFileAndClose(file_name)
}


