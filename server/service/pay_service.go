package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ppoonk/AirGo/global"
	"github.com/ppoonk/AirGo/model"
	"github.com/ppoonk/AirGo/utils/encrypt_plugin"
	"github.com/ppoonk/AirGo/utils/net_plugin"
	"github.com/smartwalle/alipay/v3"
)

// 支付宝-alipay初始化
func InitAlipayClient(pay model.Pay) (*alipay.Client, error) {
	//false时用开发网关，https://openapi.alipaydev.com/gateway.do；true时用正式环境网关，https://openapi.alipay.com/gateway.do
	client, err := alipay.New(pay.AliPay.AlipayAppID, pay.AliPay.AlipayAppPrivateKey, true)
	if err != nil {
		//fmt.Println("初始化支付宝失败, 错误信息为", err)
		//os.Exit(-1)
		return nil, err
	}

	// 加载内容密钥（可选），详情查看 https://opendocs.alipay.com/common/02mse3
	client.SetEncryptKey(pay.AliPay.AlipayEncryptKey)

	// 下面两种方式只能二选一
	var cert = false
	if cert {
		// 使用支付宝证书
		fmt.Println("加载证书", client.LoadAppCertPublicKeyFromFile("appPublicCert.crt"))
		fmt.Println("加载证书", client.LoadAliPayRootCertFromFile("alipayRootCert.crt"))
		fmt.Println("加载证书", client.LoadAlipayCertPublicKeyFromFile("alipayPublicCert.crt"))
	} else {
		// 使用支付宝公钥
		fmt.Println("加载公钥", client.LoadAliPayPublicKey(pay.AliPay.AlipayAliPublicKey))
	}
	return client, nil
}

// 支付宝-统一收单线下交易预创建,生成二维码后，展示给用户，由用户扫描二维码完成订单支付（当面付）
func TradePreCreatePay(client *alipay.Client, sysOrder *model.Orders) (*alipay.TradePreCreateRsp, error) {
	//创建支付宝订单 请求模板
	// order := alipay.TradePreCreate{
	// 	Trade: alipay.Trade{
	// 		//NotifyURL:  "" ,  //异步通知的URL，该URL是支付宝服务器端自动调用商户服务端接口的地址，支付成功后调用，再根据支付宝转发的参数进行订单状态处理,异步的post请求
	// 		//ReturnURL: "http://", //同步返回URL，是一个页面跳转通知（支付成功后，从支付宝跳转到指定的地址）。同步的get请求
	// 		//AppAuthToken: "", // 可选

	// 		// biz content，这四个参数是必须的
	// 		Subject:     "一个馒头",                 // 订单标题
	// 		OutTradeNo:  "1008610010",           // 商户订单号，64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	// 		TotalAmount: "0.01",                 // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	// 		ProductCode: "FACE_TO_FACE_PAYMENT", // 销售产品码，与支付宝签约的产品码名称。 参考官方文档,
	// 		//App 支付时默认值为 QUICK_MSECURITY_PAY
	// 		//手机网站支付产品alipay.trade.wap.pay接口中，product_code为：QUICK_WAP_WAY
	// 		//电脑网站支付产品alipay.trade.page.pay接口中，product_code为：FAST_INSTANT_TRADE_PAY
	// 		//当面付条码支付产品alipay.trade.pay接口中，product_code为：FACE_TO_FACE_PAYMENT
	// 	},
	// }
	//创建支付宝订单
	var order alipay.TradePreCreate
	//order.NotifyURL = global.Server.AliPaySetting.ReturnURL  //支付结果放在轮询里判断
	order.Subject = sysOrder.Subject
	order.OutTradeNo = sysOrder.OutTradeNo
	order.TotalAmount = sysOrder.TotalAmount
	order.ProductCode = "FACE_TO_FACE_PAYMENT"
	res, err := client.TradePreCreate(order)
	//fmt.Println("TradePreCreate:", res, err)
	return res, err
	//响应模板
	// 	{
	// 	"code": 0,
	// 	"msg": "alipay TradePreCreatePay success:",
	// 	"data": {
	// 		"alipay_trade_precreate_response": {
	// 			"code": "10000",
	// 			"msg": "Success",
	// 			"sub_code": "",
	// 			"sub_msg": "",
	// 			"out_trade_no": "5",
	// 			"qr_code": "https://qr.alipay.com/bax07220zdz0k58x5abw2504"
	// 		},
	// 		"sign": "EmZmz7Jix2GLtScaysE9D0DF9Sw9ZDuuums7CXywFO83g/dnOasZiAQnDhsgoMq9JmPnygIog4+myEcxXqmoLM2qZX2zy3Aof7CbVzLwA931kq09u6y54h28R+BvILLZAR5gmSYW2oh4/gWO24yK8awHLndCAQhNuHFOkMwCAcDRKGjhKkDb9XIx/do99V/Xa9w8pJhHSt1ONaIjyWufK6b4YcVg3bGldBTG+xpqDvzXSYFc5lBRfgAJxn8NklTKVj/PLFr3nM4IJ/fYFaJuHS2/pjQThyDiEsjPvEhA9aPEeXK03J8Qea0HFAuM9i2kw1OqmeN0oiHCrVVSCFGPRg=="
	// 	}
	// }

}

