package main

import (
	Initialize "github.com/marlon/golangsimplecrud/initializer"
	"github.com/marlon/golangsimplecrud/models"
)

func init() {
	Initialize.LoadEnv()
	Initialize.InitializeDBConnection()
}

func main() {
	Initialize.DB.AutoMigrate(&models.Payment{})
}
