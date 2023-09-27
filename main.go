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
			"getById": "/getPayment/{id}",
			"create":  "/addPayment",
			"update":  "/updatePayment/{id}",
			"delete":  "/deletePayment/{id}",
		})
	})
	r.GET("/getAllPayments", controllers.GetAllPayments)
	r.GET("/getPayment/:id", controllers.GetPayment)
	r.POST("/addPayment", controllers.AddPayment)
	r.PUT("/updatePayment/:id", controllers.UpdatePayment)
	r.DELETE("/deletePayment/:id", controllers.DeletePayment)
	r.Run(":4005")
}
