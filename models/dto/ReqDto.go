package dto

//每个类型的数据
type ReqItemInfo struct {
	Title					string	`form:"title" valid:"Required"`			//标题
	Type 					int32	`form:"type" valid:"Required;Range(1,5)"` 							//资源类型
	Address					string	`form:"address" valid:"Required"`						//多个地址，以逗号分隔
	UpHost					string	`form:"upHost"`
	ChildType				int32	`form:"childType" valid:"Required;Range(1,8)"`							//子类型的类别
	MovieHeader				string	`form:"movieHeader"`						//影片的索引的图片地址
}

//广告请求数据
type ReqAdInfo struct {
	Title					string	`form:"title" valid:"Required"`			//标题
	Address					string	`form:"address" valid:"Required"`						//多个地址，以逗号分隔
	Type					int32	`form:"type" valid:"Required"`						//数据的类型
	UpHost					string	`form:"upHost"`
	Link					string	`form:"link" valid:"Required"`						//影片的索引的图片地址
}
