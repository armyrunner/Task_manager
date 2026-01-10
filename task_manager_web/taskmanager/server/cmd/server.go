package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/armyrunner/task_manager/api"
	"net/http"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting server...")
		
		http.HandleFunc("/api/tasks", api.TaskHandler)

		fmt.Println("Server starting on port 8080...")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("Error starting server:", err)
			return
		}
	},
}