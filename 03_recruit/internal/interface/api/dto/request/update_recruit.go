package request

import (
	"github.com/google/uuid"
	"github.com/shige1114/03_recruit/internal/app/command"
)

type UpdateRequest struct {
	ID           string `json:"Id"`
	GoodPoint    string `json:"GoodPoint"`
	SelfPR       string `json:"SelfPR"`
	ConcernPoint string `json:"ConcernPoint"`
}

func (req *UpdateRequest) ToUpdateSellerCommand() (*command.UpdateCommand, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}
	return &command.UpdateCommand{
		ID:           id,
		SelfPR:       req.SelfPR,
		GoodPoint:    req.GoodPoint,
		ConcernPoint: req.ConcernPoint,
	}, nil
}
