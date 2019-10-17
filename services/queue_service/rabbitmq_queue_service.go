package queue_service

import (
	"fmt"
	"github.com/streadway/amqp"
	"goexqu/helpers"
	"goexqu/services/handler_service"
	"time"
)

type RabbitmqQueueService struct {
	QueueService
}

func (service RabbitmqQueueService) Process(handlerService handler_service.IHandlerService) {
	connection, err := amqp.Dial(service.Config.QueueService.Connection)
	helpers.CheckError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()

	channel, err := connection.Channel()
	helpers.CheckError(err, "Failed to open a channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		service.Config.QueueService.Name,
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.CheckError(err, "Failed to declare a queue")

	err = channel.Qos(
		1,
		0,
		false,
	)
	helpers.CheckError(err, "Failed to set QoS")

	messages, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	helpers.CheckError(err, "Failed to register a consumer")
	service.processMessageList(handlerService, messages)
}

func (service RabbitmqQueueService) processMessageList(handlerService handler_service.IHandlerService, messages <-chan amqp.Delivery) {
	forever := make(chan bool)
	go func() {
		for message := range messages {
			service.sendMessage(handlerService, message)
		}
	}()
	println("Service Started!")
	<-forever
}

func (RabbitmqQueueService) sendMessage(handlerService handler_service.IHandlerService, message amqp.Delivery) {
	startDate := time.Now().Format(time.RFC3339Nano)
	fmt.Printf("\n [%s] New message: \n<<[ %s ]>>\n", startDate, message.Body)
	handlerService.NewMessageEvent(message.Body)
	_ = message.Ack(false)
	endDate := time.Now().Format(time.RFC3339Nano)
	fmt.Printf("\n\t^\t(%s) [FINISHED]\n\n", endDate)
}
