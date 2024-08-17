package response

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID           uuid.UUID
	CompanyID    uuid.UUID
	UserID       uuid.UUID
	Name         string
	SelfPR       string
	GoodPoint    string
	ConcernPoint string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

type ResponseList struct {
	Recruits []*Response `json:"Recruits"`
}
