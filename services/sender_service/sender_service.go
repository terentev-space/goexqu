package sender_service

type ISenderService interface {
	Send(uri string, data string) bool
}
