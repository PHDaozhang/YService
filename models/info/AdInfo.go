package info

import "github.com/astaxie/beego/orm"

type AdInfo struct {
	Id            	int64  `orm:"description(主键)"`
	Title         	string  `orm:"description(标题)"`
	Type			int32 	`orm:"description(类型)"`					//主页面上面的三个滚动图片1：最上面，2：中间 3：最下面
	CrtTime 	  	uint64  `orm:"description(创建时间)"`
	UpHost    	  	string `orm:"description(上传人)"`
	Address    	  	string `orm:"description(单个地址信息)"`
	UpdateTime		uint64	`orm:"description(更新时间戳)"`
	Link			string `orm:"description(连接地址)"`
}

func init(){
	orm.RegisterModel(new(AdInfo))
}

func (this* AdInfo)TableName() string {
	return "ad_info"
}