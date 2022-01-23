package v1

import (
	"Golang-blog/model"
	"Golang-blog/utils"
	"Golang-blog/utils/errormessage"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 添加文章接口
func AddArticle(ctx *gin.Context) {
	var article model.Article
	_ = ctx.ShouldBindJSON(&article) // 绑定JSON 自动提取JSON数据
	code := model.CreateArticle(&article)

	if code == errormessage.SUCCESS {
		utils.Success(ctx, gin.H{"status": code, "data": article}, "添加文章成功！")
	} else {
		utils.Fail(ctx, gin.H{"status": code, "data": nil}, "添加文章失败！"+errormessage.GetErrorMsg(code))
	}
}

// 删除文章接口
func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id")) // 读取path中id参数
	code := model.DeleteArticle(id)
	if code == errormessage.SUCCESS {
		utils.Success(ctx, nil, "删除文章成功！")
	} else {
		utils.Fail(ctx, nil, "删除文章失败！")
	}
}

// 查询分类下所有文章接口
func GetCategoryArticle(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize")) // 读取querystring参数
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	id, _ := strconv.Atoi(ctx.Param("id")) // 读取path中id参数

	// 不需要分页时: limit(-1) offset(-1) 终止限制 即全返回
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 1
	}
	articles, code := model.GetCategoryArticle(id, pageSize, pageNum)
	if code == errormessage.SUCCESS {
		utils.Success(ctx, gin.H{"status": code, "data": articles}, "查询分类文章成功！")
	} else {
		utils.Fail(ctx, gin.H{"status": code, "data": nil}, "查询分类文章失败！"+errormessage.GetErrorMsg(code))
	}
}

// 查询单个文章接口
func GetSingleArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id")) // 读取path中id参数
	article, code := model.GetSingleArticle(id)
	if code == errormessage.SUCCESS {
		utils.Success(ctx, gin.H{"status": code, "data": article}, "查询单个文章成功！")
	} else {
		utils.Fail(ctx, gin.H{"status": code, "data": nil}, "查询单个文章失败！"+errormessage.GetErrorMsg(code))
	}
}

// 查询文章列表接口 是否需要分页
func GetArticle(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize")) // 读取querystring参数
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	// 不需要分页时: limit(-1) offset(-1) 终止限制 即全返回
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 1
	}
	users, code := model.GetArticle(pageSize, pageNum)
	if code == errormessage.SUCCESS {
		utils.Success(ctx, gin.H{"status": code, "data": users}, "查询文章列表成功！")
	} else {
		utils.Fail(ctx, gin.H{"status": code, "data": nil}, "查询文章列表失败 ！")
	}
}

// 编辑文章信息接口
func EditArticle(ctx *gin.Context) {
	var article model.Article
	_ = ctx.ShouldBindJSON(&article)
	id, _ := strconv.Atoi(ctx.Param("id")) // 读取path中id参数
	code := model.EditArticle(id, &article)

	if code == errormessage.SUCCESS {
		utils.Success(ctx, nil, "编辑文章成功！")
	} else {
		utils.Success(ctx, nil, "编辑文章失败！"+errormessage.GetErrorMsg(code))
		ctx.Abort()
	}
}
