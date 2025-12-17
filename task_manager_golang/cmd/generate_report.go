package cmd

import (
	"fmt"
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

		report_categories := make(map[string][]models.Task)
	
		for _, category := range models.CategoryOrder {
			report_category,err := db.Select_Initial_Tasks_By_Category(models.Task{Category: category})
			if err != nil {
				fmt.Println("Failed to generate initial tasks report:", err)
				return
			}
			report_categories[category] = report_category
		}

		report_completed,err := db.SelectCompletedTasks()
		if err != nil {
			fmt.Println("Failed to generate initial tasks report:", err)
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


