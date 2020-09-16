package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

// 添加文章
func AddArticle(c*gin.Context)  {
	var art model.Article
	_ =c.ShouldBindJSON(&art)

	errCode :=model.CreateArticle(&art)

	c.JSON(http.StatusOK,gin.H{
		"status":errCode,
		"message":errmsg.GetErrMsg(errCode),
		"data":art,
	})
}
// 某分类下的文章
func GetCateOfArticle(c *gin.Context) {
	pageSize ,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum ,_:=strconv.Atoi(c.Query("pagenum"))
	id,_:=strconv.Atoi(c.Param("id"))
	if pageSize==0{
		pageSize=-1
	}
	if pageNum==0{
		pageNum=-1
	}

	arts,errCode := model.GetCateOfArticle(id,pageSize,pageNum)


	c.JSON(http.StatusOK,gin.H{
		"status":errCode,
		"message":errmsg.GetErrMsg(errCode),
		"data":arts,
	})
}

// 查询单个文章
func GetArticleInfo(c*gin.Context)  {
	id ,_:=strconv.Atoi(c.Param("id"))

	art ,errCode :=model.GetArticleInfo(id)

	c.JSON(http.StatusOK,gin.H{
		"status":errCode,
		"message":errmsg.GetErrMsg(errCode),
		"data":art,

	})
}

// 查询文章列表
func GetArticles(c *gin.Context){
	pageSize ,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum ,_:=strconv.Atoi(c.Query("pagenum"))

	if pageSize==0{
		pageSize=-1
	}
	if pageNum==0{
		pageNum=-1
	}

	arts,errCode := model.GetArticles(pageSize,pageNum)


	c.JSON(http.StatusOK,gin.H{
		"status":errCode,
		"message":errmsg.GetErrMsg(errCode),
		"data":arts,
	})
}
// 编辑文章
func EditArticle(c *gin.Context){
	var art model.Article
	_ =c.ShouldBindJSON(&art)
	id ,_:=strconv.Atoi(c.Param("id"))

	errCode :=model.EditArticle(id,&art)

	c.JSON(http.StatusOK,gin.H{
		"status":errCode,
		"message":errmsg.GetErrMsg(errCode),
	})
}


// 删除文章
func DeleteArticle(c *gin.Context)  {
	// param 和query 一定要分清啊
	id ,_:=strconv.Atoi(c.Param("id"))
	errCode :=model.DeleteArticle(id)

	c.JSON(http.StatusOK,gin.H{
		"status":errCode,
		"message":errmsg.GetErrMsg(errCode),

	})
}