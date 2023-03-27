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
	return mysql.GetMySQLClient().Model(&StudentFileInfo{}).Where("identity_number = ?", identityNumber).Update("first_draft", firstDraft).Error
}

func UpdatePreliminaryReviewFormByIdentityNumber(identityNumber string, preliminaryReviewForm string) error {
	return mysql.GetMySQLClient().Model(&StudentFileInfo{}).Where("identity_number = ?", identityNumber).Update("preliminary_review_form", preliminaryReviewForm).Error
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
