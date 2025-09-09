package crmclient

type SpecialField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type CustomerPayload struct {
	Email         string         `json:"email"`
	IDInProject   string         `json:"id_in_project"` // Personel id (personelin sistemdeki id'si)
	Name          string         `json:"name"`
	CompanyName   string         `json:"company_name"`
	CompanyID     string         `json:"company_id_in_project"` // Firma id (firmanin sistemdeki id'si)
	BirthDay      string         `json:"birthday"`
	Phone         string         `json:"phone"`
	IsSubUser     *bool          `json:"is_sub_user"`
	ParentUserID  string         `json:"parent_user_id"`
	SpecialFields []SpecialField `json:"special_fields"`
	CreatedAt     string         `json:"created_at"`
}

type CustomerDeletePayload struct {
	IDInProject  string `json:"id_in_project"` // Personel id (personelin sistemdeki id'si)
	IsSubUser    *bool  `json:"is_sub_user"`
	ParentUserID string `json:"parent_user_id"`
}

type TransactionPayload struct {
	CustomerID   string  `json:"customer_id"`
	IsSubUser    *bool   `json:"is_sub_user"`
	ParentUserID string  `json:"parent_user_id"`
	Amount       float64 `json:"amount"`
	Note         string  `json:"note"`
	BankInfo     int8    `json:"bank_info"`
	PackageKey   string  `json:"package_key"`
	Currency     string  `json:"currency"`
}

type TicketPayload struct {
	CustomerIDInProject string `json:"customer_id_in_project"` // main user id, not subuser
	Title               string `json:"title"`
	Description         string `json:"description,omitempty"`
}

// Table name, key and val
type TableData map[string]map[string]any
