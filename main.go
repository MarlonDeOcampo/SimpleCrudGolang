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
	r.GET("/getAllPayments", controllers.GetAllPayments)
	r.GET("/getPayment", controllers.GetPayment)
	r.GET("/addPayment", controllers.AddPayment)
	r.GET("/updatePayment", controllers.UpdatePayment)
	r.GET("/deletePayment", controllers.DeletePayment)
	r.Run()
}
