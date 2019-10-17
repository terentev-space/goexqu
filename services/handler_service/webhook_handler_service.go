package handler_service

import (
	"encoding/json"
	"goexqu/services/handler_service/handler_service_entities"
	"goexqu/services/sender_service"
)

type WebhookHandlerService struct {
	IHandlerService
	senderService sender_service.ISenderService
}

func (service *WebhookHandlerService) Initialize(senderService sender_service.ISenderService) {
	service.senderService = senderService
}

func (service WebhookHandlerService) NewMessageEvent(message []byte) bool {
	var messageEntity handler_service_entities.WebhookMessageEntity
	err := json.Unmarshal(message, &messageEntity)
	if err != nil {
		return false
	}

	var sentCount uint = 0
	for _, url := range messageEntity.URLs {
		isSuccess := service.senderService.Send(url, messageEntity.Data)
		if isSuccess {
			sentCount++
			print("\n\t+\t(SENT:'" + url + "')")
		} else {
			print("\n\t-\t(FAIL:'" + url + "')")
		}
	}
	return sentCount > 0
}
