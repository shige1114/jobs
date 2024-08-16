package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shige1114/03_recruit/internal/app/command"
	"github.com/shige1114/03_recruit/internal/app/interfaces"
	"github.com/shige1114/03_recruit/internal/app/mapper"
	"github.com/shige1114/03_recruit/internal/app/query"
	"github.com/shige1114/03_recruit/internal/domain/repository"
	"github.com/shige1114/03_recruit/internal/domain/value"
)

type RecruitService struct {
	repository repository.RepositoryInterface
}

func NewRecruitService(
	repository repository.RepositoryInterface,
) interfaces.ResultService {
	return &RecruitService{repository: repository}
}

func (rec RecruitService) Create(recruitCommand *command.CreateCommand) error {
	recruit, err := value.New(
		recruitCommand.CompanyID,
		recruitCommand.UserID,
		recruitCommand.SelfPR,
		recruitCommand.GoodPoint,
		recruitCommand.ConcernPoint,
	)
	if err != nil {
		return err
	}
	if rec.repository == nil {
		return errors.New("repository is not initialized")
	}
	if err := rec.repository.Insert(recruit); err != nil {
		return err
	}
	return nil
}

func (rec RecruitService) FindByUserId(userId uuid.UUID) (*query.RecruitQueryResultList, error) {
	recruits, err := rec.repository.GetByUserId(userId)

	if err != nil {
		return nil, err
	}

	var queryRecruitList query.RecruitQueryResultList

	for _, recruit := range *recruits {
		queryRecruitList.Result = append(queryRecruitList.Result, mapper.NweRecruitResult(&recruit))
	}
	return &queryRecruitList, nil
}

func (rec RecruitService) Update(recruitCommand *command.UpdateCommand) (*query.RecruitQueryResult, error) {
	return nil, nil
}

func (rec RecruitService) Delete(id uuid.UUID) error {
	return nil
}
