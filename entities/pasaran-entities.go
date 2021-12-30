package entities

type Model_pasaranHome struct {
	Idcomppasaran          int    `json:"idcomppasaran"`
	Nmpasarantogel         string `json:"nmpasarantogel"`
	Tipepasaran            string `json:"tipepasaran"`
	PasaranDiundi          string `json:"pasarandiundi"`
	Jamtutup               string `json:"jamtutup"`
	Jamjadwal              string `json:"jamjadwal"`
	Jamopen                string `json:"jamopen"`
	Displaypasaran         int    `json:"displaypasaran"`
	StatusPasaran          string `json:"statuspasaran"`
	StatusPasaranActive    string `json:"statuspasaranactive"`
	StatusPasarancss       string `json:"statuspasaran_css"`
	StatusPasaranActivecss string `json:"statuspasaranactive_css"`
}
