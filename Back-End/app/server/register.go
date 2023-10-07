package server

import (
	"backend/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 解析信息
type RegisterData struct {
	Email string `jion:"email" binding:"required"`
	// Username string `json:"username" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var registerdata RegisterData
	err1 := c.ShouldBindJSON(&registerdata)
	if err1 != nil {
		utils.ResponseError(c, utils.ParameterErrorCode, utils.ParameterErrorMsg)
		return
	}
	// 邮箱格式
	if utils.VerifyEmailFormat(registerdata.Email) {
		utils.ResponseError(c, utils.FormatWrongCode, utils.FormatWrongMsg)
		return
	}
	// 邮箱是否已经被注册
	err1 = utils.CheckUserExistByEmail(registerdata.Email)
	if err1 == nil {
		utils.ResponseError(c, utils.HaveExistCode, utils.UserHaveExistMsg)
		return
	} else if err1 != nil && err1 != gorm.ErrRecordNotFound {
		utils.ResponseInternalError(c)
		return
	}
	// 默认用户名
	if len(registerdata.Username) == 0 {
		registerdata.Username = "普通用户"
	}

	// 创建用户
	err1 = utils.CreateNewUser(registerdata.Username, registerdata.Password, registerdata.Email)
	if err1 == nil {
		utils.ResponseSuccess(c, nil)
	} else {
		utils.ResponseError(c, utils.OperationFailedCode, utils.OperationFailedMsg)
	}
}
