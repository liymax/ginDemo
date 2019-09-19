package main

import (
	"github.com/liymax/gindemo/dao"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
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
		log.Printf("%v\n", bu)
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
		log.Printf("%v\n", ids)

		bs,err := dao.FindByIds(ids)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 0,"data": bs})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 5, "msg": "查询失败"})
		}
	})
	r.GET("/count", func(c *gin.Context) {
		count,err := dao.Count()
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 0,"data": count})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 5, "msg": "查询失败"})
		}
	})
	r.GET("/businessList", func(c *gin.Context) {
		var pv dao.PageVO
		if c.ShouldBind(&pv) == nil {
			log.Printf("%v\n", pv)
			bs,err := dao.FindList(pv)
			if err == nil {
				c.JSON(http.StatusOK, gin.H{"code": 0,"data": bs})
			} else {
				c.JSON(http.StatusOK, gin.H{"code": 5, "msg": "查询失败"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"code": 4, "msg": "请求参数异常"})
		}
	})
	
	r.GET("/sys/memo", func(c *gin.Context) {
		cmd := exec.Command("/bin/bash", "-c", "free -h")
		stdout, _ := cmd.StdoutPipe()
		if err := cmd.Start(); err != nil{
			msg := "Execute failed when Start:" + err.Error()
			c.JSON(http.StatusOK, gin.H{"code": 5, "msg": msg})
			return
		}

		outBytes, _ := ioutil.ReadAll(stdout)
		_ = stdout.Close()

		if err := cmd.Wait(); err != nil {
			msg := "Execute failed when Wait:" + err.Error()
			c.JSON(http.StatusOK, gin.H{"code": 5, "msg": msg})
		}else {
			c.JSON(http.StatusOK, gin.H{"code": 0, "data": string(outBytes)})
		}
	})

	r.StaticFile("/", "./webroot/build/index.html")
	r.StaticFile("/favicon.ico", "./webroot/build/favicon.ico")
	r.Static("/static", "./webroot/build/static")

	if err := r.Run(":8080"); err != nil {
		log.Printf("%v\n", err.Error())
	}
}
