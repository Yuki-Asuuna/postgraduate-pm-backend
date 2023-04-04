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
	"postgraduate-pm-backend/utils/redis"
	"postgraduate-pm-backend/utils/sessions"
)

func Login(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	identityNumber := params["identityNumber"].(string)
	password := params["password"].(string)

	// 通过md5生成counsellorID
	user, err := database.GetUserByIdentityNumber(identityNumber)

	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"Login Failed, err= %v", err)
		return
	}
	if user == nil {
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-3, "identityNumber not found", nil))
		return
	}
	if password != user.Password {
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-3, "Incorrect Password", nil))
		return
	}
	session, _ := sessions.GetSessionClient().Get(c.Request, "dotcomUser")
	session.Values["authenticated"] = true
	session.Values["identityNumber"] = identityNumber
	err = redis.SetOnline(identityNumber)
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"Login Failed, err= %v", err)
		return
	}
	err = sessions.GetSessionClient().Save(c.Request, c.Writer, session)
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"Login Failed, err= %v", err)
		return
	}
	if err := database.UpdateLoginTimeByIdentityNumber(identityNumber); err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"Login Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func Logout(c *gin.Context) {
	session, _ := sessions.GetSessionClient().Get(c.Request, "dotcomUser")
	session.Values["authenticated"] = false
	err := sessions.GetSessionClient().Save(c.Request, c.Writer, session)
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"Logout Failed, err= %v", err)
		return
	}
	err = redis.SetOffline(session.Values["identityNumber"].(string))
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"Logout Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func Me(c *gin.Context) {
	user := sessions.GetUserInfoBySession(c)
	if user == nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service + "Me Get Personal Info Failed, user is nil")
		return
	}
	result := &api.MeResponse{
		IdentityNumber: user.IdentityNumber,
		Name:           user.Name,
		Role:           user.Role,
		Gender:         user.Gender,
		Age:            user.Age,
		PhoneNumber:    user.PhoneNumber,
		LastLogin:      user.LastLogin,
		Avatar:         user.Avatar,
		Email:          user.Email,
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", result))
}

func ChangePassword(c *gin.Context) {
	user := sessions.GetUserInfoBySession(c)
	if user == nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service + "ChangePassword Get Personal Info Failed, user is nil")
		return
	}
	params := make(map[string]interface{})
	c.BindJSON(&params)
	oldPassword := params["oldPassword"].(string)
	newPassword := params["newPassword"].(string)
	if user.Password != oldPassword {
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-2, "旧密码不正确", nil))
		return
	}
	err := database.UpdatePasswordByIdentityNumber(user.IdentityNumber, newPassword)
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"ChangePassword Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func AvatarUpload(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	url := params["url"].(string)
	user := sessions.GetUserInfoBySession(c)
	if user == nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service + "AvatarUpload Failed, user is nil")
		return
	}
	err := database.UpdateUserAvatarByIdentityNumber(user.IdentityNumber, url)
	if err != nil {
		logrus.Errorf(constant.Service+"AvatarUpload Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func PostMe(c *gin.Context) {
	user := sessions.GetUserInfoBySession(c)
	if user == nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service + "Me Get Personal Info Failed, user is nil")
		return
	}
	params := make(map[string]interface{})
	c.BindJSON(&params)
	name := params["name"].(string)
	gender := int64(params["gender"].(float64))
	age := int64(params["age"].(float64))
	phoneNumber := params["phoneNumber"].(string)
	email := params["email"].(string)
	identityNumber := user.IdentityNumber
	err := database.UpdateUserByIdentityNumber(identityNumber, name, gender, age, phoneNumber, email)
	if err != nil {
		logrus.Errorf(constant.Service+"PostMe Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))

}
