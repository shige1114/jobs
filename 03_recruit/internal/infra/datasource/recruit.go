package datasource

import (
	"time"

	"github.com/google/uuid"
)

type Recruit struct {
	ID           uuid.UUID `gorm:"primaryKey;not null;"`
	CompanyID    uuid.UUID `gorm:"not null;"`
	UserID       uuid.UUID `gorm:"not null;"`
	Name         string    `grom:"not null;"`
	SelfPR       string
	GoodPoint    string
	ConcernPoint string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
