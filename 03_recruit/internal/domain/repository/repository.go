package repository

import (
	"github.com/google/uuid"
	"github.com/shige1114/03_recruit/internal/domain/value"
)

type RepositoryInterface interface {
	Insert(recruit *value.Recruit) error
	GetByUserId(id uuid.UUID) ([]*value.Recruit, error)
	GetById(id uuid.UUID) (*value.Recruit, error)
	Delete(id uuid.UUID) error
	Put(recruit *value.Recruit) error
}
