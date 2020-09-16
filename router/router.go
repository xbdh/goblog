package router

import (
	"github.com/gin-gonic/gin"
	v1 "goblog/api/v1"
	"goblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// user
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)

		// category
		router.POST("category/add", v1.AddCategory)
		router.GET("categorys", v1.GetCategorys)

		router.PUT("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCategory)

		// article
		router.POST("article/add", v1.AddArticle)

		router.GET("articles", v1.GetArticles)
		router.GET("article/info/:id", v1.GetArticleInfo)
		router.GET("article/category/:id",v1.GetCateOfArticle)

		router.PUT("article/:id", v1.EditArticle)
		router.DELETE("article/:id", v1.DeleteArticle)

	}

	_ = r.Run(utils.HttpPort)
}
