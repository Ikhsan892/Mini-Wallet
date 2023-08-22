package models

type Transaction struct {
	BaseModel
	WalletId    string  `json:"wallet_id"`
	Type        string  `json:"type" gorm:"not null"`
	Amount      float64 `json:"amount"`
	IssuedBy    string  `json:"issued_by"`
	ReferenceId string  `json:"reference_id"`
	Status      string  `json:"status"`
}
