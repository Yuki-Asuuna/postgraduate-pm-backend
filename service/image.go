package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"postgraduate-pm-backend/exception"
	"postgraduate-pm-backend/utils"
	"postgraduate-pm-backend/utils/upload_image"
)

func ImageUpload(c *gin.Context) {
	f, err := c.FormFile("source")
	if err != nil {
		logrus.Errorf("ImageUpload Failed, err= %v", err)
		c.Error(exception.ParameterError())
		return
	}
	url, err := upload_image.GetImageUrl(f)
	if err != nil {
		logrus.Errorf("ImageUpload Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", url))
}
