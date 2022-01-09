package model

import (
	"Golang-blog/utils/errormessage"
	"gorm.io/gorm"
)

//User 用户数据模型
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null" json:"name"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

/**
 * @Description: 查询用户是否存在接口
 * @return code int
 */
func CheckUser(name string) int {
	var user User
	db.Where("name = ?", name).First(&user)
	if user.ID != 0 {
		return errormessage.ERROR_USERNAME_USED // 用户已存在
	}
	return errormessage.SUCCESS
}

/**
 * @Description: 创建新用户接口
 * @return code int
 */
func CreateUser(user *User) int {
	err := db.Create(&user).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

/**
 * @Description: 查询用户列表接口 分页功能
 * @return []User
 */
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}
