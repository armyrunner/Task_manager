package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"strconv"
	"github.com/armyrunner/task_manager/db"
	"github.com/armyrunner/task_manager/models"
)

var deleteTaskCmd = &cobra.Command{
	Use:   "deleteTask",
	Short: "Delete a task",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		taskID, err := strconv.Atoi(task_id)
		if err != nil {
			fmt.Println("Invalid task ID:", err)
			return
		}
		err = db.DeleteData(models.Task{ID: taskID})
		if err != nil {
			fmt.Println("Failed to add task:", err)
		} else {
			fmt.Println("Task deleted successfully!")
		}
	},
}

var task_id string

func init() {
	deleteTaskCmd.Flags().StringVarP(&task_id, "id", "i", "", "Task ID")
}