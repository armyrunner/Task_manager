package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func OpenDatabase(db_path, schema_filename string) error {

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting working dir: %w", err)
	}

	dbFullPath := filepath.Join(cwd, db_path)
	schemaFullPath := filepath.Join(cwd, schema_filename)

	options :=
		"?" + "_busy_timeout=10000" +
			"&" + "_case_sensitive_like=OFF" +
			"&" + "_foreign_keys=ON" +
			"&" + "_journal_mode=OFF" +
			"&" + "_locking_mode=NORMAL" +
			"&" + "mode=rwc" +
			"&" + "_synchronous=OFF"
	DB, err = sql.Open("sqlite3", dbFullPath+options)
	if err != nil {
		return fmt.Errorf("opening db: %w", err)
	}

	absPath, _ := filepath.Abs(schemaFullPath)

	if _, err := os.Stat(absPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("schema file does not exist: %s", absPath)
		}
		return fmt.Errorf("error checking schema file: %w", err)
	}

	schema, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("reading schema: %w", err)
	}

	_, err = DB.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("executing schema: %w", err)
	}

	// Run migrations for new columns
	if err := runMigrations(); err != nil {
		return fmt.Errorf("running migrations: %w", err)
	}

	return nil
}

// columnExists checks if a column exists in a table
func columnExists(tableName, columnName string) (bool, error) {
	query := fmt.Sprintf("PRAGMA table_info(%s)", tableName)
	rows, err := DB.Query(query)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var cid int
		var name, ctype string
		var notnull, pk int
		var dfltValue sql.NullString
		if err := rows.Scan(&cid, &name, &ctype, &notnull, &dfltValue, &pk); err != nil {
			return false, err
		}
		if name == columnName {
			return true, nil
		}
	}
	return false, nil
}

// addColumnIfNotExists adds a column to a table if it doesn't already exist
// Uses DEFAULT ” for TEXT columns to prevent NULL values
func addColumnIfNotExists(tableName, columnName, columnType string) error {
	exists, err := columnExists(tableName, columnName)
	if err != nil {
		return err
	}
	if !exists {
		// Add DEFAULT '' for TEXT columns to prevent NULLs
		defaultClause := ""
		if columnType == "TEXT" {
			defaultClause = " DEFAULT ''"
		}
		query := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s%s", tableName, columnName, columnType, defaultClause)
		_, err = DB.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

// runMigrations adds any new columns that don't exist yet
func runMigrations() error {
	// Add user_id column to initial_tasks if it doesn't exist
	if err := addColumnIfNotExists("initial_tasks", "user_id", "INTEGER"); err != nil {
		return err
	}
	// Add user_id column to completed_tasks if it doesn't exist
	if err := addColumnIfNotExists("completed_tasks", "user_id", "INTEGER"); err != nil {
		return err
	}

	// Add category_id column to initial_tasks if it doesn't exist
	if err := addColumnIfNotExists("initial_tasks", "category_id", "INTEGER"); err != nil {
		return err
	}
	// Add category_id column to completed_tasks if it doesn't exist
	if err := addColumnIfNotExists("completed_tasks", "category_id", "INTEGER"); err != nil {
		return err
	}

	return nil
}
