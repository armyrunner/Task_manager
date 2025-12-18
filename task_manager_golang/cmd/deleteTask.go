package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/armyrunner/task_manager/db"
	"github.com/armyrunner/task_manager/models"
	"github.com/spf13/cobra"
)

var deleteTaskCmd = &cobra.Command{
	Use:   "deleteTask",
	Short: "Delete a task",
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
		err = db.DeleteData(&models.Task{ID: taskID})
		if err != nil {
			fmt.Println("Failed to add task:", err)
		} else {
			fmt.Println("Task deleted successfully!")
		}
	},
}

func init() {
	deleteTaskCmd.Flags().StringVarP(&task_id, "id", "i", "", "Task ID")
}
