package admin

import (
	"YService/controllers/system"
	"YService/core/errorcode"
	"YService/models/common"
	"YService/models/dto"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strings"
	"tsEngine/tsTime"
)

type AdminController struct {
	system.BaseController
}



// @Title 添加图片
// @Description 添加图片
// @Success 200 {object} respSonList
// @Param    Title    query    string    false  标题
// @Param    Type       query    int    true   类型
// @Param    Address   query    string    true   多个地址，以逗号分隔
// @router /addItem [post]
func (this* AdminController)AddItem(){
	var reqItemDto dto.ReqItemInfo

	if err := this.ParseForm(&reqItemDto); err != nil {
		this.Error(errorcode.PARAM_ERROR)
	}

	valid := validation.Validation{}
	b,err := valid.Valid(&reqItemDto)
	if err != nil {
		this.Error(errorcode.VALID_ERROR)
	}

	if !b {
		this.Error(errorcode.VALID_ERROR)
	}

	o := orm.NewOrm()

	strList := strings.Split(reqItemDto.Address,",")

	itemInfo := common.ItemInfo{}
	itemInfo.Addres = reqItemDto.Address
	itemInfo.CrtTime = tsTime.CurrSe()
	itemInfo.Title = reqItemDto.Title
	itemInfo.Type = reqItemDto.Type
	itemInfo.FocusNum = 0
	itemInfo.Num = int32(len(strList))
	itemInfo.ChildType = reqItemDto.ChildType
	itemInfo.MovieType = reqItemDto.MovieType


	o.Insert(&itemInfo)

	this.Success(1)
}


// @Title 添加图片
// @Description 添加图片
// @Success 200 {object} respSonList
// @Param    id    query    int    false  标题
// @router /deleteItem [post]
func (this* AdminController)DeleteItem(){
	id,_ := this.GetInt("id",-1)

	op := orm.NewOrm()
	res,err	:= op.Raw("delete from `item_info` where `id = ? ` " + fmt.Sprintf("%d",id)).Exec()
	if err != nil {
		logs.Error("....错误",err)
		return
	}

	_,err = res.RowsAffected()

	return
}