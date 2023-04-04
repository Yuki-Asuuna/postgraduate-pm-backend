package api

type StudentStatusInfo struct {
	IdentityNumber                        string `json:"identityNumber"`
	Name                                  string `json:"name"`
	College                               string `json:"college"`
	Class                                 string `json:"class"`
	Length                                int64  `json:"length"`
	GraduateTime                          string `json:"graduateTime"`
	DegreeType                            int64  `json:"degreeType"`
	Status                                int64  `json:"status"`
	SupervisorID                          string `json:"supervisorID"`
	FirstDraftURL                         string `json:"firstDraftURL"`
	IsFirstDraftConfirmed                 bool   `json:"isFirstDraftConfirmed"`
	PreliminaryReviewFormURL              string `json:"preliminaryReviewFormURL"`
	IsPreliminaryReviewFormConfirmed      bool   `json:"isPreliminaryReviewFormConfirmed"`
	ResearchEvaluationMaterialURL         string `json:"researchEvaluationMaterialURL"`
	IsResearchEvaluationMaterialConfirmed bool   `json:"isResearchEvaluationMaterialConfirmed"`
	BlindScore                            int64  `json:"blindScore"`
	DefenseScore                          int64  `json:"defenseScore"`
	ApplyDegree                           bool   `json:"applyDegree"`
	DegreeConfirmed                       bool   `json:"degreeConfirmed"`
}

type SupervisorGetStudentListResponse struct {
	Stus []*StudentStatusInfo `json:"stus"`
}

type SupervisorGetCommentResponse struct {
	IdentityNumber    string `json:"identityNumber"`
	StudentComment    string `json:"studentComment"`
	SupervisorComment string `json:"supervisorComment"`
}
