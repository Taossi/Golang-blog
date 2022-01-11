package v1

import (
	"Golang-blog/model"
	"Golang-blog/utils"
	"Golang-blog/utils/errormessage"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 添加分类接口
func AddCategory(ctx *gin.Context) {
	var category model.Category
	_ = ctx.ShouldBindJSON(&category) // 绑定JSON 自动提取JSON数据
	code := model.CheckCategory(category.Name)

	if code == errormessage.SUCCESS {
		model.CreateCategory(&category)
		utils.Success(ctx, gin.H{"status": code, "data": category}, "添加分类成功！")
	} else {
		utils.Success(ctx, gin.H{"status": code, "data": category}, "添加分类失败！"+errormessage.GetErrorMsg(code))
	}
}

// 删除分类接口
func DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id")) // 读取path中id参数
	code := model.DeleteCategory(id)
	if code == errormessage.SUCCESS {
		utils.Success(ctx, nil, "删除分类成功！")
	} else {
		utils.Fail(ctx, nil, "删除分类失败！")
	}
}

// 查询分类列表接口 是否需要分页
func GetCategory(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize")) // 读取querystring参数
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	// 不需要分页时: limit(-1) offset(-1) 终止限制 即全返回
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 1
	}
	users := model.GetCategory(pageSize, pageNum)
	code := errormessage.SUCCESS
	utils.Success(ctx, gin.H{"status": code, "data": users}, "查询分类列表成功！")
}

// 编辑分类信息接口
func EditCategory(ctx *gin.Context) {
	var category model.Category
	_ = ctx.ShouldBindJSON(&category)
	id, _ := strconv.Atoi(ctx.Param("id")) // 读取path中id参数
	code := model.CheckCategory(category.Name)

	// 分类未被使用 允许更新
	if code == errormessage.SUCCESS {
		model.EditCategory(id, &category)
		utils.Success(ctx, nil, "编辑分类信息成功！")
	} else {
		utils.Success(ctx, nil, "编辑分类信息失败！"+errormessage.GetErrorMsg(code))
		ctx.Abort()
	}
}

// 查询分类下的所有文章接口
