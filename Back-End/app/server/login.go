package server

import (
	"backend/app/model"
	"backend/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 解析信息
type LoginData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var logindata LoginData
	err := c.ShouldBindJSON(&logindata)
	if err != nil {
		utils.ResponseError(c, utils.ParameterErrorCode, utils.ParameterErrorMsg)
		return
	}
	//格式-邮箱
	if utils.VerifyEmailFormat(logindata.Email) {
		utils.ResponseError(c, utils.FormatWrongCode, utils.FormatWrongMsg)
		return
	}

	err = utils.CheckLogin(logindata.Email, logindata.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseInternalError(c)
		} else {
			utils.ResponseError(c, utils.NotFoundCode, utils.UserNotFoundMsg)
		}
		return
	}

	// 这里为了方便检查，以后会改
	var user *model.Usermodel
	user, err = utils.GetUser(logindata.Email)
	if err != nil {
		utils.ResponseInternalError(c)
		return
	}
	utils.ResponseSuccess(c, user)
}
