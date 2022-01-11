package v1

import (
	"Golang-blog/model"
	"Golang-blog/utils"
	"Golang-blog/utils/errormessage"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 添加用户接口
func AddUser(ctx *gin.Context) {
	var user model.User
	_ = ctx.ShouldBindJSON(&user) // 绑定JSON 自动提取JSON数据
	code := model.CheckUser(user.Name)

	if code == errormessage.SUCCESS {
		model.CreateUser(&user)
		utils.Success(ctx, gin.H{"status": code, "data": user}, "添加用户成功！")
	} else {
		utils.Success(ctx, gin.H{"status": code, "data": user}, "添加用户失败！"+errormessage.GetErrorMsg(code))
	}
}

// 删除用户接口
func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id")) // 读取path中id参数
	code := model.DeleteUser(id)
	if code == errormessage.SUCCESS {
		utils.Success(ctx, nil, "删除用户成功！")
	} else {
		utils.Fail(ctx, nil, "删除用户失败！")
	}
}

// 查询用户列表接口 是否需要分页
func GetUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize")) // 读取querystring参数
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	// 不需要分页时: limit(-1) offset(-1) 终止限制 即全返回
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 1
	}
	users := model.GetUsers(pageSize, pageNum)
	code := errormessage.SUCCESS
	utils.Success(ctx, gin.H{"status": code, "data": users}, "查询用户列表成功！")
}

// 编辑用户信息(name, role)接口
func EditUser(ctx *gin.Context) {
	var user model.User
	_ = ctx.ShouldBindJSON(&user)
	id, _ := strconv.Atoi(ctx.Param("id")) // 读取path中id参数
	code := model.CheckUser(user.Name)

	// 用户名未被使用 允许更新
	if code == errormessage.SUCCESS {
		model.EditUser(id, &user)
		utils.Success(ctx, nil, "编辑用户信息成功！")
	} else {
		utils.Success(ctx, nil, "编辑用户信息失败！"+errormessage.GetErrorMsg(code))
		ctx.Abort()
	}
}
