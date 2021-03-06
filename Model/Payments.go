package Model

import "C"
import (
	"bytes"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"projectExample/Config"
	"sync"
)
var Mutex sync.Mutex
func AddAmount(addReq AddRequest, Userid string)(addRes *AddResponse, err error){
	addAmount := addReq.Amount
	var tradingAcc TradingAccount
	var addResponse AddResponse
	var razorpayRes RazorpayResponse
	var payments Payments
	var callbackResponse CallbackResponse
	if err = Config.DB.Table("trading_account").Where("user_id = ?", Userid).First(&tradingAcc).Error; err != nil {
		return nil, err
	}
	razorRequest := RazorpayRequest{Amount: addAmount, CallbackURL: "http://localhost:8080/payments/payment_status", CallbackMethod: "get"}
	jsonReq, err := json.Marshal(razorRequest)
	resp, err := http.Post("https://api.razorpay.com/v1/payment_links", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &razorpayRes)
	if err != nil {
		return nil, err
	}
	pay := Payments{UserId: Userid,Amount: addAmount,RazorpayLink: razorpayRes.ShortURL ,RazorpayLinkId: razorpayRes.ID}
	if err = Config.DB.Create(pay).Error; err != nil {
		return nil, err
	}
	if err = Config.DB.Where("user_id = ?, razorpay_link_id=? ,razorpay_link=?", Userid,razorpayRes.ShortURL,razorpayRes.ID).First(&callbackResponse).Error; err != nil {
		return nil, err
	}
	if callbackResponse.RazorpayPaymentLinkStatus == "success" {
		tradingAcc.Balance += addAmount
	}
	if err = Config.DB.Table("trading_account").Where("user_id = ?", Userid).UpdateColumn("balance", tradingAcc.Balance).Error; err != nil{
		return nil, err
	}
	if err = Config.DB.Table("payments").Where("user_id = ?, razorpay_link_id=? ,razorpay_link=?", Userid,razorpayRes.ShortURL,razorpayRes.ID).UpdateColumn("current_balance", tradingAcc.Balance).Error; err != nil{
		return nil, err
	}
	addResponse.Userid = payments.UserId
	addResponse.Amount = payments.Amount
	addResponse.Type = "add"
	addResponse.CurrentBalance = payments.CurrentBalance
	addResponse.Message = "Process Successful"
	return &addResponse,err

}
func WithdrawAmount(withdrawReq WithdrawRequest, Userid string )(withdrawRes *WithdrawResponse, err error){
	withdrawAmount := withdrawReq.Amount
	var tradingAcc TradingAccount
	var withdrawResponse WithdrawResponse
	if err = Config.DB.Table("trading_account").Where("user_id = ?", Userid).First(&tradingAcc).Error; err != nil {
		return nil, err
	}
	if tradingAcc.Balance < withdrawAmount {
			return nil,err
	} else {
			tradingAcc.Balance -= withdrawAmount}
	if err = Config.DB.Table("trading_account").Where("user_id = ?", Userid).UpdateColumn("balance", tradingAcc.Balance).Error; err != nil{
		return nil, err
	}
	pay := Payments{UserId: Userid,Amount: withdrawAmount,CurrentBalance: tradingAcc.Balance}
	if err = Config.DB.Create(pay).Error; err != nil {
		return nil, err
	}
	withdrawResponse.Userid = pay.UserId
	withdrawResponse.Amount = pay.Amount
	withdrawResponse.Type = "Withdraw"
	withdrawResponse.CurrentBalance = pay.CurrentBalance
	withdrawResponse.Message = "Process Successful"
	return &withdrawResponse,err
}

func Callback ()(callbackRes *CallbackResponse, err error){
	var payments Payments
	var callbackResponse CallbackResponse
	if err = Config.DB.Table("payments").Where("36522razorpay_link_id=? ,razorpay_link=?",payments.RazorpayLinkId,payments.RazorpayLink).First(&payments).Error; err != nil{
		return nil, err
	}
	callbackResponse.RazorpayPaymentID = payments.RazorpayLinkId
	callbackResponse.RazorpayPaymentLinkID = payments.RazorpayLink
	callbackResponse.RazorpayPaymentLinkStatus = "success"
	return &callbackResponse,nil
}