package rabbitmq

import (
	"Michael-Min/octopus/config"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

//rabbitMQ结构体
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	MqUrl string
}

//创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, MqUrl: config.MqUrl}
}


func (r *RabbitMQ) DoConnect()  {
	var err error
	r.Conn, err = amqp.Dial(r.MqUrl)
	r.failOnErr(err, "failed to Connect rabbitmq!")
	//获取channel
	r.Channel, err = r.Conn.Channel()
	if err!=nil{
		r.Conn.Close()
	}
	r.failOnErr(err, "failed to open a channel")

}

//断开channel 和 connection
func (r *RabbitMQ) Destroy() {
	_=r.Channel.Close()
	_=r.Conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}