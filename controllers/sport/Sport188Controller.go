package sport

import (
	"YService/controllers/system"
)

type Sport188Controller struct {
	system.BaseController
}

//注单数据
type wagers struct {
	Id							uint64			`json:"id"`
	MemberCode					string 		`json:"memberCode"`
	CurrencyCode				string		`json:"currencyCode"`
	IpAddress					string		`json:"ipAddress"`
	CreateTime					string		`json:"createTime"`
	SettleTime					string		`json:"settleTime"`
	WagerStatus					int			`json:"wagerStatus"`
	Stake						float32		`json:"stake"`
	ReturnAmount				float32		`json:"returnAmount"`
	Channel                     int			`json:"channel"`
	OddType						int			`json:"oddType"`
	Odds						float32		`json:"odds"`
	EventStartTime				string 		`json:"eventStartTime"`
	Prefix                      string		`json:"prefix"`
	BetType						string 		`json:"betType"`
	ActiveBet					string		`json:"activeBet"`
	OrderDate					string 		`json:"orderDate"`
	Bets						[]Bet		`json:"bets"`
}

type Bet struct {
	Id								uint64		`json:"id"`
	Seq								int32 		`json:seq`
	BetStatus						int			`json:"betStatus"`
	Stake							int			`json:"stake"`
	ReturnAmount					float32		`json:"returnAmount"`
	Odds							float32		`json:"odds"`
	BetType							string 		`json:"betType"`
	Selection                 		string 		`json:"selection"`
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

	isSettled,_ := this.GetInt("isSettled",1)

	if isSettled == 1 {
		//如果是1：未结账，2：已结账
		for i := 0; i < 10;i ++ {
			item := wagers{}
			item.Id = 39092561659302 + uint64(i)
			item.MemberCode = "yt_2045182"
			item.CurrencyCode = "CNY"
			item.IpAddress = "18.162.173.173"
			item.CreateTime = "2010-01-10 01:59:29"
			item.SettleTime = "2010-01-10 17:53:54"
			item.WagerStatus = 2
			item.Stake = 10.00
			item.ReturnAmount = 19.80
			item.Channel = 10
			item.OddType  = 1
			item.Odds = 1.980
			item.Prefix = "yt"
			item.BetType = "1x2"
			item.ActiveBet = "0.00"
			item.OrderDate = "2020-01-10 16:00:00"

			bets := []Bet{}
			//只有一个数据的情况
			bet := Bet{}
			bet.Id = 4774517234
			bet.Seq = 1
			bet.BetStatus = 1
			bet.Stake = 10
			bet.ReturnAmount = 19.8
			bet.Odds = 1.98
			bet.BetType = "1x2"



			bets = append(bets,bet)
			item.Bets = bets

			retMap = append(retMap, item)
		}

	} else {
		for i := 0; i < 2;i ++ {
			item := wagers{}
			item.Id = 39092561659302 + uint64(i)
			item.MemberCode = "yt_2045182"
			item.CurrencyCode = "CNY"
			item.IpAddress = "18.162.173.173"
			item.CreateTime = "2010-01-10 01:59:29"
			item.SettleTime = "0000-00-00 00:00:00"
			item.WagerStatus = 1
			item.Stake = 10.00
			item.ReturnAmount = 0.00
			item.Channel = 10
			item.OddType  = 1
			item.Odds = 1.980
			item.Prefix = "yt"
			item.BetType = "1x2"
			item.ActiveBet = "0.00"
			item.OrderDate = "2020-01-10 16:00:00"

			bets := []Bet{}
			//只有一个数据的情况
			bet := Bet{}
			bet.Id = 4774517234
			bet.Seq = 1
			bet.BetStatus = 1
			bet.Stake = 10
			bet.ReturnAmount = 19.8
			bet.Odds = 1.98
			bet.BetType = "1x2"



			bets = append(bets,bet)
			item.Bets = bets

			retMap = append(retMap, item)
		}
	}

	this.Result = retMap


	this.TraceSportJson();
}

func (this *Sport188Controller) TraceSportJson() {
	this.Data["json"] = &map[string]interface{}{"code": "COMM0000", "msg": "Success", "data": this.Result,"carrier":map[string]interface{}{}}
	this.ServeJSON()
	this.StopRun()
}
