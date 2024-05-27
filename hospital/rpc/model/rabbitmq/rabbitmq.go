package rabbitmq

//封装mq的封装函数
import (
	"demo/config"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	key       string
	MqUrl     string
}

func InitMQ() *RabbitMQ {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.GlobalConfig.Rabbitmq.Username,
		config.GlobalConfig.Rabbitmq.Password,
		config.GlobalConfig.Rabbitmq.Host,
		config.GlobalConfig.Rabbitmq.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Println("连接创建失败")
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Println("通讯创建失败")
	}
	//指定交换机的名称
	//QueueName := "遇见GGB"
	err = channel.ExchangeDeclare(config.GlobalConfig.Rabbitmq.Exchange,
		"fanout",
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil)
	if err != nil {
		log.Println("交换机声明失败")
	}
	mq := &RabbitMQ{
		conn:    conn,
		channel: channel,
		//QueueName: QueueName,
		Exchange: config.GlobalConfig.Rabbitmq.Exchange,
		key:      "",
		MqUrl:    url,
	}
	return mq
}

// 订阅发送
func PublishPub(r *RabbitMQ, message string) {
	err := r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Println("消息发送失败")
	}
}

func RecieveSub(r *RabbitMQ) <-chan amqp.Delivery {
	//fmt.Println(r.QueueName, "3333333333333")
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	fmt.Println(q.Name)
	if err != nil {
		log.Println("队列声明失败", err)
	}
	fmt.Println(r.Exchange)
	err = r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
	)
	msgs, err := r.channel.Consume(q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Println("消息读取失败")
	}
	//在这里使用for循环输出会消耗消息导致后面接不到消息
	//for msg := range msgs {
	//	fmt.Println(string(msg.Body), "66666")
	//}
	return msgs
}
