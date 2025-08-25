package db

import "fmt"

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

func SelectAllData() ([][]string, error) {
	stmt, err := DB.Prepare("SELECT id, task_description, due_date, start_date, finish_date, status, notes FROM initial_tasks")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results [][]string

	for rows.Next() {
		var id int
		var description, due_date, start_date, finish_date, status, notes string
		err = rows.Scan(&id, &description, &due_date, &start_date, &finish_date, &status, &notes)
		if err != nil {
			return nil, err
		}
		results = append(results, []string{
			fmt.Sprintf("%d", id),
			description,
			due_date,
			start_date,
			finish_date,
			status,
			notes,
		})
	}

	return results, nil
}
