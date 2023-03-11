package main

import (
	"github.com/sirupsen/logrus"
	"postgraduate-pm-backend/constant"
	"postgraduate-pm-backend/middleware"
	"postgraduate-pm-backend/service"
)

func httpHandlerInit() {
	logrus.Info(constant.Main + "Init httpHandlerInit")
	// 支持跨域访问
	r.Use(middleware.Cors())

	r.GET("/ping", service.Ping)
}
