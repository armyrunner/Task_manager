package db

import (
	"github.com/armyrunner/task_manager/models"
)

func InsertData(description, due_date, start_date, finish_date, status, notes string) error {
	stmt, err := DB.Prepare("INSERT INTO initial_tasks (task_description, due_date, start_date, finish_date, status, notes) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(description, due_date, start_date, finish_date, status, notes)
	if err != nil {
		return err
	}

	return nil
}

func UpdateData(id int, description, due_date, start_date, finish_date, status, notes string) error {
	stmt, err := DB.Prepare("UPDATE initial_tasks SET task_description = ?, due_date = ?, start_date = ?, finish_date = ?, status = ?, notes = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(description, due_date, start_date, finish_date, status, notes, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteData(id int) error {
	stmt, err := DB.Prepare("DELETE FROM initial_tasks WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func SelectData(id int) (string, string, string, string, string, string, error) {
	stmt, err := DB.Prepare("SELECT task_description, due_date, start_date, finish_date, status, notes FROM initial_tasks WHERE id = ?")
	if err != nil {
		return "", "", "", "", "", "", err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return "", "", "", "", "", "", err
	}
	defer rows.Close()

	rows.Next()

	var description, due_date, start_date, finish_date, status, notes string
	err = rows.Scan(&description, &due_date, &start_date, &finish_date, &status, &notes)
	if err != nil {
		return "", "", "", "", "", "", err
	}

	return description, due_date, start_date, finish_date, status, notes, nil
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
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// MoveCompletedTask moves a task from initial_tasks to completed_tasks when status is "complete"
func MoveCompletedTask(id int) error {
	// First, get the task data
	description, due_date, start_date, finish_date, status, notes, err := SelectData(id)
	if err != nil {
		return err
	}

	// Insert into completed_tasks table
	stmt, err := DB.Prepare("INSERT INTO completed_tasks (task_id, task_description, due_date, start_date, finish_date, status, notes) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, description, due_date, start_date, finish_date, status, notes)
	if err != nil {
		return err
	}

	// Delete from initial_tasks table
	err = DeleteData(id)
	if err != nil {
		return err
	}

	return nil
}

// SelectCompletedTasks retrieves all completed tasks
func SelectCompletedTasks() ([]models.Task, error) {
	stmt, err := DB.Prepare(`
		SELECT COALESCE(task_id, 0), task_description, due_date, start_date, finish_date, status, notes
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


func Select_Completed_Tasks() ([]models.Task, error) {
	stmt, err := DB.Prepare(`
		SELECT id, task_description, due_date, start_date, finish_date, status, notes
		FROM completed_tasks
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
		tasks = append(tasks, task)
	}

	return tasks, nil
}