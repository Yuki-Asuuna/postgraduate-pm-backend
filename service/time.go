package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"postgraduate-pm-backend/constant"
	"postgraduate-pm-backend/utils/redis"
)

func GetCurrentTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"result":  redis.GetCurrentTime(),
	})
}

func PostCurrentTime(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	err := redis.SetCurrentTime(int64(params["current_time"].(float64)))
	if err != nil {
		logrus.Errorf(constant.Service+"PostCurrentTime failed: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
	})
}
