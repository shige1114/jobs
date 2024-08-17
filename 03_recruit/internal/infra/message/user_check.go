package message

// import (
// 	"context"

// 	"github.com/IBM/sarama"
// 	"github.com/google/uuid"
// )

// type UserCheckRequest struct {
// 	UserID   uuid.UUID
// 	Producer sarama.SyncProducer
// 	Topic    string
// }

// type UserCheckResponse struct {
// 	UserID   uuid.UUID
// 	Consumer sarama.Consumer
// 	Topic    string
// }

// func (s *UserCheckRequest) Send(ctx context.Context) error {

// 	msg := &sarama.ProducerMessage{
// 		Topic: s.Topic,
// 		Key:   sarama.StringEncoder(s.UserID.String()),
// 		Value: sarama.StringEncoder("UserCheck"),
// 	}

// 	if _, _, err := s.Producer.SendMessage(msg); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *UserCheckResponse) Recieve(ctx context.Context) error {

// }
