package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/armyrunner/task_manager/db"
	"github.com/armyrunner/task_manager/models"
	"github.com/spf13/cobra"
)

var verbose bool

var listTasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List all Initial tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {



		tasks, err := db.Select_Initial_Tasks()
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
			fmt.Printf("Total tasks: %d\n\n", len(tasks))
			for i, task := range tasks {
				printVerboseTask(task, i)
			}
			fmt.Printf("\nEnd of task list. Total: %d tasks\n", len(tasks))
		} else {
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tDescription\tDue Date\tStart\tFinish\tStatus\tNotes\tCategory")
			fmt.Fprintln(w, "---\t-----------\t--------\t-----\t------\t------\t------\t------")
			for _, task := range tasks {
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
					task.ID, task.Description, task.DueDate, task.StartDate,
					task.FinishDate, task.Status, task.Notes, task.CategoryName)
			}
			w.Flush()
		}
	},
}

func printVerboseTask(task models.Task, i int) {
	fmt.Printf("Task %d:\n", i+1)
	fmt.Printf("  ID: %d\n", task.ID)
	fmt.Printf("  Description: %s\n", task.Description)
	fmt.Printf("  Due Date: %s\n", task.DueDate)
	fmt.Printf("  Start Date: %s\n", task.StartDate)
	fmt.Printf("  Finish Date: %s\n", task.FinishDate)
	fmt.Printf("  Status: %s\n", task.Status)
	fmt.Printf("  Notes: %s\n", task.Notes)
	fmt.Printf("  Category: %s\n", task.CategoryName)		
	fmt.Println("---")
}

func init() {
	listTasksCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}
