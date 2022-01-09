package v1

import (
	"Golang-blog/model"
	"Golang-blog/utils"
	"Golang-blog/utils/errormessage"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 查询用户列表 是否需要分页
func GetUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
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
	utils.Success(ctx, gin.H{"status": code, "data": users}, errormessage.GetErrorMsg(code))
}

// 添加用户
func AddUser(ctx *gin.Context) {
	var user model.User
	_ = ctx.ShouldBindJSON(&user) // 绑定JSON 自动提取JSON数据
	code := model.CheckUser(user.Name)

	if code == errormessage.SUCCESS {
		model.CreateUser(&user)
	}
	utils.Success(ctx, gin.H{"status": code, "data": user}, errormessage.GetErrorMsg(code))
}
