package server

import (
	"backend/app/utils"

	"github.com/gin-gonic/gin"
)

type DeleteUserData struct {
	Email string `json:"email" binding:"required"`
	// Username string `json:"username"`
	// Password string `json:"password" binding:"required"`
}

func DeleteUser(c *gin.Context) {
	var deleteuserdata DeleteUserData
	err1 := c.ShouldBindJSON(&deleteuserdata)
	if err1 != nil {
		utils.ResponseError(c, utils.ParameterErrorCode, utils.ParameterErrorMsg)
		return
	}
	err1 = utils.DeleteUser(deleteuserdata.Email)
	if err1 != nil {
		utils.ResponseError(c, utils.OperationFailedCode, utils.OperationFailedMsg)
		return
	} else {
		utils.ResponseSuccess(c, nil)
		return
	}
}

type GetUserData struct {
	Email string `json:"email"`
}

func GetUser(c *gin.Context) {
	var getuserdata GetUserData
	err := c.ShouldBindJSON(&getuserdata)
	if err != nil {
		utils.ResponseError(c, utils.ParameterErrorCode, utils.ParameterErrorMsg)
		return
	}
	user, err := utils.GetUser(getuserdata.Email)
	if err != nil {
		utils.ResponseError(c, utils.OperationFailedCode, utils.OperationFailedMsg)
		return
	} else {
		utils.ResponseSuccess(c, user)
		return
	}
}
