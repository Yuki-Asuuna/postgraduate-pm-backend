package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"postgraduate-pm-backend/api"
	"postgraduate-pm-backend/constant"
	"postgraduate-pm-backend/database"
	"postgraduate-pm-backend/exception"
	"postgraduate-pm-backend/utils"
	"postgraduate-pm-backend/utils/minio"
	"postgraduate-pm-backend/utils/sessions"
	"time"
)

func FirstDraftUpload(c *gin.Context) {
	f, err := c.FormFile("source")
	if err != nil {
		logrus.Error(constant.Service+"FirstDraftUpload Failed, err= %v", err)
		c.Error(exception.ParameterError())
		return
	}
	url, err := minio.UploadFile("first-draft", f)
	if err != nil {
		logrus.Error(constant.Service+"FirstDraftUpload Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	err = database.UpdateFirstDraftByIdentityNumber(sessions.GetUserInfoBySession(c).IdentityNumber, url)
	if err != nil {
		logrus.Error(constant.Service+"FirstDraftUpload Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", url))
}

func PreliminaryReviewFormUpload(c *gin.Context) {
	f, err := c.FormFile("source")
	if err != nil {
		logrus.Error(constant.Service+"PreliminaryReviewFormUpload Failed, err= %v", err)
		c.Error(exception.ParameterError())
		return
	}
	url, err := minio.UploadFile("preliminary-review-form", f)
	if err != nil {
		logrus.Error(constant.Service+"PreliminaryReviewFormUpload Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	err = database.UpdatePreliminaryReviewFormByIdentityNumber(sessions.GetUserInfoBySession(c).IdentityNumber, url)
	if err != nil {
		logrus.Error(constant.Service+"PreliminaryReviewFormUpload Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", url))
}

func GetStudentStatusInfo(c *gin.Context) {
	info, err := database.GetStudentStatusInfoByIdentityNumber(sessions.GetUserInfoBySession(c).IdentityNumber)
	if err != nil {
		logrus.Error(constant.Service+"Get Student Status Info Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	supervisor, err := database.GetUserByIdentityNumber(info.SupervisorID)
	if err != nil {
		logrus.Error(constant.Service+"Get Student Status Info Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	result := &api.StudentStatusInfoResponse{
		IdentityNumber: info.IdentityNumber,
		College:        info.College,
		Class:          info.Class,
		Length:         info.Length,
		GraduateTime:   info.GraduateTime.Unix(),
		DegreeType:     info.DegreeType,
		Status:         info.Status,
		SupervisorName: supervisor.Name,
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", result))
}

func PostStudentStatusInfo(c *gin.Context) {
	user := sessions.GetUserInfoBySession(c)
	if user == nil {
		c.Error(exception.ServerError())
		logrus.Error(constant.Service + "Get Student Status Info Failed, user is nil")
		return
	}
	params := make(map[string]interface{})
	c.BindJSON(&params)
	college := params["college"].(string)
	class := params["class"].(string)
	length := int64(params["length"].(float64))
	identityNumber := params["identityNumber"].(string)
	degreeType := int64(params["degreeType"].(float64))
	status := int64(params["status"].(float64))
	graduateTime := int64(params["graduateTime"].(float64))

	err := database.UpdateStudentStatusInfoByIdentityNumber(identityNumber, college, class, length, degreeType, status, time.Unix(graduateTime, 0))
	if err != nil {
		logrus.Error(constant.Service+"Me Post Student Status Info Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func GetStudentFileInfo(c *gin.Context) {
	info, err := database.GetStudentFileInfoByIdentityNumber(sessions.GetUserInfoBySession(c).IdentityNumber)
	if err != nil {
		logrus.Error(constant.Service+"Get Student File Info Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	result := &api.StudentFileInfoResponse{
		IdentityNumber:                   info.IdentityNumber,
		FirstDraft:                       info.FirstDraft,
		PreliminaryReviewForm:            info.PreliminaryReviewForm,
		IsFirstDraftConfirmed:            info.IsFirstDraftConfirmed,
		IsPreliminaryReviewFormConfirmed: info.IsPreliminaryReviewFormConfirmed,
		IsFirstDraftSubmitted:            info.IsFirstDraftSubmitted,
		IsPreliminaryReviewFormSubmitted: info.IsPreliminaryReviewFormSubmitted,
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", result))
}

func StudentGetComment(c *gin.Context) {
	identityNumber := sessions.GetUserIdentityNumberBySession(c)
	info, err := database.GetStudentFileInfoByIdentityNumber(identityNumber)
	if err != nil {
		logrus.Error(constant.Service+"StudentGetComment Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	result := &api.StudentGetCommentResponse{
		IdentityNumber:    info.IdentityNumber,
		StudentComment:    info.StudentComment,
		SupervisorComment: info.SupervisorComment,
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", result))
}

func StudentPostComment(c *gin.Context) {
	identityNumber := sessions.GetUserIdentityNumberBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	studentComment := params["studentComment"].(string)
	err := database.UpdateStudentCommentByIdentityNumber(identityNumber, studentComment)
	if err != nil {
		logrus.Error(constant.Service+"StudentPostComment Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}
