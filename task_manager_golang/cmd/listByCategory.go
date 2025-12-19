package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

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
	tasks, err := db.Select_Initial_Tasks_By_Category(&models.Task{Category: categoryStr})
	if err != nil {
		fmt.Println("Failed to fetch tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Printf("No tasks found for category: %s\n", categoryStr)
		return
	}
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
			printTasks(tasks)
			fmt.Println()
			printed[cat] = true
		}
	}

	// Print any remaining categories not in the predefined order
	for cat, tasks := range categories {
		if !printed[cat] {
			printTasks(tasks)
			fmt.Println()
		}
	}
}

func printTasks(tasks []models.Task) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tDescription\tDue Date\tStart\tFinish\tStatus\tNotes\tCategory")
	fmt.Fprintln(w, "---\t-----------\t--------\t-----\t------\t------\t------\t------")
	for _, task := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			task.ID, task.Description, task.DueDate, task.StartDate,
			task.FinishDate, task.Status, task.Notes, task.Category)
	}
	w.Flush()
}

func init() {
	listByCategoryCmd.Flags().StringVarP(&category, "category", "g", "", "Category to filter by (personal, work, military)")
}