// 支付宝-统一收单线下交易查询
func TradeQuery(client *alipay.Client, sysOrder *model.Orders) (*alipay.TradeQueryRsp, error) {
	// TradeQuery 统一收单线下交易查询接口请求参数 https://docs.open.alipay.com/api_1/alipay.trade.query/
	//type TradeQuery struct {
	//	AppAuthToken string   `json:"-"`                       // 可选
	//	OutTradeNo   string   `json:"out_trade_no,omitempty"`  // 订单支付时传入的商户订单号, 与 TradeNo 二选一
	//	TradeNo      string   `json:"trade_no,omitempty"`      // 支付宝交易号
	//	QueryOptions []string `json:"query_options,omitempty"` // 可选 查询选项，商户通过上送该字段来定制查询返回信息 TRADE_SETTLE_INFO(交易结算信息)
	//}
	var p = alipay.TradeQuery{
		OutTradeNo: sysOrder.OutTradeNo,
		//QueryOptions: []string{"TRADE_SETTLE_INFO"},
	}
	//fmt.Println("统一收单线下交易查询 p:", p)
	rsp, err := client.TradeQuery(p)
	return rsp, err

	//  if rsp.Content.Code != alipay.CodeSuccess
	//	fmt.Println(rsp.Content.Code, rsp.Content.Msg, rsp.Content.SubMsg)
	//Code   Msg                SubCode                    SubMsg       TradeNo                       OutTradeNo                BuyerLogonId   TradeStatus
	//40004 Business  Failed    ACQ.TRADE_NOT_EXIST        交易不存在                                  168425758005579100000000
	//10000 Success                                                     2023051822001475841447588320  168438189998030100010000 249***@qq.com   WAIT_BUYER_PAY
	//10000 Success                                                     2023051822001475841447588320  168438189998030100010000 249***@qq.com   TRADE_SUCCESS

}

// 支付宝-统一收单交易关闭接口
func TradeClose(client *alipay.Client, sysOrder *model.Orders) (*alipay.TradeCloseRsp, error) {
	//type TradeClose struct {
	//	AppAuthToken string `json:"-"`                      // 可选
	//	NotifyURL    string `json:"-"`                      // 可选
	//	OutTradeNo   string `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	//	TradeNo      string `json:"trade_no,omitempty"`     // 与 OutTradeNo 二选一
	//	OperatorId   string `json:"operator_id,omitempty"`  // 可选
	//}
	var p = alipay.TradeClose{
		OutTradeNo: sysOrder.OutTradeNo,
	}
	rsp, err := client.TradeClose(p)
	return rsp, err
}

// 支付宝-轮询
func PollAliPay(order *model.Orders, client *alipay.Client) {
	t := time.NewTicker(10 * time.Second)
	//defer t.Stop()
	i := 0
	for {
		if i == 18 { // 18*10s 3分钟，3分钟未付款则超时取消交易
			if order.TradeNo != "" {
				res, _ := TradeClose(client, order) //超时，取消订单
				//fmt.Println("支付宝取消订单结果:", res)
				global.Logrus.Error("支付宝取消订单结果:", res)
			}
			order.TradeStatus = "TRADE_CLOSED" //更新数据库订单状态(超时已取消)
			UpdateOrder(order)                 //更新数据库状态
			t.Stop()
			return
		}
		<-t.C
		rsp, _ := TradeQuery(client, order)
		//fmt.Println("支付宝TradeQuery rsp.Content.TradeStatus:", rsp.TradeStatus)
		if rsp.TradeStatus == "TRADE_SUCCESS" || rsp.TradeStatus == "TRADE_FINISHED" { //交易结束
			if global.Server.Subscribe.EnabledRebate {
				global.GoroutinePool.Submit(func() {
					ReferrerRebate(order.UserID, rsp.ReceiptAmount) //处理推荐人返利
				})
			}
			order.TradeStatus = "TRADE_SUCCESS"       //交易成功
			order.BuyerLogonId = rsp.BuyerLogonId     //买家支付宝账号
			order.ReceiptAmount = rsp.ReceiptAmount   //实收金额
			order.BuyerPayAmount = rsp.BuyerPayAmount //付款金额
			PaymentSuccessfullyOrderHandler(order)
			t.Stop()
			return
		}
		if rsp.TradeStatus == "WAIT_BUYER_PAY" && order.TradeStatus != "WAIT_BUYER_PAY" { //等待付款
			order.TradeNo = rsp.TradeNo
			UpdateOrder(order) //更新数据库状态
		}
		i++
	}
}

