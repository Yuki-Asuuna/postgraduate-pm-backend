package database

import (
	"github.com/sirupsen/logrus"
	"postgraduate-pm-backend/constant"
	"postgraduate-pm-backend/utils/mysql"
	"time"
)

func GetStudentStatusInfoByIdentityNumber(identityNumber string) (*StudentStatusInfo, error) {
	info := new(StudentStatusInfo)
	if err := mysql.GetMySQLClient().First(info, "identity_number = ?", identityNumber).Error; err != nil {
		logrus.Errorf(constant.DAO+"GetStudentStatusInfoByIdentityNumber Failed, err= %v", err)
		return nil, err
	}
	return info, nil
}

func UpdateStudentStatusInfoByIdentityNumber(identityNumber string, college string, class string, length int64, degreeType int64, status int64, graduateTime time.Time, isConfirmed int64) error {
	if err := mysql.GetMySQLClient().Model(&StudentStatusInfo{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"college":       college,
		"class":         class,
		"length":        length,
		"degree_type":   degreeType,
		"status":        status,
		"graduate_time": graduateTime,
		"is_confirmed":  isConfirmed,
	}).Error; err != nil {
		logrus.Errorf(constant.DAO+"UpdateStudentStatusInfoByIdentityNumber Failed, err= %v", err)
		return err
	}
	return nil
}

func GetStudentStatusInfoListBySupervisorID(supervisorID string, page int, size int) ([]*StudentStatusInfo, error) {
	var studentStatusInfoList []*StudentStatusInfo
	studentStatusInfoList = make([]*StudentStatusInfo, 0)
	query := mysql.GetMySQLClient()
	query = query.Where("supervisor_id = ?", supervisorID).Offset(page * size).Limit(size)
	if err := query.Find(&studentStatusInfoList).Error; err != nil {
		logrus.Errorf(constant.DAO+"GetStudentStatusInfoListBySupervisorID Failed, err= %v", err)
		return nil, err
	}
	return studentStatusInfoList, nil
}

func UpdateSupervisorIDByIdentityNumber(studentID, identityNumber string) error {
	if err := mysql.GetMySQLClient().Model(&StudentStatusInfo{}).Where("identity_number = ?", studentID).Updates(map[string]interface{}{
		"supervisor_id": identityNumber,
	}).Error; err != nil {
		logrus.Errorf(constant.DAO+"UpdateSupervisorIDByIdentityNumber Failed, err= %v", err)
		return err
	}
	return nil
}

func GetStudentStatusInfoList(page int, size int, identityNumber string) ([]*StudentStatusInfo, error) {
	counsellors := make([]*StudentStatusInfo, 0)
	query := mysql.GetMySQLClient()
	if identityNumber != "" {
		query = query.Where("identity_number like ?", "%"+identityNumber+"%")
	}
	err := query.Offset(page * size).Limit(size).Find(&counsellors).Error
	if err != nil {
		logrus.Errorf(constant.DAO+"GetStudentStatusInfoList Failed, err= %v", err)
		return nil, err
	}
	return counsellors, nil
}

func UpdateBlindScore(identityNumber string, blindScore int64) error {
	if err := mysql.GetMySQLClient().Model(&StudentStatusInfo{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"blind_score": blindScore,
	}).Error; err != nil {
		logrus.Errorf(constant.DAO+"UpdateBlindScore Failed, err= %v", err)
		return err
	}
	return nil
}

func UpdateDefenseScore(identityNumber string, defenseScore int64) error {
	if err := mysql.GetMySQLClient().Model(&StudentStatusInfo{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"defense_score": defenseScore,
	}).Error; err != nil {
		logrus.Errorf(constant.DAO+"UpdateDefenseScore Failed, err= %v", err)
		return err
	}
	return nil
}

func UpdateDegreeConfirmed(identityNumber string, degreeConfirmed int64) error {
	if err := mysql.GetMySQLClient().Model(&StudentStatusInfo{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"degree_confirmed": degreeConfirmed,
	}).Error; err != nil {
		logrus.Errorf(constant.DAO+"UpdateDegreeConfirmed Failed, err= %v", err)
		return err
	}
	return nil
}

func UpdateApplyDegree(identityNumber string, applyDegree int64) error {
	if err := mysql.GetMySQLClient().Model(&StudentStatusInfo{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"apply_degree": applyDegree,
	}).Error; err != nil {
		logrus.Errorf(constant.DAO+"UpdateApplyDegree Failed, err= %v", err)
		return err
	}
	return nil
}
