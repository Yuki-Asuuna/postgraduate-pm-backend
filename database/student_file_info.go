package database

import (
	"github.com/sirupsen/logrus"
	"postgraduate-pm-backend/constant"
	"postgraduate-pm-backend/utils/mysql"
)

func GetStudentFileInfoByIdentityNumber(identityNumber string) (*StudentFileInfo, error) {
	info := new(StudentFileInfo)
	if err := mysql.GetMySQLClient().First(info, "identity_number = ?", identityNumber).Error; err != nil {
		logrus.Errorf(constant.DAO+"GetStudentFileInfoByIdentityNumber Failed, err= %v", err)
		return nil, err
	}
	return info, nil
}

func UpdateFirstDraftByIdentityNumber(identityNumber string, firstDraft string) error {
	err := mysql.GetMySQLClient().Model(&StudentFileInfo{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"is_first_draft_confirmed": false,
		"is_first_draft_submitted": true,
		"first_draft":              firstDraft,
	}).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateFirstDraftByIdentityNumber Failed, err= %v", err)
		return err
	}
	return nil
}

func UpdatePreliminaryReviewFormByIdentityNumber(identityNumber string, preliminaryReviewForm string) error {
	err := mysql.GetMySQLClient().Model(&StudentFileInfo{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"is_preliminary_review_form_confirmed": false,
		"is_preliminary_review_form_submitted": true,
		"preliminary_review_form":              preliminaryReviewForm,
	}).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdatePreliminaryReviewFormByIdentityNumber Failed, err= %v", err)
		return err
	}
	return nil
}

func UpdateResearchEvaluationMaterialByIdentityNumber(identityNumber string, researchEvaluationMaterial string) error {
	err := mysql.GetMySQLClient().Model(&StudentFileInfo{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"is_research_evaluation_material_confirmed": false,
		"is_research_evaluation_material_submitted": true,
		"research_evaluation_material":              researchEvaluationMaterial,
	}).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"UpdateResearchEvaluationMaterialByIdentityNumber Failed, err= %v", err)
		return err
	}
	return nil
}

func UpdateStudentCommentByIdentityNumber(identityNumber string, studentComment string) error {
	return mysql.GetMySQLClient().Model(&StudentFileInfo{}).Where("identity_number = ?", identityNumber).Update("student_comment", studentComment).Error
}

func UpdateSupervisorCommentByIdentityNumber(identityNumber string, supervisorComment string) error {
	return mysql.GetMySQLClient().Model(&StudentFileInfo{}).Where("identity_number = ?", identityNumber).Update("supervisor_comment", supervisorComment).Error
}

func GetStudentFileInfosByIdentityNumbers(identityNumbers []string) (map[string]*StudentFileInfo, error) {
	infos := make([]*StudentFileInfo, 0)
	if err := mysql.GetMySQLClient().Where("identity_number in (?)", identityNumbers).Find(&infos).Error; err != nil {
		logrus.Errorf(constant.DAO+"GetStudentFileInfosByIdentityNumbers Failed, err= %v", err)
		return nil, err
	}
	infoMap := make(map[string]*StudentFileInfo)
	for _, info := range infos {
		infoMap[info.IdentityNumber] = info
	}
	return infoMap, nil
}

func UpdateIsFirstDraftConfirmedByIdentityNumber(identityNumber string, isFirstDraftConfirmed int64) error {
	if err := mysql.GetMySQLClient().Model(&StudentFileInfo{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"is_first_draft_confirmed": isFirstDraftConfirmed,
	}).Error; err != nil {
		logrus.Errorf(constant.DAO+"UpdateIsFirstDraftConfirmedByIdentityNumber Failed, err= %v", err)
		return err
	}
	return nil
}

func UpdateIsPreliminaryReviewFormConfirmedByIdentityNumber(identityNumber string, isPreliminaryReviewFormConfirmed int64) error {
	if err := mysql.GetMySQLClient().Model(&StudentFileInfo{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"is_preliminary_review_form_confirmed": isPreliminaryReviewFormConfirmed,
	}).Error; err != nil {
		logrus.Errorf(constant.DAO+"UpdateIsFirstDraftConfirmedByIdentityNumber Failed, err= %v", err)
		return err
	}
	return nil
}

func UpdateIsResearchEvaluationMaterialConfirmedByIdentityNumber(identityNumber string, isResearchEvaluationMaterialConfirmed int64) error {
	if err := mysql.GetMySQLClient().Model(&StudentFileInfo{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"is_research_evaluation_material_confirmed": isResearchEvaluationMaterialConfirmed,
	}).Error; err != nil {
		logrus.Errorf(constant.DAO+"UpdateIsFirstDraftConfirmedByIdentityNumber Failed, err= %v", err)
		return err
	}
	return nil

}
