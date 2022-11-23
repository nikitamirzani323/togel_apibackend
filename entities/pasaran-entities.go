package entities

type Model_pasaranDashboardHome struct {
	Idcomppasaran    int         `json:"idcomppasaran"`
	Nmpasarantogel   string      `json:"nmpasarantogel"`
	PasaranDiundi    string      `json:"pasarandiundi"`
	PasaranURL       string      `json:"pasaranurl"`
	Jamtutup         string      `json:"jamtutup"`
	Jamjadwal        string      `json:"jamjadwal"`
	Jamopen          string      `json:"jamopen"`
	StatusPasaran    string      `json:"statuspasaran"`
	StatusPasarancss string      `json:"statuspasaran_css"`
	Listperiode      interface{} `json:"listperiode"`
}
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
type Model_pasaranEdit struct {
	Idpasarantogel                    string  `json:"idpasarantogel"`
	Nmpasarantogel                    string  `json:"nmpasaran"`
	PasaranDiundi                     string  `json:"pasarandiundi"`
	PasaranURL                        string  `json:"pasaranurl"`
	Jamtutup                          string  `json:"jamtutup"`
	Jamjadwal                         string  `json:"jamjadwal"`
	Jamopen                           string  `json:"jamopen"`
	Limitline4d                       int     `json:"limitline_4d"`
	Limitline3d                       int     `json:"limitline_3d"`
	Limitline3dd                      int     `json:"limitline_3dd"`
	Limitline2d                       int     `json:"limitline_2d"`
	Limitline2dd                      int     `json:"limitline_2dd"`
	Limitline2dt                      int     `json:"limitline_2dt"`
	Bbfs                              int     `json:"bbfs"`
	Minbet_432d                       float32 `json:"minbet_432d"`
	Maxbet4d_432d                     float32 `json:"maxbet4d_432d"`
	Maxbet3d_432d                     float32 `json:"maxbet3d_432d"`
	Maxbet3dd_432d                    float32 `json:"maxbet3dd_432d"`
	Maxbet2d_432d                     float32 `json:"maxbet2d_432d"`
	Maxbet2dd_432d                    float32 `json:"maxbet2dd_432d"`
	Maxbet2dt_432d                    float32 `json:"maxbet2dt_432d"`
	Maxbet4d_fullbb_432d              float32 `json:"maxbet4d_fullbb_432d"`
	Maxbet3d_fullbb_432d              float32 `json:"maxbet3d_fullbb_432d"`
	Maxbet3dd_fullbb_432d             float32 `json:"maxbet3dd_fullbb_432d"`
	Maxbet2d_fullbb_432d              float32 `json:"maxbet2d_fullbb_432d"`
	Maxbet2dd_fullbb_432d             float32 `json:"maxbet2dd_fullbb_432d"`
	Maxbet2dt_fullbb_432d             float32 `json:"maxbet2dt_fullbb_432d"`
	Maxbuy4d_432d                     float32 `json:"maxbuy4d_432d"`
	Maxbuy3d_432d                     float32 `json:"maxbuy3d_432d"`
	Maxbuy3dd_432d                    float32 `json:"maxbuy3dd_432d"`
	Maxbuy2d_432d                     float32 `json:"maxbuy2d_432d"`
	Maxbuy2dd_432d                    float32 `json:"maxbuy2dd_432d"`
	Maxbuy2dt_432d                    float32 `json:"maxbuy2dt_432d"`
	Limitotal4d_432d                  float32 `json:"limitotal4d_432d"`
	Limitotal3d_432d                  float32 `json:"limitotal3d_432d"`
	Limitotal3dd_432d                 float32 `json:"limitotal3dd_432d"`
	Limitotal2d_432d                  float32 `json:"limitotal2d_432d"`
	Limitotal2dd_432d                 float32 `json:"limitotal2dd_432d"`
	Limitotal2dt_432d                 float32 `json:"limitotal2dt_432d"`
	Limitglobal4d_432d                float32 `json:"limitglobal4d_432d"`
	Limitglobal3d_432d                float32 `json:"limitglobal3d_432d"`
	Limitglobal3dd_432d               float32 `json:"limitglobal3dd_432d"`
	Limitglobal2d_432d                float32 `json:"limitglobal2d_432d"`
	Limitglobal2dd_432d               float32 `json:"limitglobal2dd_432d"`
	Limitglobal2dt_432d               float32 `json:"limitglobal2dt_432d"`
	Limitotal4d_fullbb_432d           float32 `json:"limitotal4d_fullbb_432d"`
	Limitotal3d_fullbb_432d           float32 `json:"limitotal3d_fullbb_432d"`
	Limitotal3dd_fullbb_432d          float32 `json:"limitotal3dd_fullbb_432d"`
	Limitotal2d_fullbb_432d           float32 `json:"limitotal2d_fullbb_432d"`
	Limitotal2dd_fullbb_432d          float32 `json:"limitotal2dd_fullbb_432d"`
	Limitotal2dt_fullbb_432d          float32 `json:"limitotal2dt_fullbb_432d"`
	Limitglobal4d_fullbb_432d         float32 `json:"limitglobal4d_fullbb_432d"`
	Limitglobal3d_fullbb_432d         float32 `json:"limitglobal3d_fullbb_432d"`
	Limitglobal3dd_fullbb_432d        float32 `json:"limitglobal3dd_fullbb_432d"`
	Limitglobal2d_fullbb_432d         float32 `json:"limitglobal2d_fullbb_432d"`
	Limitglobal2dd_fullbb_432d        float32 `json:"limitglobal2dd_fullbb_432d"`
	Limitglobal2dt_fullbb_432d        float32 `json:"limitglobal2dt_fullbb_432d"`
	Disc4d_432d                       float32 `json:"disc4d_432d"`
	Disc3d_432d                       float32 `json:"disc3d_432d"`
	Disc3dd_432d                      float32 `json:"disc3dd_432d"`
	Disc2d_432d                       float32 `json:"disc2d_432d"`
	Disc2dd_432d                      float32 `json:"disc2dd_432d"`
	Disc2dt_432d                      float32 `json:"disc2dt_432d"`
	Win4d_432d                        float32 `json:"win4d_432d"`
	Win3d_432d                        float32 `json:"win3d_432d"`
	Win3dd_432d                       float32 `json:"win3dd_432d"`
	Win2d_432d                        float32 `json:"win2d_432d"`
	Win2dd_432d                       float32 `json:"win2dd_432d"`
	Win2dt_432d                       float32 `json:"win2dt_432d"`
	Win4dnodisc_432d                  float32 `json:"win4dnodisc_432d"`
	Win3dnodisc_432d                  float32 `json:"win3dnodisc_432d"`
	Win3ddnodisc_432d                 float32 `json:"win3ddnodisc_432d"`
	Win2dnodisc_432d                  float32 `json:"win2dnodisc_432d"`
	Win2ddnodisc_432d                 float32 `json:"win2ddnodisc_432d"`
	Win2dtnodisc_432d                 float32 `json:"win2dtnodisc_432d"`
	Win4dbb_kena_432d                 float32 `json:"win4dbb_kena_432d"`
	Win3dbb_kena_432d                 float32 `json:"win3dbb_kena_432d"`
	Win3ddbb_kena_432d                float32 `json:"win3ddbb_kena_432d"`
	Win2dbb_kena_432d                 float32 `json:"win2dbb_kena_432d"`
	Win2ddbb_kena_432d                float32 `json:"win2ddbb_kena_432d"`
	Win2dtbb_kena_432d                float32 `json:"win2dtbb_kena_432d"`
	Win4dbb_432d                      float32 `json:"win4dbb_432d"`
	Win3dbb_432d                      float32 `json:"win3dbb_432d"`
	Win3ddbb_432d                     float32 `json:"win3ddbb_432d"`
	Win2dbb_432d                      float32 `json:"win2dbb_432d"`
	Win2ddbb_432d                     float32 `json:"win2ddbb_432d"`
	Win2dtbb_432d                     float32 `json:"win2dtbb_432d"`
	Minbet_cbebas                     float32 `json:"minbet_cbebas"`
	Maxbet_cbebas                     float32 `json:"maxbet_cbebas"`
	Maxbuy_cbebas                     float32 `json:"maxbuy_cbebas"`
	Win_cbebas                        float32 `json:"win_cbebas"`
	Disc_cbebas                       float32 `json:"disc_cbebas"`
	Limitglobal_cbebas                float32 `json:"limitglobal_cbebas"`
	Limittotal_cbebas                 float32 `json:"limittotal_cbebas"`
	Minbet_cmacau                     float32 `json:"minbet_cmacau"`
	Maxbet_cmacau                     float32 `json:"maxbet_cmacau"`
	Maxbuy_cmacau                     float32 `json:"maxbuy_cmacau"`
	Win2d_cmacau                      float32 `json:"win2d_cmacau"`
	Win3d_cmacau                      float32 `json:"win3d_cmacau"`
	Win4d_cmacau                      float32 `json:"win4d_cmacau"`
	Disc_cmacau                       float32 `json:"disc_cmacau"`
	Limitglobal_cmacau                float32 `json:"limitglobal_cmacau"`
	Limitotal_cmacau                  float32 `json:"limitotal_cmacau"`
	Minbet_cnaga                      float32 `json:"minbet_cnaga"`
	Maxbet_cnaga                      float32 `json:"maxbet_cnaga"`
	Maxbuy_cnaga                      float32 `json:"maxbuy_cnaga"`
	Win3_cnaga                        float32 `json:"win3_cnaga"`
	Win4_cnaga                        float32 `json:"win4_cnaga"`
	Disc_cnaga                        float32 `json:"disc_cnaga"`
	Limitglobal_cnaga                 float32 `json:"limitglobal_cnaga"`
	Limittotal_cnaga                  float32 `json:"limittotal_cnaga"`
	Minbet_cjitu                      float32 `json:"minbet_cjitu"`
	Maxbet_cjitu                      float32 `json:"maxbet_cjitu"`
	Maxbuy_cjitu                      float32 `json:"maxbuy_cjitu"`
	Winas_cjitu                       float32 `json:"winas_cjitu"`
	Winkop_cjitu                      float32 `json:"winkop_cjitu"`
	Winkepala_cjitu                   float32 `json:"winkepala_cjitu"`
	Winekor_cjitu                     float32 `json:"winekor_cjitu"`
	Desc_cjitu                        float32 `json:"desc_cjitu"`
	Limitglobal_cjitu                 float32 `json:"limitglobal_cjitu"`
	Limittotal_cjitu                  float32 `json:"limittotal_cjitu"`
	Minbet_5050umum                   float32 `json:"minbet_5050umum"`
	Maxbet_5050umum                   float32 `json:"maxbet_5050umum"`
	Maxbuy_5050umum                   float32 `json:"maxbuy_5050umum"`
	Keibesar_5050umum                 float32 `json:"keibesar_5050umum"`
	Keikecil_5050umum                 float32 `json:"keikecil_5050umum"`
	Keigenap_5050umum                 float32 `json:"keigenap_5050umum"`
	Keiganjil_5050umum                float32 `json:"keiganjil_5050umum"`
	Keitengah_5050umum                float32 `json:"keitengah_5050umum"`
	Keitepi_5050umum                  float32 `json:"keitepi_5050umum"`
	Discbesar_5050umum                float32 `json:"discbesar_5050umum"`
	Disckecil_5050umum                float32 `json:"disckecil_5050umum"`
	Discgenap_5050umum                float32 `json:"discgenap_5050umum"`
	Discganjil_5050umum               float32 `json:"discganjil_5050umum"`
	Disctengah_5050umum               float32 `json:"disctengah_5050umum"`
	Disctepi_5050umum                 float32 `json:"disctepi_5050umum"`
	Limitglobal_5050umum              float32 `json:"limitglobal_5050umum"`
	Limittotal_5050umum               float32 `json:"limittotal_5050umum"`
	Minbet_5050special                float32 `json:"minbet_5050special"`
	Maxbet_5050special                float32 `json:"maxbet_5050special"`
	Maxbuy_5050special                float32 `json:"maxbuy_5050special"`
	Keiasganjil_5050special           float32 `json:"keiasganjil_5050special"`
	Keiasgenap_5050special            float32 `json:"keiasgenap_5050special"`
	Keiasbesar_5050special            float32 `json:"keiasbesar_5050special"`
	Keiaskecil_5050special            float32 `json:"keiaskecil_5050special"`
	Keikopganjil_5050special          float32 `json:"keikopganjil_5050special"`
	Keikopgenap_5050special           float32 `json:"keikopgenap_5050special"`
	Keikopbesar_5050special           float32 `json:"keikopbesar_5050special"`
	Keikopkecil_5050special           float32 `json:"keikopkecil_5050special"`
	Keikepalaganjil_5050special       float32 `json:"keikepalaganjil_5050special"`
	Keikepalagenap_5050special        float32 `json:"keikepalagenap_5050special"`
	Keikepalabesar_5050special        float32 `json:"keikepalabesar_5050special"`
	Keikepalakecil_5050special        float32 `json:"keikepalakecil_5050special"`
	Keiekorganjil_5050special         float32 `json:"keiekorganjil_5050special"`
	Keiekorgenap_5050special          float32 `json:"keiekorgenap_5050special"`
	Keiekorbesar_5050special          float32 `json:"keiekorbesar_5050special"`
	Keiekorkecil_5050special          float32 `json:"keiekorkecil_5050special"`
	Discasganjil_5050special          float32 `json:"discasganjil_5050special"`
	Discasgenap_5050special           float32 `json:"discasgenap_5050special"`
	Discasbesar_5050special           float32 `json:"discasbesar_5050special"`
	Discaskecil_5050special           float32 `json:"discaskecil_5050special"`
	Disckopganjil_5050special         float32 `json:"disckopganjil_5050special"`
	Disckopgenap_5050special          float32 `json:"disckopgenap_5050special"`
	Disckopbesar_5050special          float32 `json:"disckopbesar_5050special"`
	Disckopkecil_5050special          float32 `json:"disckopkecil_5050special"`
	Disckepalaganjil_5050special      float32 `json:"disckepalaganjil_5050special"`
	Disckepalagenap_5050special       float32 `json:"disckepalagenap_5050special"`
	Disckepalabesar_5050special       float32 `json:"disckepalabesar_5050special"`
	Disckepalakecil_5050special       float32 `json:"disckepalakecil_5050special"`
	Discekorganjil_5050special        float32 `json:"discekorganjil_5050special"`
	Discekorgenap_5050special         float32 `json:"discekorgenap_5050special"`
	Discekorbesar_5050special         float32 `json:"discekorbesar_5050special"`
	Discekorkecil_5050special         float32 `json:"discekorkecil_5050special"`
	Limitglobal_5050special           float32 `json:"limitglobal_5050special"`
	Limittotal_5050special            float32 `json:"limittotal_5050special"`
	Minbet_5050kombinasi              float32 `json:"minbet_5050kombinasi"`
	Maxbet_5050kombinasi              float32 `json:"maxbet_5050kombinasi"`
	Maxbuy_5050kombinasi              float32 `json:"maxbuy_5050kombinasi"`
	Belakangkeimono_5050kombinasi     float32 `json:"belakangkeimono_5050kombinasi"`
	Belakangkeistereo_5050kombinasi   float32 `json:"belakangkeistereo_5050kombinasi"`
	Belakangkeikembang_5050kombinasi  float32 `json:"belakangkeikembang_5050kombinasi"`
	Belakangkeikempis_5050kombinasi   float32 `json:"belakangkeikempis_5050kombinasi"`
	Belakangkeikembar_5050kombinasi   float32 `json:"belakangkeikembar_5050kombinasi"`
	Tengahkeimono_5050kombinasi       float32 `json:"tengahkeimono_5050kombinasi"`
	Tengahkeistereo_5050kombinasi     float32 `json:"tengahkeistereo_5050kombinasi"`
	Tengahkeikembang_5050kombinasi    float32 `json:"tengahkeikembang_5050kombinasi"`
	Tengahkeikempis_5050kombinasi     float32 `json:"tengahkeikempis_5050kombinasi"`
	Tengahkeikembar_5050kombinasi     float32 `json:"tengahkeikembar_5050kombinasi"`
	Depankeimono_5050kombinasi        float32 `json:"depankeimono_5050kombinasi"`
	Depankeistereo_5050kombinasi      float32 `json:"depankeistereo_5050kombinasi"`
	Depankeikembang_5050kombinasi     float32 `json:"depankeikembang_5050kombinasi"`
	Depankeikempis_5050kombinasi      float32 `json:"depankeikempis_5050kombinasi"`
	Depankeikembar_5050kombinasi      float32 `json:"depankeikembar_5050kombinasi"`
	Belakangdiscmono_5050kombinasi    float32 `json:"belakangdiscmono_5050kombinasi"`
	Belakangdiscstereo_5050kombinasi  float32 `json:"belakangdiscstereo_5050kombinasi"`
	Belakangdisckembang_5050kombinasi float32 `json:"belakangdisckembang_5050kombinasi"`
	Belakangdisckempis_5050kombinasi  float32 `json:"belakangdisckempis_5050kombinasi"`
	Belakangdisckembar_5050kombinasi  float32 `json:"belakangdisckembar_5050kombinasi"`
	Tengahdiscmono_5050kombinasi      float32 `json:"tengahdiscmono_5050kombinasi"`
	Tengahdiscstereo_5050kombinasi    float32 `json:"tengahdiscstereo_5050kombinasi"`
	Tengahdisckembang_5050kombinasi   float32 `json:"tengahdisckembang_5050kombinasi"`
	Tengahdisckempis_5050kombinasi    float32 `json:"tengahdisckempis_5050kombinasi"`
	Tengahdisckembar_5050kombinasi    float32 `json:"tengahdisckembar_5050kombinasi"`
	Depandiscmono_5050kombinasi       float32 `json:"depandiscmono_5050kombinasi"`
	Depandiscstereo_5050kombinasi     float32 `json:"depandiscstereo_5050kombinasi"`
	Depandisckembang_5050kombinasi    float32 `json:"depandisckembang_5050kombinasi"`
	Depandisckempis_5050kombinasi     float32 `json:"depandisckempis_5050kombinasi"`
	Depandisckembar_5050kombinasi     float32 `json:"depandisckembar_5050kombinasi"`
	Limitglobal_5050kombinasi         float32 `json:"limitglobal_5050kombinasi"`
	Limittotal_5050kombinasi          float32 `json:"limittotal_5050kombinasi"`
	Minbet_kombinasi                  float32 `json:"minbet_kombinasi"`
	Maxbet_kombinasi                  float32 `json:"maxbet_kombinasi"`
	Maxbuy_kombinasi                  float32 `json:"maxbuy_kombinasi"`
	Win_kombinasi                     float32 `json:"win_kombinasi"`
	Disc_kombinasi                    float32 `json:"disc_kombinasi"`
	Limitglobal_kombinasi             float32 `json:"limitglobal_kombinasi"`
	Limittotal_kombinasi              float32 `json:"limittotal_kombinasi"`
	Minbet_dasar                      float32 `json:"minbet_dasar"`
	Maxbet_dasar                      float32 `json:"maxbet_dasar"`
	Maxbuy_dasar                      float32 `json:"maxbuy_dasar"`
	Keibesar_dasar                    float32 `json:"keibesar_dasar"`
	Keikecil_dasar                    float32 `json:"keikecil_dasar"`
	Keigenap_dasar                    float32 `json:"keigenap_dasar"`
	Keiganjil_dasar                   float32 `json:"keiganjil_dasar"`
	Discbesar_dasar                   float32 `json:"discbesar_dasar"`
	Disckecil_dasar                   float32 `json:"disckecil_dasar"`
	Discgenap_dasar                   float32 `json:"discgenap_dasar"`
	Discganjil_dasar                  float32 `json:"discganjil_dasar"`
	Limitglobal_dasar                 float32 `json:"limitglobal_dasar"`
	Limittotal_dasar                  float32 `json:"limittotal_dasar"`
	Minbet_shio                       float32 `json:"minbet_shio"`
	Maxbet_shio                       float32 `json:"maxbet_shio"`
	Maxbuy_shio                       float32 `json:"maxbuy_shio"`
	Win_shio                          float32 `json:"win_shio"`
	Disc_shio                         float32 `json:"disc_shio"`
	Shioyear_shio                     string  `json:"shioyear_shio"`
	Limitglobal_shio                  float32 `json:"limitglobal_shio"`
	Limittotal_shio                   float32 `json:"limittotal_shio"`
	Displaypasaran                    int     `json:"displaypasaran"`
	StatusPasaranActive               string  `json:"statuspasaranactive"`
	Create                            string  `json:"create"`
	Createdate                        string  `json:"createdate"`
	Update                            string  `json:"update"`
	Updatedate                        string  `json:"updatedate"`
}
