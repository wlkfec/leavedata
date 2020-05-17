package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"makeuse/config"
)

func Publish(body string) error {
	log.Printf("dialing %q", config.Config.MQuri)
	connection, err := amqp.Dial(config.Config.MQuri)
	if err != nil {
		return fmt.Errorf("Dial: %s", err)
	}
	defer connection.Close()
	log.Printf("got Connection, getting Channel")
	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}
	log.Printf("got Channel, declaring %q Exchange (%q)", config.Config.MQexchangeType, config.Config.MQexchangeName)
	if err := channel.ExchangeDeclare(
		config.Config.MQexchangeName, // name
		config.Config.MQexchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}
	log.Printf("declared Exchange, publishing %dB body (%q)", len(body), body)
	if _, err = channel.QueueDeclare(config.Config.MQqueueName, true, false, false, false, nil); err != nil {
		return fmt.Errorf("Queue Declare: %s", err)
	}
	if err = channel.QueueBind(config.Config.MQqueueName, config.Config.MQroutingKey, config.Config.MQexchangeName, false, nil); err != nil {
		return fmt.Errorf("Queue Bind: %s", err)
	}
	// send message
	if err = channel.Publish(
		config.Config.MQexchangeName, // publish to an exchange
		config.Config.MQroutingKey,   // routing to 0 or more queues
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}
	return nil
}
