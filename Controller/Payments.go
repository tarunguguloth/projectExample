package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectExample/Model"
)

func AddAmount(c *gin.Context){
	var addReq Model.AddRequest
	id := c.Params.ByName("Userid")
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
	var withdrawReq Model.WithdrawRequest
	id := c.Params.ByName("Userid")
	c.BindJSON(&withdrawReq)
	withdrawRes, err := Model.WithdrawAmount(withdrawReq,id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, withdrawRes)
	}
}

func Callback(c *gin.Context){
	callBackResponse, err := Model.Callback()
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, callBackResponse)
	}
}

