package crmclient

type SpecialField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type CustomerPayload struct {
	Email         string         `json:"email"`
	IDInProject   string         `json:"id_in_project"`
	Name          string         `json:"name"`
	CompanyName   string         `json:"company_name"`
	CompanyID     string         `json:"company_id_in_project"`
	BirthDay      string         `json:"birthday"`
	Phone         string         `json:"phone"`
	SpecialFields []SpecialField `json:"special_fields"`
}

type TransactionPayload struct {
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
	Note       string  `json:"note"`
	BankInfo   int8    `json:"bank_info"`
	PackageKey string  `json:"package_key"`
	Currency   string  `json:"currency"`
}

type TicketPayload struct {
	CustomerIDInProject string `json:"customer_id_in_project"`
	Title               string `json:"title"`
	Description         string `json:"description,omitempty"`
}

// Table name, key and val
type TableData map[string]map[string]any
