package database

type StudentFileInfo struct {
	IdentityNumber                   string
	FirstDraft                       string
	PreliminaryReviewForm            string
	IsFirstDraftSubmitted            int64
	IsPreliminaryReviewFormSubmitted int64
	IsFirstDraftConfirmed            int64
	IsPreliminaryReviewFormConfirmed int64
}

func (StudentFileInfo) TableName() string {
	return "student_file_info"
}
