package entities

type Controller_periodecancelbet struct {
	Sdata           string `json:"sData" validate:"required"`
	Page            string `json:"page"`
	Permainan       string `json:"permainan"`
	Idinvoice       int    `json:"idinvoice" validate:"required"`
	Idinvoicedetail int    `json:"idinvoicedetail" validate:"required"`
}
