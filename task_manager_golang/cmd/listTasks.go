package cmd

import (
	"fmt"

	"github.com/armyrunner/task_manager/db"
	"github.com/spf13/cobra"
)

var verbose bool

var listTasksCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.SelectAllData()
		if err != nil {
			fmt.Println("Failed to fetch tasks:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		if verbose {
			fmt.Println("=== TASK MANAGER - VERBOSE OUTPUT ===")
			fmt.Printf("Total tasks found: %d\n\n", len(tasks))
		}

		if !verbose {
			fmt.Println("ID | Description | Due Date | Start Date | Finish Date | Status | Notes")
			fmt.Println("---|-------------|----------|------------|-------------|--------|------")
		}

		for i, task := range tasks {
			if verbose {
				fmt.Printf("Task %d:\n", i+1)
				fmt.Printf("  ID: %s\n", task[0])
				fmt.Printf("  Description: %s\n", task[1])
				fmt.Printf("  Due Date: %s\n", task[2])
				fmt.Printf("  Start Date: %s\n", task[3])
				fmt.Printf("  Finish Date: %s\n", task[4])
				fmt.Printf("  Status: %s\n", task[5])
				fmt.Printf("  Notes: %s\n", task[6])
				fmt.Println("---")
			} else {
				fmt.Printf("%s | %s | %s | %s | %s | %s | %s\n",
					task[0], task[1], task[2], task[3], task[4], task[5], task[6])
			}
		}

		if verbose {
			fmt.Printf("\nEnd of task list. Total: %d tasks\n", len(tasks))
		}
	},
}

func init() {
	listTasksCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}
