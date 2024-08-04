package query

import "github.com/shige1114/03_recruit/internal/app/common"

type RecruitQueryResult struct {
	Result *common.RecruitResult
}

type RecruitQueryResultList struct {
	Result []*common.RecruitResult
}
