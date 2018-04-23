package model

//Invoice representa uma invoice do sistema
type Invoice struct {
	ID                int     `json:"id"`
	CreatedAt         string  `json:"due_date"`
	DueDate           string  `json:"created_at"`
	Value             float64 `json:"value"`
	PaidAt            string  `json:"paid_at"`
	Description       string  `json:"description"`
	Company           Company `json:"company"`
	FilePage          int64   `json:"file_page"`
	HasPaymentReceipt bool    `json:"has_payment_receipt"`
	HasFile           bool    `json:"has_file"`
	HasPassword       bool    `json:"has_password"`
}

//Company representa uma empresa
type Company struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Email      string `json:"email"`
}

//InvoiceAPIResponse representa a respota da API
type InvoiceAPIResponse struct {
	Data    Invoice `json:"data"`
	Success bool    `json:"success"`
}
