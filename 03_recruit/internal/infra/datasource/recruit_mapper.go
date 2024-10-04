package datasource

import (
	"github.com/shige1114/03_recruit/internal/domain/value"
)

func ToDB(rec *value.Recruit) *Recruit {
	return &Recruit{
		UserID:       rec.UserID,
		CompanyID:    rec.CompanyID,
		ID:           rec.ID,
		Name:         rec.Name,
		SelfPR:       rec.SelfPR.ToDB(),
		GoodPoint:    rec.GoodPoint.ToDB(),
		ConcernPoint: rec.ConcernPoint.ToDB(),
		UpdatedAt:    rec.UpdatedAt,
		CreatedAt:    rec.CreatedAt,
	}
}

func FromDB(rec *Recruit) *value.Recruit {
	return &value.Recruit{
		UserID:       rec.UserID,
		CompanyID:    rec.CompanyID,
		ID:           rec.ID,
		Name:         rec.Name,
		SelfPR:       value.SelfPR{Sentence: rec.SelfPR},
		GoodPoint:    value.GoodPoint{Sentence: rec.GoodPoint},
		ConcernPoint: value.ConcernPoint{Sentence: rec.ConcernPoint},
		UpdatedAt:    rec.UpdatedAt,
		CreatedAt:    rec.CreatedAt,
	}
}
