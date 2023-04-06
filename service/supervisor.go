package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"postgraduate-pm-backend/api"
	"postgraduate-pm-backend/constant"
	"postgraduate-pm-backend/database"
	"postgraduate-pm-backend/utils"
	"postgraduate-pm-backend/utils/helper"
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
	stus, err := database.GetStudentStatusInfoListBySupervisorID(supervisorIdentityNumber, 0, 999999)
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
		c.JSON(http.StatusUnauthorized, utils.GenSuccessResponse(-3, "student not belong to current supervisor", nil))
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
	stus, err := database.GetStudentStatusInfoListBySupervisorID(supervisorIdentityNumber, 0, 999999)
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
		c.JSON(http.StatusUnauthorized, utils.GenSuccessResponse(-3, "student not belong to current supervisor", nil))
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
	page := helper.S2I(c.DefaultQuery("page", "0"))
	size := helper.S2I(c.DefaultQuery("size", "10"))
	identityNumber := sessions.GetUserIdentityNumberBySession(c)
	stus, err := database.GetStudentStatusInfoListBySupervisorID(identityNumber, page, size)
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
		file := FileInfoMap[stu.IdentityNumber]
		result.Stus = append(result.Stus, &api.StudentStatusInfo{
			IdentityNumber:                        stu.IdentityNumber,
			College:                               stu.College,
			Class:                                 stu.Class,
			Length:                                stu.Length,
			GraduateTime:                          stu.GraduateTime.String(),
			DegreeType:                            stu.DegreeType,
			Status:                                stu.Status,
			Name:                                  UserMap[stu.IdentityNumber].Name,
			FirstDraftURL:                         file.FirstDraft,
			IsFirstDraftConfirmed:                 helper.I2B(file.IsFirstDraftConfirmed),
			PreliminaryReviewFormURL:              file.PreliminaryReviewForm,
			IsPreliminaryReviewFormConfirmed:      helper.I2B(file.IsPreliminaryReviewFormConfirmed),
			IsResearchEvaluationMaterialConfirmed: helper.I2B(file.IsResearchEvaluationMaterialConfirmed),
			ResearchEvaluationMaterialURL:         file.ResearchEvaluationMaterial,
			BlindScore:                            stu.BlindScore,
			DefenseScore:                          stu.DefenseScore,
			DegreeConfirmed:                       helper.I2B(stu.DegreeConfirmed),
			ApplyDegree:                           helper.I2B(stu.ApplyDegree),
		})
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", result))
}

func SupervisorBindStudent(c *gin.Context) {
	var err error
	identityNumber := sessions.GetUserIdentityNumberBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	studentID := params["studentID"].(string)
	IsBind := params["bind"].(bool)
	if !IsBind {
		err = database.UpdateSupervisorIDByIdentityNumber(studentID, "")
	} else {
		err = database.UpdateSupervisorIDByIdentityNumber(studentID, identityNumber)
	}
	if err != nil {
		logrus.Errorf(constant.Service+"SupervisorBindStudent Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func SupervisorConfirmStudent(c *gin.Context) {
	// identityNumber := sessions.GetUserIdentityNumberBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	studentID := params["studentID"].(string)
	confirmType := int64(params["confirmType"].(float64))
	if confirmType == 0 { // first draft
		err := database.UpdateIsFirstDraftConfirmedByIdentityNumber(studentID, 1)
		if err != nil {
			logrus.Errorf(constant.Service+"SupervisorConfirmStudent Failed, err= %v", err)
			return
		}
	} else if confirmType == 1 { // preliminary review form
		err := database.UpdateIsPreliminaryReviewFormConfirmedByIdentityNumber(studentID, 1)
		if err != nil {
			logrus.Errorf(constant.Service+"SupervisorConfirmStudent Failed, err= %v", err)
			return
		}
	} else if confirmType == 2 { // research evaluation material
		err := database.UpdateIsResearchEvaluationMaterialConfirmedByIdentityNumber(studentID, 1)
		if err != nil {
			logrus.Errorf(constant.Service+"SupervisorConfirmStudent Failed, err= %v", err)
			return
		}
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}
