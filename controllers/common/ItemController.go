package common

import (
	"YService/controllers/system"
	"YService/core/errorcode"
	"YService/models/dto"
	"YService/models/info"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strings"
	"tsEngine/tsDb"
	"tsEngine/tsOpCode"
	"tsEngine/tsTime"
)

type ItemController struct {
	system.BaseController
}

// @Title 添加图片
// @Description 添加图片
// @Success 200 {object} respSonList
// @Param    Title    formdata    string    false  标题
// @Param    Type       formdata    int    true   类型
// @Param    Address   formdata    string    true   多个地址，以逗号分隔
// @Param    ChildType   formdata    string    true   多个地址，以逗号分隔
// @router / [post]
func (this* ItemController)Add(){
	var reqItemDto dto.ReqItemInfo

	if err := this.ParseForm(&reqItemDto); err != nil {
		this.Error(errorcode.ERROR_PARAM)
	}

	valid := validation.Validation{}
	b,err := valid.Valid(&reqItemDto)
	if err != nil {
		this.Error(errorcode.ERROR_VALID)
	}

	if !b {
		this.Error(errorcode.ERROR_VALID)
	}

	o := orm.NewOrm()

	strList := strings.Split(reqItemDto.Address,",")

	itemInfo := info.ItemInfo{}
	itemInfo.Title = reqItemDto.Title
	itemInfo.Address = reqItemDto.Address
	itemInfo.Type = reqItemDto.Type
	itemInfo.ChildType = reqItemDto.ChildType
	itemInfo.CrtTime = tsTime.CurrSe()
	itemInfo.FocusNum = 0
	itemInfo.Num = int32(len(strList))

	o.Insert(&itemInfo)

	this.Success(nil)
}


// @Title 添加图片
// @Description 添加图片
// @Success 200 {object} respSonList
// @Param    id    query    int    false  标题
// @router / [delete]
func (this* ItemController)Delete(){
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
// @Success 200 {object} admin.Node
// @Param    Id           query    int       true    主键
// @Param    Title    query    string    false  标题
// @Param    Type       query    int    true   类型
// @Param    Address   query    string    true   多个地址，以逗号分隔
// @Param    ChildType   query    string    true   多个地址，以逗号分隔
// @router   / [put]
func (this* ItemController)Edit(){
	id,_ := this.GetInt64("id",-1)
	if id == -1 {
		this.Error(errorcode.ERROR_PARAM)
	}

	var req dto.ReqItemInfo
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
	info := info.ItemInfo{
		Id:          id,
		Title:       req.Title,
		Address:      req.Address,
		Type:        req.Type,
		ChildType:   req.ChildType,
		UpdateTime:  tsTime.CurrSe(),
		MovieHeader: req.MovieHeader,
	}

	strList := strings.Split(info.Address,",")
	if len(strList) == 0 {
		this.Error(errorcode.ERROR_VALID,"地址为空")
		info.Num = int32(len(strList))
	}

	err =db.DbUpdate(&info,"Id","Title","Address","Type","ChildType","UpdateTime","MovieHeader")
	if err != nil {
		logs.Error(err)
		this.Error(tsOpCode.OPERATION_DB_FAILED)
	}



	this.Success("success")
}

// @Title 编辑节点
// @Description 编辑节点
// @Success 200 {object} info.adInfo
// @Param    Id           formData    int       true    主键
// @router   / [get]
func (this* ItemController)Get(){

	db := tsDb.NewDbBase()

	info := info.ItemInfo{}
	info.Id,_ = this.GetInt64("id",0)

	err := db.DbGet(&info)
	if err != nil {
		logs.Error(err)
	}

	this.Success(info)
}

// @Title 请求各个分类的索引
// @Description 请求各个分类的索引 这里后期需要改成缓存读取方式，从redis中读取相关的数据，隔一段时间更新一次
// @Success 200 [{Object}]  FreshItemInfo
// @Parrouteram    :tp   		 path    int    true  			1:视频，2：电影 3：图片 4：小说 5：博彩
// @Parrouteram    :childTp    	 path    int    true			//每个分类别下面的子类型
// @Parrouteram    count    	query    int    true			//需要请求的数据个数
// @Parrouteram    offset    	query    int    true			//请求数据的偏移量
// @router /:tp/:childTp [get]
func (this* ItemController)List(){
	tp,_ := this.GetInt64(":tp", 1)
	childTp,_ := this.GetInt64(":childTp", 0)
	limit,_ := this.GetInt32("PageSize",10)
	offset,_:= this.GetInt32("Page",0)

	o := orm.NewOrm()
	qs := o.QueryTable("item_info")
	qs = qs.Filter("Type",tp)

	//如果小于或是等于0，则代表是这个类别下面的所有的数据
	if childTp > 0 {
		qs = qs.Filter("ChildType",childTp)
	}

	qs = qs.OrderBy("CrtTime")
	qs = qs.Limit(limit,offset * limit)

	list := []orm.Params{}

	_,err := qs.Values( &list)		// 请求所有

	if err != nil {
		logs.Info("go this...")
	}

	tsConn := tsDb.NewDbBase()
	itemInfo := info.ItemInfo{}
	dbCount := int64(0)
	if childTp > 0 {
		dbCount,_= tsConn.DbCount(&itemInfo,"Type",tp,"ChildType",childTp)
	} else {
		dbCount,_= tsConn.DbCount(&itemInfo,"Type",tp)
	}


	this.Success(beego.M{"count":dbCount,"list":list})
}