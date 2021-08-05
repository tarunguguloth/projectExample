package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectExample/Model"
)
func DailyPendingOrders(c *gin.Context){
	var reports []Model.PendingOrders
	id := c.Params.ByName("id")
	c.BindJSON(&reports)
	penOrderRes, err := Model.DailyPendingOrders(reports, id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, penOrderRes)
	}
}
func Portfolio(c *gin.Context){
	var reports []Model.Holdings
	id := c.Params.ByName("id")
	from := c.Params.ByName("From")
	to := c.Params.ByName("To")
	c.BindJSON(&reports)
	PortfolioRes, err := Model.Portfolio(reports, id, from, to)
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
	id := c.Params.ByName("id")
	from := c.Params.ByName("From")
	to := c.Params.ByName("To")
	c.BindJSON(&report1)
	c.BindJSON(&report2)
	ordHisRes ,err := Model.OrdersHistory(report1, report2, id, from, to)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, ordHisRes)
	}
}
func ProfitLossHistory(c *gin.Context){
	var reports []Model.OrderHistory
	id := c.Params.ByName("id")
	from := c.Params.ByName("From")
	to := c.Params.ByName("To")
	c.BindJSON(&reports)
	proLosRes, err := Model.ProfitLossHistory(reports, id, from, to)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, proLosRes)
	}
}
