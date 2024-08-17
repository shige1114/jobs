package mapper

import (
	"github.com/shige1114/03_recruit/internal/app/common"
	"github.com/shige1114/03_recruit/internal/domain/value"
)

func NweRecruitResult(recruit *value.Recruit) *common.RecruitResult {

	if recruit == nil {
		return nil
	}

	return &common.RecruitResult{
		ID:           recruit.ID,
		UserID:       recruit.UserID,
		Name:         recruit.Name,
		CompanyID:    recruit.CompanyID,
		GoodPoint:    *recruit.GoodPoint.Sentence,
		ConcernPoint: *recruit.ConcernPoint.Sentence,
		SelfPR:       *recruit.SelfPR.Sentence,
		UpdatedAt:    recruit.UpdatedAt,
		CreatedAt:    recruit.CreatedAt,
	}
}
