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

	r.GET("/current_time", service.GetCurrentTime)
	r.POST("/current_time", service.PostCurrentTime)

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
		stuGroup.POST("/research_evaluation_material_upload", middleware.AuthMiddleWare(), service.ResearchEvaluationMaterialUpload)
		stuGroup.GET("/status_info", middleware.AuthMiddleWare(), service.GetStudentStatusInfo)
		stuGroup.POST("/status_info", middleware.AuthMiddleWare(), service.PostStudentStatusInfo)
		stuGroup.GET("/file_info", middleware.AuthMiddleWare(), service.GetStudentFileInfo)
		stuGroup.GET("/comment", middleware.AuthMiddleWare(), service.StudentGetComment)
		stuGroup.POST("/comment", middleware.AuthMiddleWare(), service.StudentPostComment)
	}

	supervisorGroup := r.Group("/supervisor")
	{
		supervisorGroup.GET("/stu_list", middleware.AuthMiddleWare(), service.SupervisorGetStudentList)
		supervisorGroup.GET("/comment", middleware.AuthMiddleWare(), service.SupervisorGetComment)
		supervisorGroup.POST("/comment", middleware.AuthMiddleWare(), service.SupervisorPostComment)
		supervisorGroup.POST("/bind", middleware.AuthMiddleWare(), service.SupervisorBindStudent)
		supervisorGroup.POST("/confirm", middleware.AuthMiddleWare(), service.SupervisorConfirmStudent)
	}

	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/stu_list", middleware.AuthMiddleWare(), service.AdminGetStudentList)
		adminGroup.POST("/upload_blind_score", middleware.AuthMiddleWare(), service.AdminUploadBlindScore)
		adminGroup.POST("/upload_defense_score", middleware.AuthMiddleWare(), service.AdminUploadDefenseScore)
	}
}