// 易支付-交易预创建（api支付）（弃用）
func EpayPreByApi(epayConfig model.Epay, sysOrder *model.Orders) (model.EpayPreCreatePayResponse, error) {
	var epayRes model.EpayPreCreatePayResponse

	client := net_plugin.ClientWithDNS("114.114.114.114", 10*time.Second)

	var formValues url.Values
	formValues.Add("pid", strconv.FormatInt(epayConfig.EpayPid, 10))
	formValues.Add("type", epayConfig.EpayType) //支付方式, alipay	支付宝 wxpay	微信支付 qqpay	QQ钱包 bank	网银支付
	formValues.Add("out_trade_no", sysOrder.OutTradeNo)
	formValues.Add("notify_url", epayConfig.EpayNotifyURL)
	formValues.Add("return_url", epayConfig.EpayReturnURL)
	formValues.Add("name", sysOrder.Subject)
	formValues.Add("money", sysOrder.Price)
	//formValues.Add("clientip", "")
	//formValues.Set("device", "")
	//formValues.Set("param", "")
	formValues.Add("sign", epayConfig.EpayKey)
	formValues.Add("sign_type", encrypt_plugin.Md5Encode(formValues.Encode()+epayConfig.EpayKey, false))

	reader := strings.NewReader(formValues.Encode())

	req, err := http.NewRequest("POST", epayConfig.EpayApiURL, reader)
	if err != nil {
		return epayRes, err
	}
	req.Header.Set("Accept", "application/x-www-form-urlencode")
	//fmt.Println("请求参数：", req)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return epayRes, err
	}
	out := net_plugin.ReadDate(resp)

	err = json.Unmarshal([]byte(out), &epayRes)
	if err != nil {
		return epayRes, err
	}
	//fmt.Println("响应：", epayRes)
	return epayRes, nil

}

// 易支付-交易预创建（网页支付）
func EpayPreByHTML(sysOrder *model.Orders, pay *model.Pay) (*model.EpayPreCreatePayToFrontend, error) {
	var epayRsp model.EpayPreCreatePayToFrontend
	var epay = model.EpayPreCreatePay{
		Pid:        pay.Epay.EpayPid,
		Type:       "", //为空则直接跳转到易支付收银台
		OutTradeNo: sysOrder.OutTradeNo,
		NotifyUrl:  global.Server.Subscribe.BackendUrl + "/api/public/epayNotify",
		ReturnUrl:  global.Server.Subscribe.BackendUrl + "/#/home",
		Name:       sysOrder.Subject,
		Money:      sysOrder.Price,
		//ClientIP:   "",
		//Device:     "",
		//Param:      "",
		//Sign:     pay.Epay.EpayKey,
		SignType: "MD5",
	}
	epay.Sign = CreateEpaySign(&epay, pay)
	epayRsp.EpayApiURL = pay.Epay.EpayApiURL
	epayRsp.EpayPreCreatePay = epay
	//fmt.Println("epay:", epay)
	return &epayRsp, nil

}

// 易支付sign生成
func CreateEpaySign(epay *model.EpayPreCreatePay, pay *model.Pay) string {
	text := "money=" + epay.Money + "&" + "name=" + epay.Name + "&" + "notify_url=" + epay.NotifyUrl + "&" + "out_trade_no=" + epay.OutTradeNo + "&" + "pid=" + strconv.FormatInt(epay.Pid, 10) + "&" + "return_url=" + epay.ReturnUrl + pay.Epay.EpayKey
	//fmt.Println("text:", text)
	return encrypt_plugin.Md5Encode(text, false)
}
