package value

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Recruit struct {
	ID           uuid.UUID
	CompanyID    uuid.UUID
	UserID       uuid.UUID
	SelfPR       SelfPR
	GoodPoint    GoodPoint
	ConcernPoint ConcernPoint
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

func New(id uuid.UUID, companyID uuid.UUID, userID uuid.UUID, selfPR string, goodPoint string, concernPoint string) *Recruit {
	pr := SelfPR{Sentence: &selfPR}
	gp := GoodPoint{Sentence: &goodPoint}
	cp := ConcernPoint{Sentence: &concernPoint}

	if pr.validate() && gp.validate() && cp.validate() {
		return &Recruit{
			ID:           id,
			CompanyID:    companyID,
			UserID:       userID,
			SelfPR:       pr,
			GoodPoint:    gp,
			ConcernPoint: cp,
			UpdatedAt:    time.Now(),
			CreatedAt:    time.Now(),
		}
	}
	return nil
}

func (rec *Recruit) ChangeSelfPR(text *string) error {
	selfPr := SelfPR{Sentence: text}
	if selfPr.validate() {
		rec.SelfPR = selfPr
		rec.UpdatedAt = time.Now()
		return nil
	}
	return errors.New("self pr: invalid sentence")
}

func (rec *Recruit) ChangeGoodPoint(text *string) error {
	goodPoint := GoodPoint{Sentence: text}
	if goodPoint.validate() {
		rec.GoodPoint = goodPoint
		rec.UpdatedAt = time.Now()
		return nil
	}
	return errors.New("good point: invalid sentence")
}
func (rec *Recruit) ChangeConcernPoint(text *string) error {
	concernPoint := ConcernPoint{Sentence: text}
	if concernPoint.validate() {
		rec.ConcernPoint = concernPoint
		rec.UpdatedAt = time.Now()
		return nil
	}
	return errors.New("concern point: invalid sentence")
}
