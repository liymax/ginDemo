package main

import (
	"fmt"
	"github.com/liymax/gindemo/dao"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()

	r.POST("/create", func(c *gin.Context) {
		var bu dao.BusinessInfo
		if err := c.ShouldBind(&bu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		bu.CreateAt = time.Now()
		bu.UpdateAt = time.Now()
		fmt.Printf("%v\n", bu)
		res := bu.Insert()
		c.JSON(http.StatusOK, res)
	})

	r.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		res := dao.Delete(id)
		c.JSON(http.StatusOK, res)
	})

	r.POST("/update", func(c *gin.Context) {
		var bu dao.BusinessInfo
		if err := c.ShouldBindJSON(&bu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if bu.Id > 0 {
			bu.UpdateAt = time.Now()
			res := bu.Update()
			c.JSON(http.StatusOK, res)
		}else {
			c.JSON(http.StatusOK, gin.H{"code": 6, "msg": "id不存在"})
		}
	})

	r.GET("/business/:id", func(c *gin.Context) {
		id := c.Param("id")
		b,err := dao.FindById(id)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 0,"data": b})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 5, "msg": "查询失败"})
		}
	})

	r.POST("/business/ids", func(c *gin.Context) {
		var ids []int
		if c.Bind(&ids) != nil {
			c.JSON(http.StatusOK, gin.H{"code": 4, "msg": "请求参数异常"})
		}
		fmt.Printf("%v\n", ids)

		bs,err := dao.FindByIds(ids)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 0,"data": bs})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 5, "msg": "查询失败"})
		}
	})

	r.GET("/businessList", func(c *gin.Context) {
		var pv dao.PageVO
		if c.ShouldBind(&pv) == nil {
			fmt.Printf("%v\n", pv)
			bs,err := dao.FindList(pv)
			if err == nil {
				c.JSON(http.StatusOK, gin.H{"code": 0,"data": bs})
			} else {
				c.JSON(http.StatusOK, gin.H{"code": 5, "msg": "查询失败"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 4, "msg": "请求参数异常"})
		}
	})
	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("%v\n", err)
	}
}
