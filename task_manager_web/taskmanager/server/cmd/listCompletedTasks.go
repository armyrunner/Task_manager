package cmd

import (
	"fmt"


	"github.com/armyrunner/task_manager/db"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
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
		if verbose {
			fmt.Println("=== TASK MANAGER - VERBOSE OUTPUT ===")
			fmt.Printf("Total tasks: %d\n\n", len(tasks))
			for i, task := range tasks {
				printVerboseTask(task, i)
			}
			fmt.Printf("\nEnd of task list. Total: %d tasks\n", len(tasks))
		} else {
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tOrigID\tDescription\tDue Date\tStart\tFinish\tStatus\tNotes\tCategory")
			fmt.Fprintln(w, "---\t-------\t-----------\t--------\t-----\t------\t------\t------\t------")
			for _, task := range tasks {
				fmt.Fprintf(w, "%d\t%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
					task.ID, task.OriginalID, task.Description, task.DueDate, task.StartDate,
					task.FinishDate, task.Status, task.Notes, task.CategoryName)
			}
			w.Flush()
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
