package rabbitmq

import (
	"testing"
	"fmt"
)

func TestNewRabbitMQSimple(t *testing.T) {

	mqSimple := NewRabbitMQSimple("xcar")
	err:=mqSimple.PublishSimple("test")
	if err != nil{
		fmt.Println(err)
	}

}
