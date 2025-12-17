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
	Short: "List all tasks by category (use -g to filter by specific category)",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		categoryStr := strings.TrimSpace(category)

		// If specific category provided, filter by it
		if categoryStr != "" {
			listSingleCategory(categoryStr)
			return
		}

		// Otherwise, show all tasks grouped by category
		listAllByCategory()
	},
}

func listSingleCategory(categoryStr string) {
	tasks, err := db.Select_Initial_Tasks_By_Category(models.Task{Category: categoryStr})
	if err != nil {
		fmt.Println("Failed to fetch tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Printf("No tasks found for category: %s\n", categoryStr)
		return
	}
	printCategoryHeader(categoryStr)
	printTasks(tasks)
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

	// Group tasks by category (lowercase for consistent grouping)
	categories := make(map[string][]models.Task)
	for _, task := range allTasks {
		cat := strings.ToLower(task.Category)
		if cat == "" {
			cat = "uncategorized"
		}
		categories[cat] = append(categories[cat], task)
	}

	// Print each category in defined order
	categoryOrder := models.CategoryOrder
	printed := make(map[string]bool)

	for _, cat := range categoryOrder {
		if tasks, exists := categories[cat]; exists {
			printCategoryHeader(cat)
			printTasks(tasks)
			fmt.Println()
			printed[cat] = true
		}
	}

	// Print any remaining categories not in the predefined order
	for cat, tasks := range categories {
		if !printed[cat] {
			printCategoryHeader(cat)
			printTasks(tasks)
			fmt.Println()
		}
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
	listByCategoryCmd.Flags().StringVarP(&category, "category", "g", "", "Category to filter by (personal, work, military)")
}
