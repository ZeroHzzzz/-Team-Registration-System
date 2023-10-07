package server

import (
	"backend/app/model"
	"backend/app/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// type GetTeamMemberData struct {
// 	TeamID int `json:"teamid" binding:"required"`
// 	// Username string `json:"username" binding:"required"`
// 	// Username string `json:"username" binding:"required"`
// 	// Password string `json:"password" binding:"required"`
// }

type CreateTeamData struct {
	TeamName        string `json:"teamname" binding:"required"`
	Password        string `json:"password" binding:"required"`
	TeamDescription string `json:"teamdescription"`
	LeaderID        string `json:"leaderid" binding:"required"` // email
}

func CreateTeam(c *gin.Context) {
	var createteamdata CreateTeamData
	var newteamid int
	err1 := c.ShouldBindJSON(&createteamdata)
	if err1 != nil {
		utils.ResponseError(c, 400, "参数错误")
		return
	}
	if utils.CheckUserExistByEmail(createteamdata.LeaderID) != nil {
		utils.ResponseError(c, 401, "然而用户并不存在")
		return
	}
	if !utils.IsInTeam(createteamdata.LeaderID) {
		utils.ResponseError(c, 402, "LeaderID 已经存在于其他团队中")
		return
	}
	newteamid, err1 = utils.CreateTeam(createteamdata)
	if err1 != nil || newteamid == -1 {
		utils.ResponseError(c, 403, "操作未成功完成")
		return
	} else {
		utils.ResponseSuccess(c, newteamid)
		return
	}
}

type AddMemberData struct {
	Email    string `json:"email" binding:"required"`
	TeamID   int    `json:"teamid" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AddMember(c *gin.Context) {
	var addmemberdata AddMemberData
	err := c.ShouldBindJSON(&addmemberdata)
	if err != nil {
		utils.ResponseError(c, utils.ParameterErrorCode, utils.ParameterErrorMsg)
		return
	}
	err = utils.CheckAdd(addmemberdata.TeamID, addmemberdata.Password)
	if err != nil {
		utils.ResponseError(c, utils.NotFoundCode, utils.TeamNotFoundFailMsg)
		return
	}
	err = utils.AddMember(addmemberdata.Email, addmemberdata.TeamID)
	if err != nil {
		utils.ResponseError(c, utils.OperationFailedCode, utils.OperationFailedMsg)
		return
	}
}

type GetTeamData struct {
	TeamID int `json:"teamid" binding:"required"`
}

func GetTeam(c *gin.Context) {
	var geteamdata GetTeamData
	err := c.ShouldBindJSON(&geteamdata)
	if err != nil {
		utils.ResponseError(c, utils.ParameterErrorCode, utils.ParameterErrorMsg)
		return
	}
	Team, err := utils.GetTeam(geteamdata.TeamID)
	// 可能会出bug
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ResponseError(c, utils.NotFoundCode, utils.NotFoundMsg)
		return
	} else {
		utils.ResponseSuccess(c, Team)
		return
	}
}

func GetTeamMember(c *gin.Context) {
	var users []model.Usermodel
	var geteammemberdata GetTeamData
	err1 := c.ShouldBindJSON(&geteammemberdata)
	if err1 != nil {
		utils.ResponseError(c, utils.ParameterErrorCode, utils.ParameterErrorMsg)
		return
	}
	users, err1 = utils.GetTeamMember(geteammemberdata.TeamID)
	if err1 != nil {
		utils.ResponseError(c, utils.OperationFailedCode, utils.OperationFailedMsg)
		return
	} else {
		utils.ResponseSuccess(c, users)
		return
	}
}

func DelTeam(c *gin.Context) {
	var geteamdata GetTeamData
	err := c.ShouldBindJSON(&geteamdata)
	if err != nil {
		utils.ResponseError(c, utils.ParameterErrorCode, utils.ParameterErrorMsg)
		return
	}
	err = utils.DelTeam(geteamdata.TeamID)
	if err != nil {
		utils.ResponseError(c, utils.OperationFailedCode, utils.OperationFailedMsg)
		return
	} else {
		utils.ResponseSuccess(c, nil)
		return
	}
}
