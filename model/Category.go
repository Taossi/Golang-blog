package model

import (
	"Golang-blog/utils/errormessage"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

/**
 * @Description: 查询分类是否存在
 * @return code int
 */
func CheckCategory(name string) int {
	var cate Category
	db.Where("name = ?", name).First(&cate)
	if cate.ID != 0 {
		return errormessage.ERROR_CATENAME_USED // 分类已存在
	}
	return errormessage.SUCCESS
}

/**
 * @Description: 创建新分类
 * @return code int
 */
func CreateCategory(cate *Category) int {
	err := db.Create(&cate).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

/**
 * @Description: 删除指定id的分类
 * @return code int
 */
func DeleteCategory(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

/**
 * @Description: 查询分类列表 分页功能
 * @return []Category 分页列表
 */
func GetCategory(pageSize int, pageNum int) []Category {
	var cate []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil {
		return nil
	}
	return cate
}

/**
 * @Description: 编辑分类信息
 * @return code int
 */
func EditCategory(id int, cate *Category) int {
	var maps = make(map[string]interface{})
	maps["name"] = cate.Name
	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

// 查询分类下的所有文章
