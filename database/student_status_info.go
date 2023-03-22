package database

import (
	"postgraduate-pm-backend/utils/mysql"
	"time"
)

func GetStudentStatusInfoByIdentityNumber(identityNumber string) (*StudentStatusInfo, error) {
	info := new(StudentStatusInfo)
	if err := mysql.GetMySQLClient().First(info, "identity_number = ?", identityNumber).Error; err != nil {
		return nil, err
	}
	return info, nil
}

func UpdateStudentStatusInfoByIdentityNumber(identityNumber string, college string, class string, length int64, degreeType int64, status int64, graduateTime time.Time) error {
	if err := mysql.GetMySQLClient().Model(&StudentStatusInfo{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"college":       college,
		"class":         class,
		"length":        length,
		"degree_type":   degreeType,
		"status":        status,
		"graduate_time": graduateTime,
	}).Error; err != nil {
		return err
	}
	return nil
}
