/**
 * @author ustcYang
 * @date 2025/4/13 18:30
 * @description
 */
package kafkamq

type MQHandle interface {
}

type MQHandleImpl struct {
}

func NewMQHandle() MQHandle {
	return &MQHandleImpl{}
}

// 生产者 消费者编写
