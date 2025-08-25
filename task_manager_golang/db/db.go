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
	fmt.Println("Reading schema from:", absPath)

	if _, err := os.Stat(schemaFullPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("schema file does not exist: %s", schemaFullPath)
		}
		return fmt.Errorf("error checking schema file: %w", err)
	}

	schema, err := os.ReadFile(schemaFullPath)
	if err != nil {
		return fmt.Errorf("reading schema: %w", err)
	}

	_, err = DB.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("executing schema: %w", err)
	}

	return nil
}
