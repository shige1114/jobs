package response

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID           uuid.UUID `json:"Id"`
	CompanyID    uuid.UUID `json:"CompanyID"`
	UserID       uuid.UUID `json:"UserID"`
	Name         string    `json:"Name"`
	SelfPR       string    `json:"SelfPR"`
	GoodPoint    string    `json:"GoodPoint"`
	ConcernPoint string    `json:"ConcernPoint"`
	UpdatedAt    time.Time `json:"UpdatedAt"`
	CreatedAt    time.Time `json:"CreatedAt"`
}

type ResponseList struct {
	Recruits []*Response `json:"Recruits"`
}
