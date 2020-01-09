package routers

import (
	"YService/controllers"
	"YService/controllers/admin"
	"YService/controllers/sport"
	"github.com/astaxie/beego"
)

func init() {


	ysApi := beego.NewNamespace("/api",
		beego.NSNamespace("/admin",beego.NSInclude(&admin.AdminController{})),
		)
	beego.AddNamespace(ysApi)


	sport188Api := beego.NewNamespace("/API",
		beego.NSNamespace("/test",beego.NSInclude(&sport.Sport188Controller{})),
	)
	beego.AddNamespace(sport188Api)

	beego.Router("/", &controllers.MainController{})
}


//bee run -gendoc=true -downdoc=true

/**
@Title
这个 API 所表达的含义，是一个文本，空格之后的内容全部解析为 title

@Description
这个 API 详细的描述，是一个文本，空格之后的内容全部解析为 Description

@Param
参数，表示需要传递到服务器端的参数，有五列参数，使用空格或者 tab 分割，五个分别表示的含义如下
参数名
参数类型，可以有的值是 formData、query、path、body、header，formData 表示是 post 请求的数据，query 表示带在 url 之后的参数，path 表示请求路径上得参数，例如上面例子里面的 key，body 表示是一个 raw 数据请求，header 表示带在 header 信息中得参数。
参数类型
是否必须
注释

@Success
成功返回给客户端的信息，三个参数，第一个是 status code。第二个参数是返回的类型，必须使用 {} 包含，第三个是返回的对象或者字符串信息，如果是 {object} 类型，那么 bee 工具在生成 docs 的时候会扫描对应的对象，这里填写的是想对你项目的目录名和对象，例如 models.ZDTProduct.ProductList 就表示 /models/ZDTProduct 目录下的 ProductList 对象。
三个参数必须通过空格分隔

@Failure
失败返回的信息，包含两个参数，使用空格分隔，第一个表示 status code，第二个表示错误信息

@router
路由信息，包含两个参数，使用空格分隔，第一个是请求的路由地址，支持正则和自定义路由，和之前的路由规则一样，第二个参数是支持的请求方法,放在 [] 之中，如果有多个方法，那么使用 , 分隔
 */