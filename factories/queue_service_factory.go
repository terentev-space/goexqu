package factories

import (
	"errors"
	"goexqu/services/queue_service"
)

func CreateQueueService(name string) (queue_service.IQueueService, error) {
	switch name {
	case "rabbitmq":
		return new(queue_service.RabbitmqQueueService), nil
	default:
		return nil, errors.New("Unknown queue service: " + name)
	}
}
