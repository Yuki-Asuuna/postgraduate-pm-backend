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
	// TODO
}

func SupervisorPostComment(c *gin.Context) {
	// TODO
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
