package main

import (
	"context"
	kafkamq "github.com/ustcLyric/helloworld/mq/kafka"

	pb "github.com/ustcLyric/trpcprotocol/helloworld"
)

type helloWorldImpl struct {
	pb.UnimplementedHelloWorld
}

func (s *helloWorldImpl) HelloEveryone(
	ctx context.Context,
	req *pb.HelloRequest,
) (*pb.HelloResponse, error) {
	kafkamq.NewMQHandle()
	return nil, nil
}
