package services

import (
	"goexqu/factories"
	"goexqu/helpers"
	"goexqu/services/config_service"
)

type ExpressService struct {
	Config *config_service.Config
}

func (service ExpressService) RunWebhook() {
	queueService, err := factories.CreateQueueService(service.Config.QueueService.Type)
	helpers.CheckError(err, "Queue service error")
	queueService.Initialize(service.Config)

	senderService, err := factories.CreateSenderService("post")
	helpers.CheckError(err, "Sender service error")

	handlerService := factories.CreateWebhookHandlerService()
	handlerService.Initialize(senderService)

	queueService.Process(handlerService)
}
