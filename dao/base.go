package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
var db *gorm.DB
func init() {
	log.SetFlags(log.Ldate|log.Ltime |log.Lshortfile)
	var err error
	db, err = gorm.Open("mysql", "root:tiger@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	//一个坑，不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}
	db.SingularTable(true)
	db.LogMode(true)
	log.Printf("连接成功：%#v\n", db)
}

type PageVO struct {
	PageSize  int `gorm:"default:20" form:"pageSize" json:"pageSize"`
	PageIndex int `gorm:"default:1" form:"pageIndex" json:"pageIndex"`
	PageCount int `form:"pageCount" json:"pageCount"`
	Total 	  int `form:"total" json:"total"`
}