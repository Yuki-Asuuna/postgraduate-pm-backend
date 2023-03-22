package database

import (
	"github.com/sirupsen/logrus"
	"postgraduate-pm-backend/constant"
	"postgraduate-pm-backend/utils/mysql"
)

func GetStudentFileInfoByIdentityNumber(identityNumber string) (*StudentFileInfo, error) {
	info := new(StudentFileInfo)
	if err := mysql.GetMySQLClient().First(info, "identity_number = ?", identityNumber).Error; err != nil {
		logrus.Error(constant.DAO+"GetStudentFileInfoByIdentityNumber Failed, err= %v", err)
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
