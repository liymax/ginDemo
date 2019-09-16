package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)
type BusinessInfo struct {
	Id int `gorm:"primary_key" form:"id" json:"id"`
	Uid string `form:"uid" json:"uid" binding:"required"`
	RegisterStatus int8 `form:"registerStatus" json:"registerStatus"`
	SrmNumber string `form:"srmNumber" json:"srmNumber"`
	CustomerBg string `form:"customerBg" json:"customerBg"`
	CreateAt time.Time `form:"createAt" json:"createAt"`
	UpdateAt time.Time `form:"createAt" json:"updateAt"`
}
func (business *BusinessInfo) Insert() gin.H{
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	if err := db.Table("business_info").Create(business).Error; err != nil {
		fmt.Printf("%v\n", err)
		return gin.H{"code": 5, "data": nil, "msg": "创建失败"}
	} else {
		return gin.H{"code": 0, "data": nil, "msg": "创建成功"}
	}
}

func Delete(id string) gin.H{
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	if err := db.Table("business_info").Where("id = ?", id).Delete(&BusinessInfo{}).Error; err != nil {
		fmt.Printf("%v\n", err)
		return gin.H{"code": 5, "data": nil, "msg": "删除失败"}
	} else {
		return gin.H{"code": 0, "data": nil, "msg": "删除成功"}
	}
}

func (business *BusinessInfo) Update() gin.H{
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	if err := db.Table("business_info").Model(business).Update(business).Error; err != nil {
		fmt.Printf("%v\n", err)
		return gin.H{"code": 5, "data": nil, "msg": "更新失败"}
	} else {
		return gin.H{"code": 0, "data": nil, "msg": "更新成功"}
	}
}

func FindById(id string)  (b BusinessInfo,err error){
	err = db.Table("business_info").Find(&b, "id = ?", id).Error
	return b, err
}

func FindByIds(ids []int) (bs []BusinessInfo,err error){
	err = db.Table("business_info").Where(ids).Find(&bs).Error
	return bs, err
}

func FindList(po PageVO) (bs []BusinessInfo,err error){
	offset := po.PageSize*(po.PageIndex - 1)
	err = db.Table("business_info").Offset(offset).Limit(po.PageSize).Find(&bs).Error
	return bs, err
}
