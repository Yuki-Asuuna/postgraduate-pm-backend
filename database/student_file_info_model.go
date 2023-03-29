package database

type StudentFileInfo struct {
	IdentityNumber                        string
	FirstDraft                            string
	PreliminaryReviewForm                 string
	ResearchEvaluationMaterial            string
	IsFirstDraftSubmitted                 int64
	IsPreliminaryReviewFormSubmitted      int64
	IsFirstDraftConfirmed                 int64
	IsPreliminaryReviewFormConfirmed      int64
	IsResearchEvaluationMaterialSubmitted int64
	IsResearchEvaluationMaterialConfirmed int64
	StudentComment                        string
	SupervisorComment                     string
}

func (StudentFileInfo) TableName() string {
	return "student_file_info"
}
