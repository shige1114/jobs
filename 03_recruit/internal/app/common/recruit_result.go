package common

import (
	"time"

	"github.com/google/uuid"
)

type RecruitResult struct {
	ID           uuid.UUID
	CompanyID    uuid.UUID
	UserID       uuid.UUID
	SelfPR       string
	GoodPoint    string
	ConcernPoint string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
