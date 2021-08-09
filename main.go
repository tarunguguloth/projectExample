package main

import (
	"fmt"
	"projectExample/Config"
	"projectExample/Model"
	"projectExample/Routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)


var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Model.Payments{}, &Model.PendingOrders{}, &Model.OrderHistory{}, &Model.Holdings{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}