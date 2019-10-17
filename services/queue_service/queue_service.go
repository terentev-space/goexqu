package queue_service

import (
	"goexqu/services/config_service"
	"goexqu/services/handler_service"
)

type IQueueService interface {
	Initialize(config *config_service.Config)
	Process(handlerService handler_service.IHandlerService)
}

type QueueService struct {
	IQueueService
	Config *config_service.Config
}

func (service *QueueService) Initialize(config *config_service.Config) {
	service.Config = config
}
