package main

import (
	"gosimplecrud/controllers"
	Initialize "gosimplecrud/initializer"
	"gosimplecrud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	Initialize.LoadEnv()
	Initialize.InitializeDBConnection()
	Initialize.DB.AutoMigrate(&models.Payment{})
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
	r.Run()
}
