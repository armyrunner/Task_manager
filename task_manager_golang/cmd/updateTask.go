package cmd

import (
	"fmt"
	"strconv"

	"github.com/armyrunner/task_manager/db"
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
		existingDesc, existingDue, existingStart, existingFinish, existingStatus, existingNotes, err := db.SelectData(taskID)
		if err != nil {
			fmt.Println("Failed to fetch existing task:", err)
			return
		}

		// Merge new values with existing values (preserve existing if new is empty)
		finalDesc := mergeString(description, existingDesc)
		finalDue := mergeString(duedate, existingDue)
		finalStart := mergeString(startdate, existingStart)
		finalFinish := mergeString(finishdate, existingFinish)
		finalStatus := mergeString(stat, existingStatus)
		finalNotes := mergeString(information, existingNotes)

		// Check if the task is being marked as complete
		if finalStatus == "complete" || finalStatus == "Complete" || finalStatus == "completed" || finalStatus == "Completed" {
			// Move the task to completed_tasks table
			err = db.MoveCompletedTask(taskID)
			if err != nil {
				fmt.Println("Failed to move completed task:", err)
				return
			}
			fmt.Println("Task marked as complete and moved to completed tasks!")
			return
		}

		err = db.UpdateData(taskID, finalDesc, finalDue, finalStart, finalFinish, finalStatus, finalNotes)
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

var task_id, description, duedate, startdate, finishdate, stat, information string

func init() {
	updateTaskCmd.Flags().StringVarP(&task_id, "id", "i", "", "Task ID")
	updateTaskCmd.Flags().StringVarP(&description, "task", "t", "", "Description of Task")
	updateTaskCmd.Flags().StringVarP(&duedate, "due", "d", "", "Due Date of Task")
	updateTaskCmd.Flags().StringVarP(&startdate, "start", "s", "", "Start Date of Task")
	updateTaskCmd.Flags().StringVarP(&finishdate, "finish", "f", "", "Finish Date of Task")
	updateTaskCmd.Flags().StringVarP(&stat, "status", "c", "", "Task Status")
	updateTaskCmd.Flags().StringVarP(&information, "notes", "n", "", "Any Notes")
}
