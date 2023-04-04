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
	"postgraduate-pm-backend/utils/helper"
)

func AdminGetStudentList(c *gin.Context) {
	page := helper.S2I(c.DefaultQuery("page", "0"))
	size := helper.S2I(c.DefaultQuery("size", "10"))
	identityNumber := c.DefaultQuery("identityNumber", "")
	stus, err := database.GetStudentStatusInfoList(page, size, identityNumber)
	if err != nil {
		logrus.Errorf(constant.Service+"AdminGetStudentList Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	IDs := make([]string, 0)
	for _, stu := range stus {
		IDs = append(IDs, stu.IdentityNumber)
	}
	UserMap, err := database.GetUsersByIdentityNumbers(IDs)
	if err != nil {
		logrus.Errorf(constant.Service+"AdminGetStudentList Failed, err= %v", err)
		return
	}
	FileInfoMap, err := database.GetStudentFileInfosByIdentityNumbers(IDs)
	if err != nil {
		logrus.Errorf(constant.Service+"AdminGetStudentList Failed, err= %v", err)
		return
	}
	result := &api.AdminGetStudentListResponse{}
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

func AdminUploadBlindScore(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	identityNumber := params["identityNumber"].(string)
	score := int64(params["score"].(float64))
	err := database.UpdateBlindScore(identityNumber, score)
	if err != nil {
		logrus.Errorf(constant.Service+"AdminUploadBlindScore Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func AdminUploadDefenseScore(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	identityNumber := params["identityNumber"].(string)
	score := int64(params["score"].(float64))
	err := database.UpdateDefenseScore(identityNumber, score)
	if err != nil {
		logrus.Errorf(constant.Service+"AdminUploadDefenseScore Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func AdminUploadDegreeConfirmed(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	identityNumber := params["identityNumber"].(string)
	applyDegree := params["degreeConfirmed"].(bool)
	var ad int64
	if applyDegree {
		ad = 1
	}
	err := database.UpdateDegreeConfirmed(identityNumber, ad)
	if err != nil {
		logrus.Errorf(constant.Service+"AdminUploadDegreeConfirmed Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}
