package command

import "github.com/google/uuid"

type UpdateCommand struct {
	ID           uuid.UUID
	SelfPR       string
	GoodPoint    string
	ConcernPoint string
}
