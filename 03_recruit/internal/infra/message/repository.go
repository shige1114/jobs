package message

import (
	"github.com/IBM/sarama"
)

// 実際のメッセージの送信作業

type Repository struct {
	Producer sarama.SyncProducer
}

func NewRepository(producer sarama.SyncProducer) *Repository {
	return &Repository{Producer: producer}
}

func (rec *Repository) AddSaga() {

}

func (rec *Repository) CheckUserID() {}

func (rec *Repository) CheckCompany() {}
