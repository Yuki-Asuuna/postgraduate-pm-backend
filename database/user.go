package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"postgraduate-pm-backend/constant"
	"postgraduate-pm-backend/utils/mysql"
	"time"
)

func GetUserByIdentityNumber(identityNumber string) (*User, error) {
	user := new(User)
	if err := mysql.GetMySQLClient().First(user, "identity_number = ?", identityNumber).Error; err != nil && err != gorm.ErrRecordNotFound {
		logrus.Errorf(constant.DAO+"GetUserByIdentityNumber Failed, err= %v", err)
		return nil, err
	}
	return user, nil
}

func UpdateLoginTimeByIdentityNumber(identityNumber string) error {
	return mysql.GetMySQLClient().Model(&User{}).Where("identity_number = ?", identityNumber).Update("last_login", time.Now()).Error
}

func UpdatePasswordByIdentityNumber(identityNumber string, password string) error {
	return mysql.GetMySQLClient().Model(&User{}).Where("identity_number = ?", identityNumber).Update("password", password).Error
}

func UpdateUserAvatarByIdentityNumber(identityNumber string, avatar string) error {
	return mysql.GetMySQLClient().Model(&User{}).Where("identity_number = ?", identityNumber).Update("avatar", avatar).Error
}

func UpdateUserByIdentityNumber(identityNumber string, name string, role int64, gender int64, age int64, phoneNumber string, email string) error {
	if err := mysql.GetMySQLClient().Model(&User{}).Where("identity_number = ?", identityNumber).Updates(map[string]interface{}{
		"name":         name,
		"role":         role,
		"gender":       gender,
		"age":          age,
		"phone_number": phoneNumber,
		"email":        email,
	}).Error; err != nil {
		logrus.Errorf(constant.DAO+"UpdateUserByIdentityNumber Failed, err= %v", err)
		return err
	}
	return nil
}

func GetUsersByIdentityNumbers(identityNumbers []string) (map[string]*User, error) {
	users := make([]*User, 0)
	if err := mysql.GetMySQLClient().Where("identity_number in (?)", identityNumbers).Find(&users).Error; err != nil {
		logrus.Errorf(constant.DAO+"GetUsersByIdentityNumbers Failed, err= %v", err)
		return nil, err
	}
	userMap := make(map[string]*User)
	for _, user := range users {
		userMap[user.IdentityNumber] = user
	}
	return userMap, nil
}
