package db

import (
	"database/sql"
	//"time"

	"github.com/armyrunner/task_manager/models"
)

func InsertData(tks models.Task) error {
	stmt, err := DB.Prepare("INSERT INTO initial_tasks (task_description, due_date, start_date, finish_date, status, notes) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tks.Description, tks.DueDate, tks.StartDate, tks.FinishDate, tks.Status, tks.Notes)
	if err != nil {
		return err
	}

	return nil
}

func UpdateData(tks models.Task) error {
	stmt, err := DB.Prepare("UPDATE initial_tasks SET task_description = ?, due_date = ?, start_date = ?, finish_date = ?, status = ?, notes = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tks.Description, tks.DueDate, tks.StartDate, tks.FinishDate, tks.Status, tks.Notes, tks.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteData(tks models.Task) error {
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

func SelectData(tks models.Task) ([]models.Task, error) {
	stmt, err := DB.Prepare("SELECT id, task_description, due_date, start_date, finish_date, status, notes FROM initial_tasks WHERE id = ?")
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
		err := rows.Scan(
			&task.ID,
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
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func Select_Initial_Tasks() ([]models.Task, error) {
	stmt, err := DB.Prepare(`
		SELECT id, task_description, due_date, start_date, finish_date, status, notes
		FROM initial_tasks
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
		err := rows.Scan(
			&task.ID,
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
		task.OriginalID = 0 // Initial tasks don't have an original ID
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// MoveCompletedTask moves a task from initial_tasks to completed_tasks when status is "complete"
func MoveCompletedTask(tks models.Task) error {
	// First, get the task data
	tasks, err := SelectData(tks)
	if err != nil {
		return err
	}

	// Update the status and finish date to reflect completion
	completedStatus := "Completed"
	//completedFinishDate := getCurrentDate()

	// Insert into completed_tasks table
	stmt, err := DB.Prepare("INSERT INTO completed_tasks (task_id, task_description, due_date, start_date, finish_date, status, notes) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tasks[0].OriginalID, tasks[0].Description, tasks[0].DueDate, tasks[0].StartDate, tasks[0].FinishDate, completedStatus, tasks[0].Notes)
	if err != nil {
		return err
	}

	// Delete from initial_tasks table
	err = DeleteData(tasks[0])
	if err != nil {
		return err
	}

	return nil
}

// getCurrentDate returns the current date in YYYY-MM-DD format
// func getCurrentDate() string {
// 	now := time.Now()
// 	return now.Format("2006-01-02")
// }

// SelectCompletedTasks retrieves all completed tasks
func SelectCompletedTasks() ([]models.Task, error) {
	stmt, err := DB.Prepare(`
		SELECT id, task_id, task_description, due_date, start_date, finish_date, status, notes
		FROM completed_tasks
		ORDER BY id DESC
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
		var originalTaskID sql.NullInt64
		err := rows.Scan(
			&task.ID,
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
		if originalTaskID.Valid {
			task.OriginalID = int(originalTaskID.Int64)
		} else {
			task.OriginalID = 0
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
