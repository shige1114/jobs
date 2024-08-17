package message

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"

// 	"github.com/IBM/sarama"
// )

// type ConsumerService struct {
// 	topic    string
// 	groupID  string
// 	stopFlag bool
// 	consumer sarama.ConsumerGroup
// }

// func NewConsumerService(brokers []string, topic string, groupID string) (*ConsumerService, error) {
// 	config := sarama.NewConfig()

// 	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ConsumerService{
// 		topic:    topic,
// 		groupID:  groupID,
// 		consumer: consumer,
// 	}, nil
// }
// func (cs *ConsumerService) Start() error {
// 	handler := ConsumerGroupHandler{topic: cs.topic}
// 	for {
// 		// コンシューマーグループがメッセージを消費する
// 		err := cs.consumer.Consume(context.Background(), []string{cs.topic}, &handler)
// 		if err != nil {
// 			return fmt.Errorf("error while consuming messages: %v", err)
// 		}
// 	}
// }

// func (cs *ConsumerService) consumeMessages(ctx context.Context) error {
// 	handler := ConsumerGroupHandler{topic: cs.topic}
// 	return cs.consumer.Consume(ctx, []string{cs.topic}, &handler)
// }

// type ConsumerGroupHandler struct {
// 	topic string
// }

// func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
// func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
// func (h *ConsumerGroupHandler) ConsumeClaim(sarama.ConsumerGroupSession, sarama.ConsumerGroupClaim) error {
// 	// Implement your message processing logic here
// 	var sagaMessage SagaMessage

// 	if err := json.Unmarshal(msg.Value, &sagaMessage); err != nil {
// 		return err
// 	}

// 	switch sagaMessage.Action {
// 	case "CheckUser":
// 		if sagaMessage.UserSuccess {

// 		}
// 		if sagaMessage.UserFail {

// 		}
// 	case "CheckCompany":
// 		if sagaMessage.CompanySuccess {

// 		}
// 		if sagaMessage.CompanyFail {
// 			return sarama.ErrNoError
// 		}
// 	}
// 	return nil
// }
