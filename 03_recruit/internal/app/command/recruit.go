package command

import (
	"github.com/google/uuid"
)

type Recruit struct {
	CompanyID    uuid.UUID
	UserID       uuid.UUID
	SelfPR       string
	GoodPoint    string
	ConcernPoint string
}
