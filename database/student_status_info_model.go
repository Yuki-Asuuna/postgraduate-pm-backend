package database

import "time"

type StudentStatusInfo struct {
	IdentityNumber string
	College        string
	Class          string
	Length         int64
	GraduateTime   time.Time
	DegreeType     int64
	Status         int64
	SupervisorID   string
	IsConfirmed    int64
}

func (StudentStatusInfo) TableName() string {
	return "student_status_info"
}
