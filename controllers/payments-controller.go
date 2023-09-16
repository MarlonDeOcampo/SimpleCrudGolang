package controllers

import "github.com/gin-gonic/gin"

func GetAllPayments(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}

func GetPayment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}
func AddPayment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}
func UpdatePayment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}
func DeletePayment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}
