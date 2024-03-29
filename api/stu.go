package api

type StudentStatusInfoResponse struct {
	IdentityNumber string `json:"identityNumber"`
	College        string `json:"college"`
	Class          string `json:"class"`
	Length         int64  `json:"length"`
	GraduateTime   int64  `json:"graduateTime"`
	DegreeType     int64  `json:"degreeType"`
	Status         int64  `json:"status"`
	SupervisorName string `json:"supervisorName"`
	IsConfirmed    int64  `json:"isConfirmed"`
}

type StudentFileInfoResponse struct {
	IdentityNumber                        string `json:"identityNumber"`
	FirstDraft                            string `json:"firstDraft"`
	PreliminaryReviewForm                 string `json:"preliminaryReviewForm"`
	ResearchEvaluationMaterial            string `json:"researchEvaluationMaterial"`
	IsFirstDraftSubmitted                 int64  `json:"isFirstDraftSubmitted"`
	IsPreliminaryReviewFormSubmitted      int64  `json:"isPreliminaryReviewFormSubmitted"`
	IsFirstDraftConfirmed                 int64  `json:"isFirstDraftConfirmed"`
	IsPreliminaryReviewFormConfirmed      int64  `json:"isPreliminaryReviewFormConfirmed"`
	IsResearchEvaluationMaterialSubmitted int64  `json:"isResearchEvaluationMaterialSubmitted"`
	IsResearchEvaluationMaterialConfirmed int64  `json:"isResearchEvaluationMaterialConfirmed"`
}

type StudentGetCommentResponse struct {
	IdentityNumber    string `json:"identityNumber"`
	StudentComment    string `json:"studentComment"`
	SupervisorComment string `json:"supervisorComment"`
}
