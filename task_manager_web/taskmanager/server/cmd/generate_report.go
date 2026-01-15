package cmd

import (
	"fmt"
	"strings"

	"github.com/armyrunner/task_manager/db"
	"github.com/armyrunner/task_manager/models"
	"github.com/armyrunner/task_manager/services"
	"github.com/spf13/cobra"
)

var generateReportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a report of all tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating report...")

		// Get all tasks and group by category name
		allTasks, err := db.Select_Initial_Tasks()
		if err != nil {
			fmt.Println("Failed to fetch tasks:", err)
			return
		}

		report_categories := make(map[string][]models.Task)
		for _, task := range allTasks {
			cat := strings.ToLower(task.CategoryName)
			if cat == "" {
				cat = "uncategorized"
			}
			report_categories[cat] = append(report_categories[cat], task)
		}

		report_completed, err := db.SelectCompletedTasks()
		if err != nil {
			fmt.Println("Failed to generate completed tasks report:", err)
			return
		}

		err = services.PDF_Initial_Tasks(report_categories, report_completed, "report.pdf")
		if err != nil {
			fmt.Println("Failed to generate report:", err)
			return
		}
		fmt.Println("Report generated successfully")
	},
}


