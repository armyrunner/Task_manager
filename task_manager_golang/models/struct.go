package models

type Task struct {
	ID          int
	OriginalID  int
	Description string
	DueDate     string
	StartDate   string
	FinishDate  string
	Status      string
	Notes       string
	Category    string
}
