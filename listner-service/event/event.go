package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", //name of the exchange
		"topic",      //type
		true,         //is the exchange durable
		false,        //auto-deleted
		false,        //this exchange is going to use internally, we will use with 3 ms
		false,        //no-wait?
		nil,          // arguments
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"MQ1", //name
		false, //durable
		false, //delete when unused
		true,  //exclusive [do not share]
		false, //no wait
		nil,   //arguments
	)
}
