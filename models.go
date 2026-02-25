package crmclient

type SpecialField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type CustomerPayload struct {
	Email         string         `json:"email"`
	IDInProject   string         `json:"id_in_project"` // Personel id (personelin sistemdeki id'si)
	Name          string         `json:"name"`
	Surname       string         `json:"surname"`
	CompanyName   string         `json:"company_name"`
	CompanyPhone  string         `json:"company_phone"`
	CompanyID     string         `json:"company_id_in_project"` // Firma id (firmanin sistemdeki id'si)
	BirthDay      string         `json:"birthday"`
	Phone         string         `json:"phone"`
	CountryCode   string         `json:"country_code"`
	IsSubUser     *bool          `json:"is_sub_user"`
	ParentUserID  string         `json:"parent_user_id"`
	SpecialFields []SpecialField `json:"special_fields"`
	CreatedAt     string         `json:"created_at"`
	AgentID       uint           `json:"agent_id"`
	AgentCode     string         `json:"agent_code"`
	CampaignCode  string         `json:"campaign_code"`
}

type CustomerDeletePayload struct {
	IDInProject string `json:"id_in_project"` // Personel id (personelin sistemdeki id'si)
}

type TransactionPayload struct {
	CustomerID             string                 `json:"customer_id"`
	Amount                 float64                `json:"amount"`
	TransactionIdInProject string                 `json:"transaction_id_in_project,omitempty"`
	Note                   string                 `json:"note"`
	BankInfo               int8                   `json:"bank_info,omitempty"`
	PackageKey             string                 `json:"package_key,omitempty"`
	SpecialFields          map[string]interface{} `json:"special_fields,omitempty"`
	Currency               string                 `json:"currency,omitempty"`
	IsRefund               *bool                  `json:"is_refund,omitempty"`
	CreatedAt              string                 `json:"created_at,omitempty"`
}

type TicketPayload struct {
	CustomerIDInProject string `json:"customer_id_in_project"` // main user id, not subuser
	Title               string `json:"title"`
	Name                string `json:"name"`
	Phone               string `json:"phone"`
	Description         string `json:"description"`
	TicketKey           string `json:"-"`
}

type SenderNamePayload struct {
	SenderName       string `json:"sender_name"`
	CustomerID       string `json:"customer_id"`
	Status           *int8  `json:"status"` // 0: pending, 1: approved, 2: rejected
	ServiceSendCount *uint  `json:"service_send_count"`
}

type SmsCancelPayload struct {
	CustomerID string `json:"customer_id"` // required
	SenderName string `json:"sender_name"` // required
	Keyword    string `json:"keyword"`     // required
	Status     *int8  `json:"status"`      // 0: pending, 1: approved, 2: rejected
	IsMailsent *bool  `json:"is_mail_sent"`
	ExpireDate string `json:"expire_date"`
}

// Table name, key and val
type TableData map[string]map[string]any
