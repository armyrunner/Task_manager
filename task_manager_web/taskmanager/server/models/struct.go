package models

type Task struct {
	ID          int		`json:"id"`
	OriginalID  int		`json:"originalid"`
	Description string	`json:"description"`
	DueDate     string	`json:"duedate"`
	StartDate   string	`json:"startdate"`
	FinishDate  string	`json:"finishdate"`
	Status      string	`json:"status"`
	Notes       string	`json:"notes"`
	Category    string	`json:"category"`
}
