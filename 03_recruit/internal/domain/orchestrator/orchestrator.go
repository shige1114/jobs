package orchestrator

import "github.com/google/uuid"

type OrchestratorInterface interface {
	CheckUserID(userId uuid.UUID) (bool, error)
	CheckCompanyID(companyID uuid.UUID) (string, error)
}
