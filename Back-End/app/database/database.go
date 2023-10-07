// 搭建数据库

package database

import (
	"backend/app/config"
	"backend/app/model"

	// "backend/app/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	user := config.Config.GetString("database.user")
	pass := config.Config.GetString("database.pass")
	port := config.Config.GetString("database.port")
	host := config.Config.GetString("database.host")
	name := config.Config.GetString("database.name")

	// user := "root"
	// pass := "asdf112233"
	// host := "127.0.0.1"
	// port := "3360"
	// name := "dd"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connect failed: ", err)
	}
	// 建表
	err1 := db.AutoMigrate(
		&model.Usermodel{},
		&model.Team{},
	)

	if err1 != nil {
		log.Fatal("Database migrate failed: ", err1)
	}

	//创建初始团队
	// newTeam := &model.Team{
	// 	TeamName:        "ZeroIndex",
	// 	TeamDescription: "ZeroIndex",
	// 	Password:        "asdf112233",
	// 	LeaderID:        "00000000",
	// }
	// result := db.Where("team_id=?", "20230000").First(&model.Team{}).Error
	// if errors.Is(result, gorm.ErrRecordNotFound) {
	// 	err = db.Create(newTeam).Error
	// 	if err != nil {
	// 		log.Fatal("Create ZeroIndex Team failed :", err)
	// 	}
	// }

	DB = db
}
