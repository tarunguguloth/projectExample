package Routes

import (
	"github.com/gin-gonic/gin"
	"projectExample/Controller"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	pay := r.Group("/Payments")
	{
		pay.POST(":id/addAmount", Controller.AddAmount)
		pay.POST(":id/withdrawAmount", Controller.WithdrawAmount)
		//will use callback according to time constraint
		//pay.GET("payments/:id/update", Controller.Callback)
	}
	return r

	rep := r.Group("/Reports")
	{
		rep.GET("pending_orders/:id", Controller.DailyPendingOrders)
		rep.GET("holdings/:id?{:From},{:To}", Controller.Portfolio)
		rep.GET("order_histroy/:id?{:From},{:To}", Controller.OrdersHistory)
		rep.GET("order_histroy/:id/profit_loss_history?{:From},{:To}", Controller.ProfitLossHistory)
	}
	return r
}