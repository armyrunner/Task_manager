package cmd

import (
	"fmt"
	"strings"

	"github.com/armyrunner/task_manager/db"
	"github.com/armyrunner/task_manager/models"
	"github.com/spf13/cobra"
)

var listByCategoryCmd = &cobra.Command{
	Use:   "category",
	Short: "List all tasks grouped by category",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Show all tasks grouped by category
		listAllByCategory()
	},
}

func listAllByCategory() {
	// Get all tasks
	allTasks, err := db.Select_Initial_Tasks()
	if err != nil {
		fmt.Println("Failed to fetch tasks:", err)
		return
	}

	if len(allTasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	// Group tasks by category name (lowercase for consistent grouping)
	categories := make(map[string][]models.Task)
	for _, task := range allTasks {
		cat := strings.ToLower(task.CategoryName)
		if cat == "" {
			cat = "uncategorized"
		}
		categories[cat] = append(categories[cat], task)
	}

	// Print each category
	for cat, tasks := range categories {
		printCategoryHeader(cat)
		printTasks(tasks)
		fmt.Println()
	}
}

func printCategoryHeader(category string) {
	fmt.Printf("\n=== %s TASKS ===\n", strings.ToUpper(category))
	fmt.Printf("\n")
	fmt.Println("ID   Description        Due Date    Start       Finish      Status       Notes")
	fmt.Println("---  -----------        --------    -----       ------      ------       -----")
}

func printTasks(tasks []models.Task) {
	for _, task := range tasks {
		fmt.Printf("%-4d %-18s %-11s %-11s %-11s %-12s %s\n",
			task.ID,
			truncateString(task.Description, 18),
			task.DueDate,
			task.StartDate,
			task.FinishDate,
			task.Status,
			task.Notes)
	}
}

func init() {
	// Category filtering via CLI is deprecated - use the API instead
}
