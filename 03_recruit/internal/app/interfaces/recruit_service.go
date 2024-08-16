package interfaces

import (
	"github.com/google/uuid"
	"github.com/shige1114/03_recruit/internal/app/command"
	"github.com/shige1114/03_recruit/internal/app/query"
)

type ResultService interface {
	Create(recruitCommand *command.CreateCommand) error
	FindByUserId(userId uuid.UUID) (*query.RecruitQueryResultList, error)
	Update(recruitCommand *command.UpdateCommand) (*query.RecruitQueryResult, error)
	Delete(id uuid.UUID) error
}
