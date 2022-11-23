package entities

type Model_periodeDashboard struct {
	Idtrxkeluaran     int    `json:"pasaran_invoice"`
	Nomorperiode      string `json:"pasaran_nomorperiode"`
	Tanggalperiode    string `json:"pasaran_tanggal"`
	Keluarantogel     string `json:"pasaran_keluaran"`
	Total_Member      int    `json:"pasaran_totalmember"`
	Total_bet         int    `json:"pasaran_totalbet"`
	Total_outstanding int    `json:"pasaran_totaloutstanding"`
}
type Controller_periodecancelbet struct {
	Sdata           string `json:"sData" validate:"required"`
	Page            string `json:"page"`
	Permainan       string `json:"permainan"`
	Idinvoice       int    `json:"idinvoice" validate:"required"`
	Idinvoicedetail int    `json:"idinvoicedetail" validate:"required"`
}
