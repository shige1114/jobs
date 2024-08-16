package datasource

import (
	"errors"

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
func (repo *Repository) GetByUserId(id uuid.UUID) (*[]value.Recruit, error) {
	var dbRecruits []Recruit

	if err := repo.db.Where("user_id = ?", id).Find(&dbRecruits).Error; err != nil {
		return nil, err
	}
	recruits := make([]value.Recruit, len(dbRecruits))
	for i, dbRecruit := range dbRecruits {
		recruits[i] = FromDB(&dbRecruit)
	}
	return &recruits, nil
}
func (repo *Repository) Delete(id uuid.UUID) error {

	return nil
}
func (repo *Repository) Put(recruit *value.Recruit) error {
	dbRecruit := ToDB(recruit)
	return repo.db.Model(&Recruit{}).Where("id = ?", dbRecruit.ID).Updates(dbRecruit).Error
}
