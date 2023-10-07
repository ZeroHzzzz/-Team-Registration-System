package utils

import (
	"backend/app/database"
	"backend/app/model"
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

// 拷贝(凑合用)
func copyFields(source, target interface{}) {
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(target).Elem() // 获取目标结构体的可写入部分

	for i := 0; i < sourceValue.NumField(); i++ {
		field := sourceValue.Type().Field(i)
		sourceFieldValue := sourceValue.Field(i)

		// 检查字段是否非空
		if !isEmpty(sourceFieldValue) {
			// 获取目标结构体中对应字段
			targetField := targetValue.FieldByName(field.Name)
			if targetField.IsValid() && targetField.CanSet() {
				// 设置目标结构体中字段的值
				targetField.Set(sourceFieldValue)
			}
		}
	}
}

// 判断字段是否为空
func isEmpty(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return field.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return field.IsNil()
	default:
		zeroValue := reflect.Zero(field.Type())
		return reflect.DeepEqual(field.Interface(), zeroValue.Interface())
	}
}

// 格式错误会返回一个true，正确返回false
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return !reg.MatchString(email)
}

// 登录判断
func CheckLogin(Email string, Password string) error {
	err := database.DB.Where("email = ? AND password = ?", Email, Password).First(&model.Usermodel{}).Error
	return err
}

// 账户存在性
func CheckUserExistByEmail(Email string) error {
	err := database.DB.Where("email = ?", Email).First(&model.Usermodel{}).Error
	return err
}

// 创建用户
func CreateNewUser(username string, password string, email string) error {
	db := database.DB
	newUser := &model.Usermodel{
		Username: username,
		Password: password,
		Email:    email,
		// TeamID:   20230000,
	}
	// if newUser.
	err := db.Create(newUser)
	// fmt.Println(err)
	return err.Error
}

// 查询用户
func GetUser(Email string) (*model.Usermodel, error) {
	var user model.Usermodel
	err := database.DB.Where("email=?", Email).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

// 校对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}

// 注销账号
func DeleteUser(email string) error {
	db := database.DB
	err := db.Delete(&model.Usermodel{}, "email=?", email)
	return err.Error
}

// 团队

// 创建团队(如果是空值可能为bug)
// func CreateTeam(teamdata server.CreateTeamData) (int, error) {
// 	db := database.DB
// var existingLeader model.Usermodel
// // Check LeaderID
// if err := db.Where("email = ?", teamdata.LeaderID).First(&existingLeader).Error; err != nil {
// 	return -1, err
// }
// if existingLeader.TeamID != 0 {
// 	return -1, errors.New("Leader is already part of another team")
// }

// 	// 使用反射获取结构体类型
// 	userType := reflect.TypeOf(user)

// 	// 使用反射获取结构体值
// 	userValue := reflect.ValueOf(&user).Elem()

// 	// 遍历结构体字段
// 	for i := 0; i < userType.NumField(); i++ {
// 		field := userType.Field(i)
// 		fieldValue := userValue.Field(i)

//			// 如果字段值为空，跳过
//			if fieldValue.Interface() == reflect.Zero(field.Type).Interface() {
//				continue
//			}
//		}
//		// newTeam := &model.Team{
//		// 	TeamName:        TeamName,
//		// 	TeamDescription: TeamDescription,
//		// 	Password:        Password,
//		// 	LeaderID:        LeaderID,
//		// }
//		err := db.Create(newTeam).Error //会调用model中的BeforeCreate
//		if err != nil {
//			// fmt.Printf("Error creating team: %+v\n", err)
//			// panic(err)
//			return -1, err
//		} else {
//			return newTeam.TeamID, nil
//		}
//	}
//
// 判断团队是否存在
func CheckTeamExist(TeamID int) error {
	result := database.DB.Where("team_id?", TeamID).First(&model.Team{})
	return result.Error
}

// 查询团队
func GetTeam(TeamID int) (*model.Team, error) {
	var team model.Team
	err := database.DB.Preload("Users").Where("team_id=?", TeamID).First(&team)
	if err.Error != nil {
		return nil, err.Error
	}
	return &team, nil
}
func CheckAdd(TeamID int, Password string) error {
	err := database.DB.Where("team_id = ? AND password = ?", TeamID, Password).First(&model.Team{}).Error
	return err
}

// check TeamID
func IsInTeam(Email string) bool {
	User, err := GetUser(Email)
	if err != nil {
		return false
	}
	return User.TeamID == nil
}

// 创建团队
func CreateTeam(teamdata interface{}) (int, error) {
	newTeam := model.Team{}
	db := database.DB

	copyFields(teamdata, &newTeam)

	err := db.Create(&newTeam).Error
	if err != nil {
		fmt.Printf("Error creating team: %+v\n", err)
		// panic(err)
		return -1, err
	} else {
		err = AddMember(newTeam.LeaderID, newTeam.TeamID)
		return newTeam.TeamID, err
	}
}

// 添加成员
func AddMember(Email string, TeamID int) error {
	db := database.DB
	if !IsInTeam(Email) {
		return errors.New("Have joined a team")
	}
	member, err := GetUser(Email)
	if err != nil {
		return err
	}
	team, err := GetTeam(TeamID)
	if err != nil {
		return err
	}
	db.Model(&team).Association("Users").Append(member)

	db.Save(&team)
	return nil
}

// 从团队查询所有用户信息(*)
func GetTeamMember(TeamID int) ([]model.Usermodel, error) {
	db := database.DB
	var team model.Team
	var users []model.Usermodel

	err := db.Preload("users").First(&team, TeamID).Error
	if err != nil {
		return nil, err
	}

	users = team.Users
	return users, nil
}

// 解散
func DelTeam(TeamID int) error {
	db := database.DB
	// if err := db.Delete(&model.Team{}, "team_id=?", TeamID).Error; err != nil {
	// 	return err
	// }
	var team model.Team
	db.Preload("Users").Where("team_id=?", TeamID).First(&team, TeamID)
	db.Model(&team).Association("Users").Delete(&team.Users)
	// return nil
	err := db.Delete(&model.Team{}, "team_id=?", TeamID).Error
	return err
}
