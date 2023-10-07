package model

import (
	"gorm.io/gorm"
)

type Team struct {
	//ID       int    `gorm:"primaryKey"`
	TeamName        string `gorm:"not null type:varchar(25)"`
	Password        string
	TeamDescription string `gorm:"type:varchar(25)"`
	LeaderID        string // Email
	TeamID          int    `gorm:"primary_key"`
	Users           []Usermodel
}

// 修改id规则
func (team *Team) BeforeCreate(tx *gorm.DB) error {
	if team.LeaderID == "00000000" {
		team.TeamID = 20230000
		return nil
	} else {
		team.TeamID = generateCustomTeamID(tx)
		return nil
	}
}
func generateCustomTeamID(tx *gorm.DB) int {
	var count int64
	tx.Model(&Team{}).Count(&count)
	return int(20230000 + count + 1)
}

// func (team *Team) BeforeDelete(tx *gorm.DB) error {
// 	teamID := team.TeamID
// 	if err := tx.Model(&Usermodel{}).Where("team_id=?", teamID).Update("team_id", nil).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
