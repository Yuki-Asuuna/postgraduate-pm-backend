package api

type AdminGetStudentListResponse struct {
	Stus []*StudentStatusInfo `json:"stus"`
}
