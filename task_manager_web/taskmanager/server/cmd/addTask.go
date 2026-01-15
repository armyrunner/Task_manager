package cmd

import (
	"fmt"

	"github.com/armyrunner/task_manager/db"
	"github.com/armyrunner/task_manager/models"
	"github.com/spf13/cobra"
)

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if task_description == "" {
			fmt.Println("Task Description Required...")
			return
		}
		err := db.InsertData(&models.Task{
			Description: task_description,
			DueDate:     due_date,
			StartDate:   start_date,
			FinishDate:  finish_date,
			Status:      "Pending",
			Notes:       notes,
			// CategoryID is set via API, not CLI for now
		})
		if err != nil {
			fmt.Println("Failed to add task:", err)
		} else {
			fmt.Println("Task added successfully!")
		}
	},
}

func init() {
	addTaskCmd.Flags().StringVarP(&task_description, "task", "t", "", "Description of Task")
	addTaskCmd.Flags().StringVarP(&due_date, "due", "d", "", "Due Date of Task")
	addTaskCmd.Flags().StringVarP(&start_date, "start", "s", "", "Start Date of Task")
	addTaskCmd.Flags().StringVarP(&finish_date, "finish", "f", "", "Finish Date of Task")
	addTaskCmd.Flags().StringVarP(&status, "status", "c", "", "Task Status")
	addTaskCmd.Flags().StringVarP(&notes, "notes", "n", "", "Any Notes")
	addTaskCmd.Flags().StringVarP(&category, "category", "g", "", "Category of Task")
}
