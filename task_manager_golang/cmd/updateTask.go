package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/armyrunner/task_manager/db"
	"github.com/armyrunner/task_manager/models"
	"github.com/spf13/cobra"
)

var updateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "update a task",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Validate and trim task ID
		taskIDStr := strings.TrimSpace(task_id)
		if taskIDStr == "" {
			fmt.Println("Error: Task ID is required. Please provide --id or -i flag")
			return
		}

		taskID, err := strconv.Atoi(taskIDStr)
		if err != nil {
			fmt.Println("Invalid task ID:", err)
			return
		}

		// Fetch existing task data
		tasks, err := db.SelectData(models.Task{ID: taskID})
		if err != nil {
			fmt.Println("Failed to fetch existing task:", err)
			return
		}

		// Merge new values with existing values (preserve existing if new is empty)
		finalDesc := mergeString(task_description, tasks[0].Description)
		finalDue := mergeString(due_date, tasks[0].DueDate)
		finalStart := mergeString(start_date, tasks[0].StartDate)
		finalFinish := mergeString(finish_date, tasks[0].FinishDate)
		finalStatus := mergeString(status, tasks[0].Status)
		finalNotes := mergeString(notes, tasks[0].Notes)
		finalCategory := mergeString(category, tasks[0].Category)

		// Check if the task is being marked as complete
		if finalStatus == "complete" || finalStatus == "Complete" || finalStatus == "completed" || finalStatus == "Completed" {
			// Move the task to completed_tasks table
			err = db.MoveCompletedTask(models.Task{ID: taskID})
			if err != nil {
				fmt.Println("Failed to move completed task:", err)
				return
			}
			fmt.Println("Task marked as complete and moved to completed tasks!")
			return
		}

		err = db.UpdateData(models.Task{
			ID:          taskID,
			Description: finalDesc,
			DueDate:     finalDue,
			StartDate:   finalStart,
			FinishDate:  finalFinish,
			Status:      finalStatus,
			Notes:       finalNotes,
			Category:    finalCategory,
		})
		if err != nil {
			fmt.Println("Failed to update task:", err)
		} else {
			fmt.Println("Task updated successfully!")
		}
	},
}

// mergeString returns newValue if it's not empty, otherwise returns existingValue
func mergeString(newValue, existingValue string) string {
	if newValue != "" {
		return newValue
	}
	return existingValue
}

func init() {
	updateTaskCmd.Flags().StringVarP(&task_id, "id", "i", "", "Task ID")
	updateTaskCmd.Flags().StringVarP(&task_description, "task", "t", "", "Description of Task")
	updateTaskCmd.Flags().StringVarP(&due_date, "due", "d", "", "Due Date of Task")
	updateTaskCmd.Flags().StringVarP(&start_date, "start", "s", "", "Start Date of Task")
	updateTaskCmd.Flags().StringVarP(&finish_date, "finish", "f", "", "Finish Date of Task")
	updateTaskCmd.Flags().StringVarP(&status, "status", "c", "", "Task Status")
	updateTaskCmd.Flags().StringVarP(&notes, "notes", "n", "", "Any Notes")
	updateTaskCmd.Flags().StringVarP(&category, "category", "g", "", "Category of Task")
}
