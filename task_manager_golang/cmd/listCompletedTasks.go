package cmd

import (
	"fmt"

	"github.com/armyrunner/task_manager/db"
	"github.com/spf13/cobra"
)

var listCompletedTasksCmd = &cobra.Command{
	Use:   "completed",
	Short: "list completed tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.SelectCompletedTasks()
		if err != nil {
			fmt.Println("Failed to fetch completed tasks:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No completed tasks found.")
			return
		}

		fmt.Println("Completed Tasks:")
		fmt.Println("ID   Description        Due Date    Start       Finish  Status       Notes")
		fmt.Println("---  -----------        --------    -----       ------  ------       -----")
		for _, task := range tasks {
			fmt.Printf("%-4d %-18s %-11s %-11s %-7s %-12s %s\n",
				task.ID,
				truncateString(task.Description, 18),
				task.DueDate,
				task.StartDate,
				task.FinishDate,
				task.Status,
				task.Notes)
		}
	},
}

// truncateString truncates a string to the specified length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
