package routers

import (
	"Golang-blog/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 路由接口初始化函数
func InitRouters() {
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// 用户模块路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.DELETE("users/:id", v1.DeleteUser)
		router.PUT("users/:id", v1.EditUser)

		// 文章模块路由接口
		router.POST("article/add", v1.AddArticle)
		router.GET("article", v1.GetArticle)
		router.GET("article/category/:id", v1.GetCategoryArticle)
		router.GET("article/single/:id", v1.GetSingleArticle)
		router.DELETE("article/:id", v1.DeleteArticle)
		router.PUT("article/:id", v1.EditArticle)

		// 分类模块路由接口
		router.POST("category/add", v1.AddCategory)
		router.GET("category", v1.GetCategory)
		router.DELETE("category/:id", v1.DeleteCategory)
		router.PUT("category/:id", v1.EditCategory)

	}

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}
