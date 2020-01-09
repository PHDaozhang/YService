package dto

type RespFreshItemInfo struct {
	Id                      uint64				//唯一ID
	Title          			string 				//标题
	CreateTime				uint64				//创建时间
	PicHeader				string				//电影的图片地址
	ShortUrl				string				//电影的地址短信息
}