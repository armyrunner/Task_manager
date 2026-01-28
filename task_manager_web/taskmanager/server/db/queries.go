package db

import (
	"database/sql"
	"fmt"
	
	"github.com/armyrunner/task_manager/models"
)

// =============================================================================
// Category Functions
// =============================================================================

// CreateCategory creates a new category for a user
func CreateCategory(cat *models.Category) error {
	stmt, err := DB.Prepare("INSERT INTO categories (user_id, name) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("preparing create category statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(cat.UserID, cat.Name)
	if err != nil {
		return fmt.Errorf("executing create category statement: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("getting last insert ID for category: %w", err)
	}
	cat.ID = int(id)
	return nil
}

// GetCategoriesByUserID retrieves all categories for a user
func GetCategoriesByUserID(userID int) ([]models.Category, error) {
	rows, err := DB.Query("SELECT id, user_id, name, created_at FROM categories WHERE user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("querying categories: %w", err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		err := rows.Scan(&cat.ID, &cat.UserID, &cat.Name, &cat.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scanning category row: %w", err)
		}
		categories = append(categories, cat)
	}
	return categories, nil
}

// GetCategoryByID retrieves a category by ID
func GetCategoryByID(id int) (*models.Category, error) {
	cat := &models.Category{}
	err := DB.QueryRow("SELECT id, user_id, name, created_at FROM categories WHERE id = ?", id).
		Scan(&cat.ID, &cat.UserID, &cat.Name, &cat.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("querying category by ID: %w", err)
	}
	return cat, nil
}

// UpdateCategory updates a category name
func UpdateCategory(cat *models.Category) error {
	result, err := DB.Exec("UPDATE categories SET name = ? WHERE id = ? AND user_id = ?", cat.Name, cat.ID, cat.UserID)
	if err != nil {
		return fmt.Errorf("updating category: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no category found with ID %d for user %d", cat.ID, cat.UserID)
	}
	return nil
}

// DeleteCategory deletes a category
func DeleteCategory(id, userID int) error {
	result, err := DB.Exec("DELETE FROM categories WHERE id = ? AND user_id = ?", id, userID)
	if err != nil {
		return fmt.Errorf("deleting category: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no category found with ID %d for user %d", id, userID)
	}
	return nil
}

// =============================================================================
// Task Functions
// =============================================================================

