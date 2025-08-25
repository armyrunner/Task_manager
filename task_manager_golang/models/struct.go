package models

import "time"


type Task struct{

	ID          int
	Description string
	DueDate     time.Time
	StartDate   *time.Time
	FinishDate  *time.Time
	Status      string
	Notes       string
}