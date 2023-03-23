package api

type StudentStatusInfo struct {
	IdentityNumber string `json:"identityNumber"`
	Name           string `json:"name"`
	College        string `json:"college"`
	Class          string `json:"class"`
	Length         int64  `json:"length"`
	GraduateTime   string `json:"graduateTime"`
	DegreeType     int64  `json:"degreeType"`
	Status         int64  `json:"status"`
	SupervisorID   string `json:"supervisorID"`
	FirstDraftURL  string `json:"firstDraftURL"`
}

type SupervisorGetStudentListResponse struct {
	Stus []*StudentStatusInfo `json:"stus"`
}
