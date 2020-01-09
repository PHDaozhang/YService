package common

import (
	"YService/controllers/system"
	"YService/models/common"
	"github.com/astaxie/beego/orm"
)

type Header struct {
	system.BaseController
}



// @Title 获取当前的最新的资源信息，包括图片，影片
// @Description 添加图片
// @Success 200 [{Object}]  FreshItemInfo
// @Parrouteram    movieType    query    int    false  类型，1：电影 2：视频
// @ /fresh [post]
func (this* Header)Fresh(){
	movieType,_ := this.GetInt("movieType",1)

	o := orm.NewOrm()
	resp := new(common.ItemInfo)
	qs := o.QueryTable(resp)
	qs.Filter("movie_type",movieType)
	qs.OrderBy("create_time")
	qs.Limit(10)

	//qs

	//tsCrypto.EncodeBase64()


}