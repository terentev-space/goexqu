package main

import (
	"goexqu/services"
	"goexqu/services/config_service"
)

func main() {
	config := config_service.LoadConfig()
	//config := config_service.DefaultConfig()
	expressService := services.ExpressService{Config: config}
	expressService.RunWebhook()
}
