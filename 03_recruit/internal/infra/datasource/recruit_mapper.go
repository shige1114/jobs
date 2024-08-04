package datasource

import (
	"github.com/shige1114/03_recruit/internal/domain/models"
	"github.com/shige1114/03_recruit/internal/domain/value"
)

func ToDB(rec *value.Recruit) *models.Recruit {
	return &models.Recruit{
		UserID:       rec.UserID,
		CompanyID:    rec.CompanyID,
		ID:           rec.ID,
		SelfPR:       rec.SelfPR.ToDB(),
		GoodPoint:    rec.GoodPoint.ToDB(),
		ConcernPoint: rec.ConcernPoint.ToDB(),
		UpdatedAt:    rec.UpdatedAt,
		CreatedAt:    rec.CreatedAt,
	}
}

func FromDB(rec *models.Recruit) *value.Recruit {
	return &value.Recruit{
		UserID:       rec.UserID,
		CompanyID:    rec.CompanyID,
		ID:           rec.ID,
		SelfPR:       value.SelfPR{Sentence: &rec.SelfPR},
		GoodPoint:    value.GoodPoint{Sentence: &rec.GoodPoint},
		ConcernPoint: value.ConcernPoint{Sentence: &rec.ConcernPoint},
		UpdatedAt:    rec.UpdatedAt,
		CreatedAt:    rec.CreatedAt,
	}
}
