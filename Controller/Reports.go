package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectExample/Model"
)
func DailyPendingOrders(c *gin.Context){
	var reports []Model.PendingOrders
	id := c.Params.ByName("Userid")
	c.BindJSON(&reports)
	penOrderRes, err := Model.DailyPendingOrders(id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, penOrderRes)
	}
}
func Portfolio(c *gin.Context){
	var reports []Model.Holdings
	var reportsParamRequest Model.ReportsParamRequest
	id := c.Params.ByName("Userid")
	if err := c.BindQuery(&reportsParamRequest);err != nil{
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&reports)
	PortfolioRes, err := Model.Portfolio(id, reportsParamRequest.From, reportsParamRequest.To)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, PortfolioRes)
	}
}
func OrdersHistory(c *gin.Context){
	var report1 []Model.OrderHistory
	var report2 []Model.Holdings
	var reportsParamRequest Model.ReportsParamRequest
	id := c.Params.ByName("Userid")
	if err := c.BindQuery(&reportsParamRequest);err != nil{
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&report1)
	c.BindJSON(&report2)
	ordHisRes ,err := Model.OrdersHistory(id,reportsParamRequest.From, reportsParamRequest.To)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, ordHisRes)
	}
}
func ProfitLossHistory(c *gin.Context){
	var reports []Model.OrderHistory
	var reportsParamRequest Model.ReportsParamRequest
	id := c.Params.ByName("Userid")
	if err := c.BindQuery(&reportsParamRequest);err != nil{
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&reports)
	proLosRes, err := Model.ProfitLossHistory(id, reportsParamRequest.From, reportsParamRequest.To)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, proLosRes)
	}
}
