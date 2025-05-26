package main

import (
	"context"
	"fmt"
	kafkamq "github.com/ustcLyric/helloworld/mq/kafka"
	pb "github.com/ustcLyric/trpcprotocol/helloworld"
	"time"
	kafka "trpc.group/trpc-go/trpc-database/kafka"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterHelloWorldService(s.Service("trpc.polarBear.helloworld.HelloWorld"), &helloWorldImpl{})
	kafka.RegisterKafkaConsumerService(s.Service("trpc.polarBear.helloworld.consumer"), &kafkamq.KafkaConsumer{})
	proxy := kafkamq.NewKafkaProduce()
	for i := 0; i < 3; i++ {
		err := proxy.SendMessage(context.Background(), []byte{'1'})
		//err := proxy.Produce(context.Background(), byte('1'))
		if err != nil {
			panic(err)
		}
		fmt.Println(i, []byte("1"), []byte("我爱你"))
		time.Sleep(time.Second)
	}
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
