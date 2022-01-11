package model

import (
	"Golang-blog/utils/errormessage"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

//User 用户数据模型
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null" json:"name"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

/**
 * @Description: 查询用户是否存在
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
 * @Description: 创建新用户
 * @return code int
 */
func CreateUser(user *User) int {
	user.Password = BcryptPassword(user.Password) // 密码加密
	err := db.Create(&user).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

/**
 * @Description: 删除指定id的用户
 * @return code int
 */
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

/**
 * @Description: 查询用户列表 分页功能
 * @return []User 用户列表
 */
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

/**
 * @Description: 编辑用户信息 (不包括密码)
 * @return code int
 */
func EditUser(id int, user *User) int {
	var maps = make(map[string]interface{})
	maps["name"] = user.Name
	maps["role"] = user.Role
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

/**
 * @Description: 密码加密函数
 * @return 加密后的密码字符串
 */
func BcryptPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}
