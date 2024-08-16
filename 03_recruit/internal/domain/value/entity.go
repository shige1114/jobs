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

func New(companyID uuid.UUID, userID uuid.UUID, selfPR string, goodPoint string, concernPoint string) (*Recruit, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	pr, err := checkSelfPr(&selfPR)
	if err != nil {
		return nil, err
	}
	gp, err := checkGP(&goodPoint)
	if err != nil {
		return nil, err
	}
	cp, err := checkCP(&concernPoint)
	if err != nil {
		return nil, err
	}
	return &Recruit{
		ID:           id,
		CompanyID:    companyID,
		UserID:       userID,
		SelfPR:       *pr,
		GoodPoint:    *gp,
		ConcernPoint: *cp,
		UpdatedAt:    time.Now(),
		CreatedAt:    time.Now(),
	}, nil
}

func checkSelfPr(text *string) (*SelfPR, error) {
	selfPr := SelfPR{Sentence: text}
	if selfPr.validate() {
		return &selfPr, nil
	}
	return nil, errors.New("self pr: invalid sentence")
}
func checkGP(text *string) (*GoodPoint, error) {
	goodPoint := GoodPoint{Sentence: text}
	if goodPoint.validate() {
		return &goodPoint, nil
	}
	return nil, errors.New("good point: invalid sentence")
}
func checkCP(text *string) (*ConcernPoint, error) {
	concernPoint := ConcernPoint{Sentence: text}
	if concernPoint.validate() {
		return &concernPoint, nil
	}
	return nil, errors.New("concern point: invalid sentence")
}
func (rec *Recruit) ChangeSelfPR(text *string) error {
	selfPr, err := checkSelfPr(text)
	if err != nil {
		return err
	}
	rec.SelfPR = *selfPr
	rec.UpdatedAt = time.Now()
	return nil
}
func (rec *Recruit) ChangeGoodPoint(text *string) error {
	goodPoint, err := checkGP(text)
	if err != nil {
		return err
	}
	rec.GoodPoint = *goodPoint
	rec.UpdatedAt = time.Now()
	return nil
}
func (rec *Recruit) ChangeConcernPoint(text *string) error {
	ConcernPoint, err := checkCP(text)
	if err != nil {
		return err
	}
	rec.ConcernPoint = *ConcernPoint
	rec.UpdatedAt = time.Now()
	return nil
}
