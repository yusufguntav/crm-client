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
	BirthDay      string         `json:"birthday"`
	Phone         string         `json:"phone"`
	SpecialFields []SpecialField `json:"special_fields"`
}

type TransactionPayload struct {
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
	Note       string  `json:"note"`
	BankInfo   int8    `json:"bank_info"`
	PackageID  *uint   `json:"package_id"`
}

type TicketPayload struct {
	ProjectKey          string `json:"project_key"`
	CustomerIDInProject string `json:"customer_id_in_project"`
	Title               string `json:"title"`
	Description         string `json:"description,omitempty"`
}

// Table name, key and val
type TableData map[string]map[string]any
