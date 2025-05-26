/**
 * @author ustcYang
 * @date 2025/4/13 18:52
 * @description
 */
package kafkamq

import (
	"context"
	"github.com/IBM/sarama"
	"trpc.group/trpc-go/trpc-database/kafka"
	"trpc.group/trpc-go/trpc-go/log"
)

type KafkaConsumer struct {
}

func (KafkaConsumer) Handle(ctx context.Context, msg *sarama.ConsumerMessage) error {
	if rawContext, ok := kafka.GetRawSaramaContext(ctx); ok {
		log.Infof("InitialOffset:%d", rawContext.Claim.InitialOffset())
	}
	log.Infof("get kafka message: %+v", msg)
	// Successful consumption is confirmed only when returning nil.
	return nil
}
