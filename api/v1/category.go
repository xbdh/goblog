package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)


// 添加分类
func AddCategory(c *gin.Context)  {
	var cate model.Category
	_ =c.ShouldBindJSON(&cate)

	errCode :=model.CheckCategory(cate.Name)
	if errCode==errmsg.SUCCSE{
		_ =model.CreateCategory(&cate)
	}

	if errCode==errmsg.ERROR_CATENAME_USED{
		errCode =errmsg.ERROR_CATENAME_USED
	}

	c.JSON(http.StatusOK,gin.H{
		"status":errCode,
		"message":errmsg.GetErrMsg(errCode),
		"data":cate,
	})
}

// 查询单个分类
func GetCategoryInfo(c *gin.Context)  {

}

// 查询分类列表
func GetCategorys(c *gin.Context){
	pageSize ,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum ,_:=strconv.Atoi(c.Query("pagenum"))

	if pageSize==0{
		pageSize=-1
	}
	if pageNum==0{
		pageNum=-1
	}

	cates:= model.GetCategorys(pageSize,pageNum)
	errCode:=errmsg.SUCCSE
	c.JSON(http.StatusOK,gin.H{
		"status":errCode,
		"message":errmsg.GetErrMsg(errCode),
		"data":cates,
	})
}
// 编辑分类
func EditCategory(c *gin.Context){
	var cate model.Category
	_ =c.ShouldBindJSON(&cate)
	id ,_:=strconv.Atoi(c.Param("id"))
	errCode :=model.CheckCategory(cate.Name)
	if errCode==errmsg.SUCCSE{
		_ =model.EditCategory(id,&cate)
	}

	if errCode==errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"status":errCode,
		"message":errmsg.GetErrMsg(errCode),
	})
}


// 删除分类

func DeleteCategory(c *gin.Context)  {
	// param 和query 一定要分清啊
	id ,_:=strconv.Atoi(c.Param("id"))
	errCode :=model.DeleteCategory(id)

	c.JSON(http.StatusOK,gin.H{
		"status":errCode,
		"message":errmsg.GetErrMsg(errCode),

	})
}