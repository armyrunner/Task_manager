package cmd

// Shared variables for CLI flags across all commands
// These are populated by cobra flags in each command's init() function

// Task fields
var (
	task_description string
	due_date         string
	start_date       string
	finish_date      string
	status           string
	notes            string
	category         string
)

// Task identifiers
var (
	task_id string
)

