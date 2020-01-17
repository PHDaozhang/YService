package info

import "github.com/astaxie/beego/orm"

type ItemInfo struct {
	Id            	int64  `orm:"description(主键)"`
	Title         	string  `orm:"description(标题)"`
	CrtTime 	  	uint64  `orm:"description(创建时间)"`
	UpHost    	  	string `orm:"description(上传人)"`
	Address    	  	string `orm:"description(地址，多个图片使用逗号分隔)"`
	FocusNum      	int32    `orm:"description(浏览的次数)"`
	Type       	  	int32    `orm:"description(1:视频，2：电影 3：图片 4：小说 5：游戏 )"`
	Num        		int32    `orm:"description(资源的数量，只有在图片时此值 才有效，另两种类型此值一直为1)"`
	ChildType		int32	`orm:"description(每个类型的子类别说明)"`
	UpdateTime		uint64	`orm:"description(更新时间戳)"`

	//只有电影或视频
	MovieHeader		string	`orm:"description(当为视频与电影时，其的简写数据)"`		//当为视频或是电影时 显示在页面上的索引图片
	//FormatType		int32	`orm:"description(mp4,armt,flv 等样式)"`
}

func init(){
	orm.RegisterModel(new(ItemInfo))
}

func (this* ItemInfo)TableName() string {
	return "item_info"
}