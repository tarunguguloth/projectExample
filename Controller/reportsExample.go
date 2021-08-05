
package Controller

/*
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectExample/Model"
	"time"
)

func DailyPendingOrders(c *gin.Context){
	var reports []Model.PendingOrders
	id := c.Params.ByName("id")
	c.BindJSON(&reports)
	err := Model.DailyPendingOrders(&reports, id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{[{
			"User_id" :reports.Userid,
			"OrderId" : reports.OrderId,
			"Stock_Name" : reports.StockName,
			"Order_Type" : reports.OrderType,
			"Book_Type" : reports.BookType,
			"Limit_Price" : reports.LimitPrice,
			"Quantity" : reports.Quantity,
			"Order_Price" : reports.OrderPrice,
			"Status" : reports.Status,
		},]})
	}
}
func Portfolio(c *gin.Context){
	var reports []Model.Holdings
	id := c.Params.ByName("id")
	from,err := time.Parse("020106150405",c.Params.ByName("From"))
	if err!= nil {
		fmt.Println(err.Error())
	}
	to,err := time.Parse("020106150405",c.Params.ByName("To"))
	if err!= nil {
		fmt.Println(err.Error())
	}
	c.BindJSON(&reports)
	err = Model.Portfolio(&reports, id, from, to)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{[
			"Id" : reports.Id,
			"User_id" :reports.Userid,
			"OrderId" : reports.OrderId,
			"Stock_Name" : reports.StockName,
			"Quantity" : reports.Quantity,
			"Buy_Price" : reports.BuyPrice,
		]})
	}
}
func OrdersHistroy(c *gin.Context){
	var report1 []Model.OrderHistory
	var report2 []Model.Holdings
	id := c.Params.ByName("id")
	from,err := time.Parse("020106150405",c.Params.ByName("From"))
	if err!= nil {
		fmt.Println(err.Error())
	}
	to,err := time.Parse("020106150405",c.Params.ByName("To"))
	if err!= nil {
		fmt.Println(err.Error())
	}
	c.BindJSON(&report1)
	c.BindJSON(&report2)
	err = Model.OrdersHistroy(&report1, &report2, id, from, to)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{[
			"User_id": report1.Userid,
			"Order_id": report1.OrderId,
			"Stock_Name":report1.StockName,
			"Quantity" : report1.Quantity,
			"Buy_Price" : report1.BuyPrice,
			"Sell_Price" : report1.SellPrice,
		]})
		c.JSON(http.StatusOK, gin.H{[
			"User_id": report2.Userid,
			"Order_id": report2.OrderId,
			"Stock_Name":report2.StockName,
			"Quantity" : report2.Quantity,
			"Buy_Price" : report2.BuyPrice,
			"Sell_Price" : "not Selled yet",
		]})
	}
}
func ProfitLossHistroy(c *gin.Context){
	var reports []Model.OrderHistory
	id := c.Params.ByName("id")
	from,err := time.Parse("020106150405",c.Params.ByName("From"))
	if err!= nil {
		fmt.Println(err.Error())
	}
	to,err := time.Parse("020106150405",c.Params.ByName("To"))
	if err!= nil {
		fmt.Println(err.Error())
	}
	c.BindJSON(&reports)
	err = Model.ProfitLossHistroy(&reports, id, from, to)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{[,
			"Id": reports.Id,
			"User_id": reports.Userid,
			"Order_id": reports.OrderId,
			"Stock_Name":reports.StockName,
			"Quantity" : reports.Quantity,
			"Buy_Price" : reports.BuyPrice,
			"Sell_Price" : reports.SellPrice,
			"P/l"        : reports.Quantity*(reports.BuyPrice-reports.SellPrice),
		]})
	}
}

*/