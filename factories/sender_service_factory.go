package factories

import (
	"errors"
	"goexqu/services/sender_service"
)

func CreateSenderService(name string) (sender_service.ISenderService, error) {
	switch name {
	case "post":
		return new(sender_service.HttpPostSenderService), nil
	default:
		return nil, errors.New("Unknown sender service: " + name)
	}
}
