package crmclient

type SpecialField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type CustomerPayload struct {
	Email         string         `json:"email"`
	IDInProject   string         `json:"id_in_project"`
	Name          string         `json:"name"`
	Phone         string         `json:"phone"`
	SpecialFields []SpecialField `json:"special_fields"`
}

type TransactionPayload struct {
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
	Note       string  `json:"note"`
	BankInfo   int8    `json:"bank_info"`
}

type TicketPayload struct {
	ProjectKey          string `json:"project_key"`
	CustomerIDInProject string `json:"customer_id_in_project"`
	Title               string `json:"title"`
	Description         string `json:"description,omitempty"`
}
