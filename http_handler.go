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

	r.PUT("/image_upload", service.ImageUpload)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", service.Login)
		authGroup.POST("/logout", service.Logout)
		authGroup.GET("/me", middleware.AuthMiddleWare(), service.Me)
		authGroup.POST("/me", middleware.AuthMiddleWare(), service.PostMe)
		authGroup.POST("/password", middleware.AuthMiddleWare(), service.ChangePassword)
		authGroup.POST("/avatar_upload", middleware.AuthMiddleWare(), service.AvatarUpload)
	}

	stuGroup := r.Group("/stu")
	{
		stuGroup.POST("/first_draft_upload", middleware.AuthMiddleWare(), service.FirstDraftUpload)
		stuGroup.POST("/preliminary_review_form_upload", middleware.AuthMiddleWare(), service.PreliminaryReviewFormUpload)
		stuGroup.GET("/status_info", middleware.AuthMiddleWare(), service.GetStudentStatusInfo)
		stuGroup.POST("/status_info", middleware.AuthMiddleWare(), service.PostStudentStatusInfo)
		stuGroup.GET("/file_info", middleware.AuthMiddleWare(), service.GetStudentFileInfo)
	}
}
