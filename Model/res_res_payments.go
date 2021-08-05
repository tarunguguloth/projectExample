package Model

type PaymentsRequest struct {
	Userid string `json:"user_id"`
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}
type PaymentsResponse struct {
	Userid         string `json:"user_id"`
	Amount         int    `json:"amount"`
	Type           string `json:"type"`
	CurrentBalance int    `json:"current_balance"`
	Message        string `json:"message"`
}
