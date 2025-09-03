package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"strconv"
	"github.com/armyrunner/task_manager/db"
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
		err = db.DeleteData(taskID)
		if err != nil {
			fmt.Println("Failed to add task:", err)
		} else {
			fmt.Println("Task added successfully!")
		}
	},
}

var taskID int

func init() {
	deleteTaskCmd.Flags().IntVarP(&taskID, "id", "i", 0, "Task ID")
}