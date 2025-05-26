/**
 * @author ustcYang
 * @date 2025/4/13 18:52
 * @description
 */
package kafkamq

import (
	"context"
	"fmt"
	"trpc.group/trpc-go/trpc-database/kafka"
)

// 声明一个全局变量
var proxy kafka.Client

func init() {
	// 默认初始化 Kafka Proxy
	proxy = kafka.NewClientProxy("trpc.polarBear.helloworld.produce")
}

type KafkaProducer interface {
	Produce(ctx context.Context, topic string, message []byte) error
	SendMessage(ctx context.Context, topic string, message []byte) error
}

type KafkaProducerImpl struct{}

// 实现 produce 方法
func (p *KafkaProducerImpl) Produce(ctx context.Context, message byte) error {
	// 这里添加生产消息的逻辑
	str := "Hello, World!"
	error := proxy.Produce(ctx, []byte(str), nil)
	if error != nil {
		fmt.Println("produce error:", error)
		return error
	}
	return nil
}

// 实现 produce 方法
func (p *KafkaProducerImpl) SendMessage(ctx context.Context, message []byte) error {
	// 这里添加生产消息的逻辑
	partition, offset, error := proxy.SendMessage(ctx, "helloworld", []byte{'1'}, message)
	if error != nil {
		fmt.Println("produce error:", error)
		return error
	}
	fmt.Printf("partition: %d, offset: %d\n", partition, offset)
	return nil
}

func NewKafkaProduce() *KafkaProducerImpl {
	return &KafkaProducerImpl{}
}
