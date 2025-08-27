package cmd

import (
	"fmt"

	"github.com/armyrunner/task_manager/db"
	"github.com/armyrunner/task_manager/services"
	"github.com/spf13/cobra"
)

var generateReportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a report of all tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating report...")
		report_completed,err := db.Select_Completed_Tasks()
		if err != nil {
			fmt.Println("Failed to generate initial tasks report:", err)
			return
		}
		report_incomplete,err := db.Select_Initial_Tasks()
		if err != nil {
			fmt.Println("Failed to generate initial tasks report:", err)
			return
		}
		err = services.PDF_Initial_Tasks(report_incomplete, report_completed, "report.pdf")
		if err != nil {
			fmt.Println("Failed to generate report:", err)
			return
		}
		fmt.Println("Report generated successfully")
	},
}


