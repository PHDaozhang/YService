package common

import (
	"YService/controllers/system"
	"YService/core/errorcode"
	"YService/models/dto"
	"YService/models/info"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"tsEngine/tsDb"
	"tsEngine/tsOpCode"
	"tsEngine/tsTime"
)

//各种广告的
type AdController struct {
	system.BaseController
}


// @Title 添加广告
// @Description 添加在图头部或是尾部的图片
// @Success 200 {object} nil
// @Param    Title    query    string    false  标题
// @Param    Type       query    int    true   类型
// @Param    Address   query    string    true   多个地址，以逗号分隔
// @Param    Link   query    string    true   多个地址，以逗号分隔
// @router / [post]
func (this* AdController)Add(){
	var req dto.ReqAdInfo

	if err := this.ParseForm(&req); err != nil {
		this.Error(errorcode.ERROR_PARAM)
	}

	valid := validation.Validation{}
	b,err := valid.Valid(&req)
	if err != nil {
		this.Error(errorcode.ERROR_VALID)
	}

	if !b {
		this.Error(errorcode.ERROR_VALID)
	}

	o := orm.NewOrm()

	adInfo := info.AdInfo{}
	adInfo.Title = req.Title
	adInfo.Address = req.Address
	adInfo.Link = req.Link
	adInfo.Type = req.Type
	adInfo.CrtTime = tsTime.CurrMs()

	o.Insert(&adInfo)

	this.Success("success")
}



// @Title 添加图片
// @Description 添加图片
// @Success 200 {object} nil
// @Param    id    query    int    false  标题
// @router / [delete]
func (this* AdController)Delete(){
	id,_ := this.GetInt("id",-1)

	o := orm.NewOrm()
	qs := o.QueryTable("item_info")
	qs = qs.Filter("id",id)

	_,err := qs.Delete()

	if err != nil {
		this.Error(errorcode.ERROR_PARAM,"删除对象下载错误")
	}

	this.Success(nil)
}


// @Title 编辑节点
// @Description 编辑节点
// @Success 200 {object} info.adInfo
// @Param    Id           formData    int       true    主键
// @Param    Title    query    string    false  标题
// @Param    Type       query    int    true   类型
// @Param    Address   query    string    true   多个地址，以逗号分隔
// @Param    Link   query    string    true   多个地址，以逗号分隔
// @router   / [put]
func (this* AdController)Edit(){
	id,_ := this.GetInt64("id",-1)
	if id == -1 {
		this.Error(errorcode.ERROR_PARAM)
	}

	var req dto.ReqAdInfo
	if err := this.ParseForm(&req);err != nil {
		this.Error(errorcode.ERROR_VALID)
	}

	valid := validation.Validation{}
	b,err := valid.Valid(&req)
	if err != nil {
		this.Error(errorcode.ERROR_VALID)
	}

	if !b {
		this.Error(errorcode.ERROR_VALID)
	}

	db := tsDb.NewDbBase()
	info := info.AdInfo{
		Id:          id,
		Title:       req.Title,
		Address:      req.Address,
		Type:        req.Type,
		UpdateTime:  tsTime.CurrSe(),
	}


	err =db.DbUpdate(&info,"Id","Title","Address","Type","UpdateTime")
	if err != nil {
		logs.Error(err)
		this.Error(tsOpCode.OPERATION_DB_FAILED)
	}

	this.Success("success")
}


// @Title 编辑节点
// @Description 编辑节点
// @Success 200 {object} info.adInfo
// @Param    Id           query    int       true    主键
// @router   / [get]
func (this* AdController)Get(){
	db := tsDb.NewDbBase()
	info := info.AdInfo{}
	info.Id,_ = this.GetInt64("id",0)

	err := db.DbGet(&info)
	if err != nil {
		logs.Error(err)
	}

	this.Success(info)
}


// @Title 请求各个分类的索引
// @Description 请求各个分类的索引 这里后期需要改成缓存读取方式，从redis中读取相关的数据，隔一段时间更新一次
// @Success 200 [{Object}]  info.adInfo
// @Parrouteram    :tp   		 path    int    true  			1:视频，2：电影 3：图片 4：小说 5：博彩
// @Parrouteram    count    	query    int    true			//需要请求的数据个数
// @Parrouteram    offset    	query    int    true			//请求数据的偏移量
// @router /:tp [get]
func (this* AdController)List(){
	tp,_ := this.GetInt64(":tp", 1)
	count,_ := this.GetInt32("count",10)
	offset,_:= this.GetInt32("offset",0)

	o := orm.NewOrm()
	qs := o.QueryTable("ad_info")
	qs = qs.Filter("Type",tp)

	qs = qs.OrderBy("CrtTime")
	qs = qs.Limit(count,offset)

	data := []orm.Params{}

	_,err :=qs.Values( &data)		// 请求所有

	if err != nil {
		logs.Error(err)
	}

	this.Success(data)
}