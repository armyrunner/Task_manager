package cmd

import (
	"fmt"
	"net/http"

	"github.com/armyrunner/task_manager/api"
	"github.com/armyrunner/task_manager/auth"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting server...")

		// Auth routes (public)
		http.HandleFunc("/api/auth/register", auth.RegisterHandler)
		http.HandleFunc("/api/auth/login", auth.LoginHandler)
		http.HandleFunc("/api/auth/refresh", auth.RefreshHandler)
		http.HandleFunc("/api/auth/logout", auth.LogoutHandler)

		// Task routes (protected with auth middleware)
		http.HandleFunc("/api/tasks", auth.AuthMiddleware(api.TaskHandler))

		fmt.Println("Server starting on port 8080...")
		fmt.Println("Auth endpoints:")
		fmt.Println("  POST /api/auth/register - Register new user")
		fmt.Println("  POST /api/auth/login    - Login user")
		fmt.Println("  POST /api/auth/refresh  - Refresh access token")
		fmt.Println("  POST /api/auth/logout   - Logout user")
		fmt.Println("Protected endpoints:")
		fmt.Println("  /api/tasks - Requires Authorization: Bearer <token>")

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("Error starting server:", err)
			return
		}
	},
}