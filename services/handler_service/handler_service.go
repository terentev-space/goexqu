package handler_service

type IHandlerService interface {
	NewMessageEvent(message []byte) bool
}
