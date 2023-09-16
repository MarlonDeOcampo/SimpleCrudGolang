package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marlon/golangsimplecrud/controllers"
	Initialize "github.com/marlon/golangsimplecrud/initializer"
)

func init() {
	Initialize.LoadEnv()
	Initialize.InitializeDBConnection()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"getAll":  "/getAllPayments",
			"getById": "/getPayment",
			"create":  "/addPayment",
			"update":  "/updatePayment",
			"delete":  "/deletePayment",
		})
	})
	r.GET("/allpayments", controllers.GetAllPayments)
	r.GET("/allpayments", controllers.GetPayment)
	r.GET("/allpayments", controllers.AddPayment)
	r.GET("/allpayments", controllers.UpdatePayment)
	r.GET("/allpayments", controllers.DeletePayment)
	r.Run()
}
