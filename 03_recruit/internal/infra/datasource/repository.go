package datasource

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/shige1114/03_recruit/internal/domain/repository"
	"github.com/shige1114/03_recruit/internal/domain/value"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.RepositoryInterface {
	return &Repository{db: db}
}

func (repo *Repository) Insert(recruit *value.Recruit) error {
	dbRecurit := ToDB(recruit)
	if repo.db == nil {
		return errors.New("repository is not initialized")
	}
	if err := repo.db.Create(dbRecurit).Error; err != nil {
		return err
	}
	return nil
}
func (repo *Repository) GetByUserId(id uuid.UUID) ([]*value.Recruit, error) {
	var dbRecruits []Recruit

	if err := repo.db.Where("user_id = ?", id).Order("created_at").Find(&dbRecruits).Error; err != nil {
		return nil, err
	}
	recruits := make([]*value.Recruit, len(dbRecruits))

	for i, dbRecruit := range dbRecruits {
		recruits[i] = FromDB(&dbRecruit)
	}
	for _, item := range recruits {
		fmt.Printf("service recruits %+v\n", item) // 各ポインタの内容を出力
	}
	return recruits, nil
}
func (repo *Repository) Delete(id uuid.UUID) error {
	return repo.db.Delete(&Recruit{}, id).Error
}
func (repo *Repository) Put(recruit *value.Recruit) error {
	dbRecruit := ToDB(recruit)
	return repo.db.Model(&Recruit{}).Where("id = ?", dbRecruit.ID).Updates(dbRecruit).Error
}

func (repo *Repository) GetById(id uuid.UUID) (*value.Recruit, error) {
	var dbRecruit Recruit
	if err := repo.db.First(&dbRecruit, id).Error; err != nil {
		return nil, err
	}
	return FromDB(&dbRecruit), nil
}
