package Rbmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

var url = "amqp://rbmq:rbmq@127.0.0.1:5672/"
var exchange = "project"
var queue = "pj_event"
var routing_key = "pj_event"
var content = map[string]interface{}{
	"name": "zelda",
}

// 生产者
func Pub_mq(uri, exchange, queue, routing_key string, content map[string]interface{}) error {
	// 建立连接
	connection, err := amqp.Dial(uri)
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err.Error())
		return err
	}
	defer connection.Close()
	// 创建一个Channel
	channel, err := connection.Channel()
	if err != nil {
		log.Println("Failed to open a channel:", err.Error())
		return err
	}
	defer channel.Close()

	// 声明exchange
	if err := channel.ExchangeDeclare(
		exchange, //name
		"direct", //exchangeType
		true,     //durable
		false,    //auto-deleted
		false,    //internal
		false,    //noWait
		nil,      //arguments
	); err != nil {
		log.Println("Failed to declare a exchange:", err.Error())
		return err
	}
	/*******死信队列参数********/
	// 声明一个queue
	args := amqp.Table{
		"x-message-ttl":             int64(60000),
		"x-dead-letter-exchange":    "dead",
		"x-dead-letter-routing-key": "",
	}
	/*******死信队列参数********/
	// 声明一个queue
	if _, err := channel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		args,  // arguments
	); err != nil {
		log.Println("Failed to declare a queue:", err.Error())
		return err
	}
	// exchange 绑定 queue
	channel.QueueBind(queue, routing_key, exchange, false, nil)

	// 发送
	messageBody, _ := json.Marshal(content)
	if err = channel.Publish(
		exchange,    // exchange
		routing_key, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            messageBody,
			//Expiration:      "60000", // 消息过期时间
		},
	); err != nil {
		log.Println("Failed to publish a message:", err.Error())
		return err
	}
	return nil
}
func Use_mq(uri, exchange, queue string) error {
	// 建立连接
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err.Error())
		return err
	}
	defer conn.Close()
	// 启动一个通道
	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel:", err.Error())
		return err
	}

	// 声明一个队列
	q, err := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println("Failed to declare a queue:", err.Error())
		return err
	}
	// 注册消费者
	msgs, err := ch.Consume(
		q.Name,    // queue
		"project", // 标签
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Println("Failed to register a consumer:", err.Error())
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Println(d.Type)
			log.Println(d.MessageId)
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
