package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	Initialize "github.com/marlon/golangsimplecrud/initializer"
	"github.com/marlon/golangsimplecrud/models"
)

// Get All Payment List
func GetAllPayments(c *gin.Context) {
	var payments []models.Payment
	Initialize.DB.Find(&payments)

	c.JSON(200, gin.H{
		"result": payments,
	})
}

// Get Specific Payment By ID
func GetPayment(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment
	Initialize.DB.First(&payment, id)

	c.JSON(200, gin.H{
		"message": payment,
	})
}

// Add Payment
func AddPayment(c *gin.Context) {
	var body struct {
		Lastname  string `json:"lastname"`
		Firstname string `json:"firstname"`
	}

	c.Bind(&body)

	payment := models.Payment{
		Paymentid: uuid.New(),
		Userid:    "01",
		Lastname:  body.Lastname,
		Firstname: body.Firstname,
	}

	result := Initialize.DB.Create(&payment)

	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
	}

	c.JSON(200, gin.H{
		"result": payment,
	})
}

// Update Payment
func UpdatePayment(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Lastname  string `json:"lastname"`
		Firstname string `json:"firstname"`
	}

	c.Bind(&body)

	var payment models.Payment
	Initialize.DB.First(&payment, id)

	Initialize.DB.Model(&payment).Updates(models.Payment{
		Lastname:  body.Lastname,
		Firstname: body.Firstname,
	})

	c.JSON(200, gin.H{
		"result": payment,
	})
}

// Delete Payment
func DeletePayment(c *gin.Context) {
	id := c.Param("id")

	Initialize.DB.Delete(&models.Payment{}, id)

	c.JSON(200, gin.H{
		"result": "Delete Successfull",
	})
}
