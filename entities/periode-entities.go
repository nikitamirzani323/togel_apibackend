package entities

type Model_periodeDashboard struct {
	Idtrxkeluaran     int     `json:"pasaran_invoice"`
	Nomorperiode      string  `json:"pasaran_nomorperiode"`
	Tanggalperiode    string  `json:"pasaran_tanggal"`
	Keluarantogel     string  `json:"pasaran_keluaran"`
	Total_Member      float32 `json:"pasaran_totalmember"`
	Total_bet         float32 `json:"pasaran_totalbet"`
	Total_outstanding float32 `json:"pasaran_totaloutstanding"`
	Total_cancelbet   float32 `json:"pasaran_totalcancelbet"`
	Winlose           float32 `json:"pasaran_winlose"`
	Revisi            int     `json:"pasaran_revisi"`
	Msgrevisi         string  `json:"pasaran_msgrevisi"`
}
type Controller_periodecancelbet struct {
	Sdata           string `json:"sData" validate:"required"`
	Page            string `json:"page"`
	Permainan       string `json:"permainan"`
	Idinvoice       int    `json:"idinvoice" validate:"required"`
	Idinvoicedetail int    `json:"idinvoicedetail" validate:"required"`
}
