package router

import (
	"github.com/gin-gonic/gin"
	v1 "goblog/api/v1"
	"goblog/middleware"
	"goblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	auth :=r.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		// user

		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)


		// category
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)


		// article
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

	}
	router := r.Group("api/v1")
	{
		// user
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)

		// category
		router.GET("categorys", v1.GetCategorys)

		// article
		router.GET("articles", v1.GetArticles)
		router.GET("article/info/:id", v1.GetArticleInfo)
		router.GET("article/category/:id",v1.GetCateOfArticle)


		// login
		router.POST("login",v1.Login)

	}

	_ = r.Run(utils.HttpPort)
}
