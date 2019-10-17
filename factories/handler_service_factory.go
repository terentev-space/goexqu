package factories

import (
	"goexqu/services/handler_service"
)

func CreateWebhookHandlerService() handler_service.WebhookHandlerService {
	return handler_service.WebhookHandlerService{}
}
