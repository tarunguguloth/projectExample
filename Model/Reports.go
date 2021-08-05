package Model

import (
	"fmt"
	_ "fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"projectExample/Config "
	"time"
)



func DailyPendingOrders(reports []PendingOrders, Userid string) (penOrderRes DailyPendingOrderResponse,err error) {
	if err = Config.DB.Where("Userid = ?", Userid).Find(&reports).Error; err != nil {
		return err
	}
	return penOrderRes,nil

}
func Portfolio(reports []Holdings, Userid string, from string,to string ) ( portfolioRes PortfolioResponse, err error){

	from,err = time.Parse("020106150405",from)
	if err!= nil {
		fmt.Println(err.Error())
	}
	to,err = time.Parse("020106150405", to)
	if err!= nil {
		fmt.Println(err.Error())
	}
	if err = Config.DB.Where("Userid = ? AND updated_at BETWEEN ? AND ? ", Userid ,from ,to).Find(&reports).Error; err != nil {
		return nil,err
	}
	return portfolioRes,nil
}
func OrdersHistory(report1 []OrderHistory,report2 []Holdings, Userid string,from string,to string) (ordHisRes OrderHistoryResponse,err error){

	from,err = time.Parse("020106150405",from)
	if err!= nil {
		fmt.Println(err.Error())
	}
	to,err = time.Parse("020106150405", to)
	if err!= nil {
		fmt.Println(err.Error())
	}
	if err = Config.DB.Where("Userid = ? AND updated_at BETWEEN ? AND ?", Userid,from,to).Find(&report1).Error; err != nil {
		return nil,err
	}
	if err = Config.DB.Where("Userid = ? AND updated_at BETWEEN ? AND ?", Userid,from,to).Find(&report2).Error; err != nil {
		return nil,err
	}
	return ordHisRes,nil
}
func ProfitLossHistory(reports []OrderHistory, Userid string, from string,to string) (proLosRes ProfitLossHistoryResponse,err error){

	from,err = time.Parse("020106150405",from)
	if err!= nil {
		fmt.Println(err.Error())
	}
	to,err = time.Parse("020106150405", to)
	if err!= nil {
		fmt.Println(err.Error())
	}
	if err = Config.DB.Where("Userid = ?  AND updated_at BETWEEN ? AND ?", Userid,from,to).Find(reports).Error; err != nil {
		return nil,err
	}
	return proLosRes,nil
}

