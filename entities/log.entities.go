package entities

type Model_log struct {
	Log_id       int    `json:"log_id"`
	Log_datetime string `json:"log_datetime"`
	Log_username string `json:"log_username"`
	Log_page     string `json:"log_page"`
	Log_tipe     string `json:"log_tipe"`
	Log_note     string `json:"log_note"`
}