func InsertData(tks *models.Task) error {
	stmt, err := DB.Prepare("INSERT INTO initial_tasks (user_id, category_id, task_description, due_date, start_date, finish_date, status, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Handle null category_id
	var categoryID interface{}
	if tks.CategoryID == 0 {
		categoryID = nil
	} else {
		categoryID = tks.CategoryID
	}

	result, err := stmt.Exec(tks.UserID, categoryID, tks.Description, tks.DueDate, tks.StartDate, tks.FinishDate, tks.Status, tks.Notes)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	tks.ID = int(id)
	return nil
}

func UpdateData(tks *models.Task) error {
	stmt, err := DB.Prepare("UPDATE initial_tasks SET category_id = ?, task_description = ?, due_date = ?, start_date = ?, finish_date = ?, status = ?, notes = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Handle null category_id
	var categoryID interface{}
	if tks.CategoryID == 0 {
		categoryID = nil
	} else {
		categoryID = tks.CategoryID
	}

	result, err := stmt.Exec(categoryID, tks.Description, tks.DueDate, tks.StartDate, tks.FinishDate, tks.Status, tks.Notes, tks.ID)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no tasks found with ID %d", tks.ID)
	}

	return nil
}

func DeleteData(tks *models.Task) error {
	stmt, err := DB.Prepare("DELETE FROM initial_tasks WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tks.ID)
	if err != nil {
		return err
	}

	return nil
}

func SelectData(tks *models.Task) ([]models.Task, error) {
	stmt, err := DB.Prepare(`
		SELECT t.id, t.user_id, t.category_id, COALESCE(c.name, '') as category_name, 
		       t.task_description, t.due_date, t.start_date, t.finish_date, t.status, t.notes 
		FROM initial_tasks t
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE t.id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(tks.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		var categoryID sql.NullInt64
		var userID sql.NullInt64
		err := rows.Scan(
			&task.ID,
			&userID,
			&categoryID,
			&task.CategoryName,
			&task.Description,
			&task.DueDate,
			&task.StartDate,
			&task.FinishDate,
			&task.Status,
			&task.Notes,
		)
		if err != nil {
			return nil, err
		}
		if categoryID.Valid {
			task.CategoryID = int(categoryID.Int64)
		}
		if userID.Valid {
			task.UserID = int(userID.Int64)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func Select_Initial_Tasks() ([]models.Task, error) {
	stmt, err := DB.Prepare(`
		SELECT t.id, t.user_id, t.category_id, COALESCE(c.name, '') as category_name,
		       t.task_description, t.due_date, t.start_date, t.finish_date, t.status, t.notes
		FROM initial_tasks t
		LEFT JOIN categories c ON t.category_id = c.id
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		var categoryID sql.NullInt64
		var userID sql.NullInt64
		err := rows.Scan(
			&task.ID,
			&userID,
			&categoryID,
			&task.CategoryName,
			&task.Description,
			&task.DueDate,
			&task.StartDate,
			&task.FinishDate,
			&task.Status,
			&task.Notes,
		)
		if err != nil {
			return nil, err
		}
		if categoryID.Valid {
			task.CategoryID = int(categoryID.Int64)
		}
		if userID.Valid {
			task.UserID = int(userID.Int64)
		}
		task.OriginalID = 0 // Initial tasks don't have an original ID
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func Select_Initial_Tasks_By_Search(search string, userID int) ([]models.Task, error) {
	stmt, err := DB.Prepare(`
		SELECT t.id, t.user_id, t.category_id, COALESCE(c.name, '') as category_name,
		       t.task_description, t.due_date, t.start_date, t.finish_date, t.status, t.notes
		FROM initial_tasks t
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE t.task_description LIKE ?
		AND t.user_id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query("%" + search + "%", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		var categoryID sql.NullInt64
		var userID sql.NullInt64
		err := rows.Scan(
			&task.ID,
			&userID,
			&categoryID,
			&task.CategoryName,
			&task.Description,
			&task.DueDate,
			&task.StartDate,
			&task.FinishDate,
			&task.Status,
			&task.Notes,
		)
		if err != nil {
			return nil, err
		}
		if categoryID.Valid {
			task.CategoryID = int(categoryID.Int64)
		}
		if userID.Valid {
			task.UserID = int(userID.Int64)
		}
		task.OriginalID = 0 // Initial tasks don't have an original ID
		tasks = append(tasks, task)
	}

	return tasks, nil
}
// MoveCompletedTask moves a task from initial_tasks to completed_tasks when status is "complete"
func MoveCompletedTask(tks *models.Task) error {
	completedStatus := "Completed"

	// Handle null category_id
	var categoryID interface{}
	if tks.CategoryID == 0 {
		categoryID = nil
	} else {
		categoryID = tks.CategoryID
	}

	stmt, err := DB.Prepare("INSERT INTO completed_tasks (user_id, category_id, task_id, task_description, due_date, start_date, finish_date, status, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tks.UserID, categoryID, tks.ID, tks.Description, tks.DueDate, tks.StartDate, tks.FinishDate, completedStatus, tks.Notes)
	if err != nil {
		return err
	}

	// Delete from initial_tasks table
	err = DeleteData(tks)
	if err != nil {
		return err
	}

	return nil
}

// SelectCompletedTasks retrieves all completed tasks
func SelectCompletedTasks() ([]models.Task, error) {
	stmt, err := DB.Prepare(`
		SELECT t.id, t.user_id, t.category_id, COALESCE(c.name, '') as category_name,
		       t.task_id, t.task_description, t.due_date, t.start_date, t.finish_date, t.status, t.notes
		FROM completed_tasks t
		LEFT JOIN categories c ON t.category_id = c.id
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		var categoryID sql.NullInt64
		var userID sql.NullInt64
		var originalTaskID sql.NullInt64
		err := rows.Scan(
			&task.ID,
			&userID,
			&categoryID,
			&task.CategoryName,
			&originalTaskID,
			&task.Description,
			&task.DueDate,
			&task.StartDate,
			&task.FinishDate,
			&task.Status,
			&task.Notes,
		)
		if err != nil {
			return nil, err
		}
		if categoryID.Valid {
			task.CategoryID = int(categoryID.Int64)
		}
		if userID.Valid {
			task.UserID = int(userID.Int64)
		}
		if originalTaskID.Valid {
			task.OriginalID = int(originalTaskID.Int64)
		} else {
			task.OriginalID = 0
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func Select_Initial_Tasks_By_Category(categoryID int) ([]models.Task, error) {
	stmt, err := DB.Prepare(`
		SELECT t.id, t.user_id, t.category_id, COALESCE(c.name, '') as category_name,
		       t.task_description, t.due_date, t.start_date, t.finish_date, t.status, t.notes
		FROM initial_tasks t
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE t.category_id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		var catID sql.NullInt64
		var userID sql.NullInt64
		err := rows.Scan(
			&task.ID,
			&userID,
			&catID,
			&task.CategoryName,
			&task.Description,
			&task.DueDate,
			&task.StartDate,
			&task.FinishDate,
			&task.Status,
			&task.Notes,
		)
		if err != nil {
			return nil, err
		}
		if catID.Valid {
			task.CategoryID = int(catID.Int64)
		}
		if userID.Valid {
			task.UserID = int(userID.Int64)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
