package sport

import (
	"YService/controllers/system"
	"strconv"
)

type Sport188Controller struct {
	system.BaseController
}

//注单数据
type wagers struct {
	Id							int			`json:"id"`
	MemberCode					string 		`json:"memberCode"`
	CurrencyCode				string		`json:"currencyCode"`
}


// @Title 添加图片
// @Description 添加图片
// @Success 200 {object} respSonList
// @router /Launch [post]
func (this* Sport188Controller)Launch(){
	this.Result = "http://www.baidu.com"
	this.TraceSportJson();
}


// @Title 添加图片
// @Description 添加图片
// @Success 200 {object} respSonList
// @router /Logout [post]
func (this* Sport188Controller)Logout(){
	this.Result = map[string]interface{}{}
	this.TraceSportJson();
}

// @Title 添加图片
// @Description 添加图片
// @Success 200 {object} respSonList
// @router /Registration [post]
func (this* Sport188Controller)Registration(){
	this.Result = map[string]interface{}{}
	this.TraceSportJson();
}


// @Title 添加图片
// @Description 添加图片
// @Success 200 {object} respSonList
// @router /MemberBalance [post]
func (this* Sport188Controller)MemberBalance(){
	retMap :=  map[string]interface{}{}
	retMap["currencyCode"] = "CNY"
	retMap["balance"] = 1000.25

	this.Result = retMap


	this.TraceSportJson();
}

// @Title 添加图片
// @Description 添加图片
// @Success 200 {object} respSonList
// @router /DepositFund [post]
func (this* Sport188Controller)DepositFund(){
	retMap :=  map[string]interface{}{}
	retMap["totalBalance"] = 1000.25
	retMap["transationId"] = 32154564

	this.Result = retMap


	this.TraceSportJson();
}


// @Title 添加图片
// @Description 添加图片
// @Success 200 {object} respSonList
// @router /WithdrawalFund [post]
func (this* Sport188Controller)WithdrawalFund(){


	retMap :=  map[string]interface{}{}
	retMap["totalBalance"] = 1000.25
	retMap["transationId"] = 32154564

	this.Result = retMap


	this.TraceSportJson();
}


// @Title 添加图片
// @Description 添加图片
// @Success 200 {object} respSonList
// @router /Wagers [post]
func (this* Sport188Controller)Wagers(){
	retMap := []wagers{}

	for i := 0; i < 10;i ++ {
		item := wagers{}
		item.Id = 123456464 + i
		item.MemberCode = "yt_abc" + strconv.Itoa(i)
		item.CurrencyCode = "CNY"

		retMap = append(retMap, item)
	}

		this.Result = retMap


	this.TraceSportJson();
}

func (this *Sport188Controller) TraceSportJson() {
	this.Data["json"] = &map[string]interface{}{"code": "COMM000", "msg": "Success", "data": this.Result,"carrier":map[string]interface{}{}}
	this.ServeJSON()
	this.StopRun()
}
