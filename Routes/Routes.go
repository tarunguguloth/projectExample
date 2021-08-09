package Routes

import (
	"github.com/gin-gonic/gin"
	"projectExample/Controller"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	payments := r.Group("/payments")
	{
		payments.POST(":Userid/addAmount", Controller.AddAmount)
		payments.POST(":Userid/withdrawAmount", Controller.WithdrawAmount)
        payments.GET(":payment_status", Controller.Callback)
	}

	reports := r.Group("/reports")
	{
		reports.GET("pending_orders/:Userid", Controller.DailyPendingOrders)
		reports.GET("holdings/:Userid", Controller.Portfolio)
		reports.GET("order_history/:Userid", Controller.OrdersHistory)
		reports.GET("profit_loss_history/:Userid?", Controller.ProfitLossHistory)
	}
	return r
}