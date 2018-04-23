package model

//Invoice representa uma invoice do sistema
type Invoice struct {
	DueDate           string  `json:"due_date"`
	Value             float64 `json:"value"`
	Description       string  `json:"description"`
	CodeLine          string  `json:"code_line"`
	BarCode           string  `json:"barcode"`
	CompanyIdentifier string  `json:"company_identifier"`
}
