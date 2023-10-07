package model

// import "backend/app/model"

// User
type Usermodel struct {
	//ID       int    `gorm:"primaryKey"`               //设置为主键
	Username   string `gorm:"default:普通用户"`
	Password   string `gorm:"not null"`
	Email      string `gorm:"primary_key"`
	Type       uint8  // 管理员/用户
	Sign       string
	AvataraUrl string
	// TeamID     *int `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TeamID *int `gorm:"index"`
	Team   Team
}
