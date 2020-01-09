package dto


type ReqItemInfo struct {
	Title					string	`form:"title"`			//标题
	Type 					int32 							//资源类型
	Address					string							//多个地址，以逗号分隔
	UpHost					string
	ChildType				int32							//子类型的类别
	MovieType				int32							//影片的类型	1：电影	2：视频
	MovieHeader				string							//影片的索引的图片地址
	FormatType				int32							//图片或是影片的格式
}
