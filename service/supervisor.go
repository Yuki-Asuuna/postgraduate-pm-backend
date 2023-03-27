package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"postgraduate-pm-backend/api"
	"postgraduate-pm-backend/constant"
	"postgraduate-pm-backend/database"
	"postgraduate-pm-backend/utils"
	"postgraduate-pm-backend/utils/sessions"
)

func SupervisorGetComment(c *gin.Context) {
	supervisorIdentityNumber := sessions.GetUserIdentityNumberBySession(c)
	identityNumber := c.Query("identityNumber")
	fileInfo, err := database.GetStudentFileInfoByIdentityNumber(identityNumber)
	if err != nil {
		logrus.Errorf(constant.Service+"SupervisorGetComment Failed, err= %v", err)
		return
	}

	// 检查student是否属于当前supervisor
	stus, err := database.GetStudentStatusInfoListBySupervisorID(supervisorIdentityNumber)
	if err != nil {
		logrus.Errorf(constant.Service+"SupervisorGetStudentList Failed, err= %v", err)
		return
	}
	var flag bool = false
	for _, stu := range stus {
		if stu.IdentityNumber == identityNumber {
			flag = true
			break
		}
	}
	if !flag {
		c.JSON(http.StatusUnauthorized, utils.GenSuccessResponse(1, "student not belong to current supervisor", nil))
		return
	}

	result := &api.SupervisorGetCommentResponse{
		IdentityNumber:    fileInfo.IdentityNumber,
		StudentComment:    fileInfo.StudentComment,
		SupervisorComment: fileInfo.SupervisorComment,
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", result))
}

func SupervisorPostComment(c *gin.Context) {
	supervisorIdentityNumber := sessions.GetUserIdentityNumberBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	supervisorComment := params["supervisorComment"].(string)
	studentID := params["studentID"].(string)
	// 检查student是否属于当前supervisor
	stus, err := database.GetStudentStatusInfoListBySupervisorID(supervisorIdentityNumber)
	if err != nil {
		logrus.Errorf(constant.Service+"SupervisorGetStudentList Failed, err= %v", err)
		return
	}
	var flag bool = false
	for _, stu := range stus {
		if stu.IdentityNumber == studentID {
			flag = true
			break
		}
	}
	if !flag {
		c.JSON(http.StatusUnauthorized, utils.GenSuccessResponse(1, "student not belong to current supervisor", nil))
		return
	}

	err = database.UpdateSupervisorCommentByIdentityNumber(studentID, supervisorComment)
	if err != nil {
		logrus.Errorf(constant.Service+"SupervisorPostComment Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func SupervisorGetStudentList(c *gin.Context) {
	identityNumber := sessions.GetUserIdentityNumberBySession(c)
	stus, err := database.GetStudentStatusInfoListBySupervisorID(identityNumber)
	if err != nil {
		logrus.Errorf(constant.Service+"SupervisorGetStudentList Failed, err= %v", err)
		return
	}
	IDs := make([]string, 0)
	for _, stu := range stus {
		IDs = append(IDs, stu.IdentityNumber)
	}
	UserMap, err := database.GetUsersByIdentityNumbers(IDs)
	if err != nil {
		logrus.Errorf(constant.Service+"SupervisorGetStudentList Failed, err= %v", err)
		return
	}

	FileInfoMap, err := database.GetStudentFileInfosByIdentityNumbers(IDs)
	if err != nil {
		logrus.Errorf(constant.Service+"SupervisorGetStudentList Failed, err= %v", err)
		return
	}

	result := &api.SupervisorGetStudentListResponse{}
	for _, stu := range stus {
		result.Stus = append(result.Stus, &api.StudentStatusInfo{
			IdentityNumber: stu.IdentityNumber,
			College:        stu.College,
			Class:          stu.Class,
			Length:         stu.Length,
			GraduateTime:   stu.GraduateTime.String(),
			DegreeType:     stu.DegreeType,
			Status:         stu.Status,
			Name:           UserMap[stu.IdentityNumber].Name,
			FirstDraftURL:  FileInfoMap[stu.IdentityNumber].FirstDraft,
		})
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", result))
}

func SupervisorBindStudent(c *gin.Context) {
	identityNumber := sessions.GetUserIdentityNumberBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	studentID := params["studentID"].(string)
	err := database.UpdateSupervisorIDByIdentityNumber(studentID, identityNumber)
	if err != nil {
		logrus.Errorf(constant.Service+"SupervisorBindStudent Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}
