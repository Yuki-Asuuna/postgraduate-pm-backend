package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"postgraduate-pm-backend/utils/zookeeper"
)

var client *gorm.DB

type MYSQLConfig struct {
	MySQLServerIP string `json:"mysqlServerIP"`
	Port          int    `json:"port"`
	DBUsername    string `json:"dbUsername"`
	DBPassword    string `json:"dbPassword"`
	DBName        string `json:"dbName"`
	Charset       string `json:"charset"`
}

var config *MYSQLConfig

func MysqlInit() error {
	var err error
	config = &MYSQLConfig{}
	err = zookeeper.GetUtilsConfig("/mysql", config)
	if err != nil {
		return err
	}
	//Example:
	//root:%ecnu#0006$@(150.158.159.26:3306)/graduate_exemption?charset=utf8mb4&parseTime=True&loc=Local
	target := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.MySQLServerIP, config.Port, config.DBName, config.Charset)
	client, err = gorm.Open("mysql", target)
	if err != nil {
		return err
	}
	return nil
}

func GetMySQLClient() *gorm.DB {
	return client
}
