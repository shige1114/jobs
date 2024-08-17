package command

import (
	"github.com/google/uuid"
)

type CreateCommand struct {
	CompanyID    uuid.UUID
	Name         string
	UserID       uuid.UUID
	SelfPR       string
	GoodPoint    string
	ConcernPoint string
}
