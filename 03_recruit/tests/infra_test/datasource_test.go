package datasource_test

import (
	"testing"

	"github.com/shige1114/03_recruit/internal/infra/datasource"
)

func TestOpen(t *testing.T) {
	db := datasource.Open()

	if db == nil {
		t.Errorf("db error")
	}
}
