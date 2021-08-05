package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectExample/Model"
)

func AddAmount(c *gin.Context){
	var addReq Model.PaymentsRequest
	id := c.Params.ByName("id")
	c.BindJSON(&addReq)
	addRes, err := Model.AddAmount(addReq,id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, addRes)
	}
}
func WithdrawAmount(c *gin.Context){
	var withdrawReq Model.PaymentsResponse
	id := c.Params.ByName("id")
	c.BindJSON(&withdrawReq)
	withdrawRes, err := Model.WithdrawAmount(withdrawReq,id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, withdrawRes)
	}
}

