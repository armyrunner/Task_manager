package pdf

import (
	"database/sql"

)

type ReportService struct {
	db *sql.DB
}