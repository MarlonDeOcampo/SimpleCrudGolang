package main

import (
	Initialize "gosimplecrud/initializer"
	"gosimplecrud/models"
)

func init() {
	Initialize.LoadEnv()
	Initialize.InitializeDBConnection()
}

func main() {
	Initialize.DB.AutoMigrate(&models.Payment{})
}
