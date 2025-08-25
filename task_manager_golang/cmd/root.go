package cmd

import (
	"fmt"
	"os"

	"github.com/armyrunner/task_manager/db"
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "taskmanager",
	Short: "A simple Task Manager CLI Tool!",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		err := db.OpenDatabase("db/tasks.db", "db/schema.sql")
		if err != nil {
			fmt.Println("Error opening database:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCMD.AddCommand(addTaskCmd)
	rootCMD.AddCommand(listTasksCmd)
}
