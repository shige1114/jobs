package datasource_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/shige1114/03_recruit/internal/domain/value"
	"github.com/shige1114/03_recruit/internal/infra/datasource"
)

func TestOpen(t *testing.T) {
	db, err := datasource.Open()

	if err != nil {
		t.Errorf("db error")
	}
	log.Print(db)
}

func TestInsert(t *testing.T) {

	db, err := datasource.Open()
	if err != nil {
		t.Errorf("db error")
	}

	repo := datasource.NewRepository(db)
	if repo == nil {
		t.Errorf("repo error")
	}

	userID := uuid.New()
	companyID := uuid.New()
	recruit, err := value.New(userID, companyID, "test-good-point", "test-self-pr", "test-concern-point")
	if repo == nil {
		t.Errorf("repo %v", err)
	}

	if err := repo.Insert(recruit); err != nil {
		t.Errorf("repo %v", err)
	}

}

func TestGet(t *testing.T) {
	db, err := datasource.Open()
	if err != nil {
		t.Errorf("db error")
	}

	repo := datasource.NewRepository(db)
	if repo == nil {
		t.Errorf("repo error")
	}
	userId, err := uuid.Parse("a9f4a1e6-4c8d-4b7e-9b89-4c9076bb27c1")
	recruits, err := repo.GetByUserId(userId)

	fmt.Println(recruits)
}
