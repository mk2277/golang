package routes

import (
	"bank-web-api/controllers"

	"github.com/gin-gonic/gin"
)

func AccRoute(router *gin.Engine) {
	//All routes related to users comes here
	router.POST("/newaccount", controllers.CreateAccount)
	router.GET("/getaccount/:userId", controllers.GetAccount())
	router.PUT("/getaccount/:userId", controllers.EditAccount())
	router.DELETE("/getaccount/:userId", controllers.DeleteAccount())
	router.GET("/getaccounts", controllers.GetAllAccounts())
	router.GET("accbalance/:userId", controllers.GetAccountBalance())
	router.PUT("/addmoney/:userId", controllers.AddMoney())
	router.PUT("/rmvmoney/:userId", controllers.WithdrawMoney())
	router.PUT("/transfer/:user1/:user2", controllers.TransferMoney())
}
