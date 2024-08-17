package mapper

import (
	"github.com/shige1114/03_recruit/internal/app/common"
	"github.com/shige1114/03_recruit/internal/interface/api/dto/response"
)

func ToResponse(recruit *common.RecruitResult) *response.Response {
	return &response.Response{
		ID:           recruit.ID,
		UserID:       recruit.UserID,
		CompanyID:    recruit.CompanyID,
		Name:         recruit.Name,
		GoodPoint:    recruit.GoodPoint,
		ConcernPoint: recruit.ConcernPoint,
		SelfPR:       recruit.SelfPR,
		UpdatedAt:    recruit.UpdatedAt,
		CreatedAt:    recruit.CreatedAt,
	}
}
func ToListResponse(recruits []*common.RecruitResult) *response.ResponseList {
	var responseList []*response.Response

	for _, recruit := range recruits {
		responseList = append(responseList, ToResponse(recruit))
	}
	return &response.ResponseList{Recruits: responseList}
}
