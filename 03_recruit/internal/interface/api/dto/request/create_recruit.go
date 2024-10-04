package request

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shige1114/03_recruit/internal/app/command"
)

type CreateRecruit struct {
	CompanyID    string `json:"CompanyID"`
	UserID       string `json:"UserID"`
	Name         string `json:"Name"`
	GoodPoint    string `json:"GoodPoint"`
	SelfPR       string `json:"SelfPR"`
	ConcernPoint string `json:"ConcernPoint"`
}

func ToCreateCommand(req *CreateRecruit) (*command.CreateCommand, error) {
	companyId, err := uuid.Parse(req.CompanyID)
	if err != nil {
		fmt.Printf("Error parsing CompanyID: %v\n", err)
		return nil, err
	}
	userId, err := uuid.Parse(req.UserID)
	if err != nil {
		fmt.Printf("Error parsing UserID: %v\n", err)
		return nil, err
	}

	return &command.CreateCommand{
		CompanyID:    companyId,
		UserID:       userId,
		Name:         req.Name,
		SelfPR:       req.SelfPR,
		GoodPoint:    req.GoodPoint,
		ConcernPoint: req.ConcernPoint,
	}, nil

}
