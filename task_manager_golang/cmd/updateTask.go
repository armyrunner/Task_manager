package cmd

import (
	"fmt"
	"strconv"

	"github.com/armyrunner/task_manager/db"
	"github.com/armyrunner/task_manager/models"
	"github.com/spf13/cobra"
)

var updateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "update a task",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(task_id)
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
		finalDesc := mergeString(description, tasks[0].Description)
		finalDue := mergeString(duedate, tasks[0].DueDate)
		finalStart := mergeString(startdate, tasks[0].StartDate)
		finalFinish := mergeString(finishdate, tasks[0].FinishDate)
		finalStatus := mergeString(stat, tasks[0].Status)
		finalNotes := mergeString(information, tasks[0].Notes)

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
			ID: taskID,
			Description: finalDesc,
			DueDate: finalDue,
			StartDate: finalStart,
			FinishDate: finalFinish,
			Status: finalStatus,
			Notes: finalNotes,
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


var task_ID, description, duedate, startdate, finishdate, stat, information string

func init() {
	updateTaskCmd.Flags().StringVarP(&task_ID, "id", "i", "", "Task ID")
	updateTaskCmd.Flags().StringVarP(&description, "task", "t", "", "Description of Task")
	updateTaskCmd.Flags().StringVarP(&duedate, "due", "d", "", "Due Date of Task")
	updateTaskCmd.Flags().StringVarP(&startdate, "start", "s", "", "Start Date of Task")
	updateTaskCmd.Flags().StringVarP(&finishdate, "finish", "f", "", "Finish Date of Task")
	updateTaskCmd.Flags().StringVarP(&stat, "status", "c", "", "Task Status")
	updateTaskCmd.Flags().StringVarP(&information, "notes", "n", "", "Any Notes")
}
