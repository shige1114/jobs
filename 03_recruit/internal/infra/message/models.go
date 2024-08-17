package message

type SagaMessage struct {
	Action         string `json:"action"`
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	CompanyID      string `json:"company_id"`
	UserSuccess    bool   `json:"user_success"`
	UserFail       bool   `json:"user_fail"`
	CompanySuccess bool   `json:"company_success"`
	CompanyFail    bool   `json:"compnay_fail"`
}
