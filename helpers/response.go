package helpers

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
	Time    string      `json:"time"`
}
type ResponsePasaran struct {
	Status       int         `json:"status"`
	Message      string      `json:"message"`
	Record       interface{} `json:"record"`
	Pasaraonline interface{} `json:"pasaranonline"`
	Time         string      `json:"time"`
}
type ResponsePeriode struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Time        string      `json:"time"`
	Totalbet    int         `json:"totalbet"`
	Subtotal    int         `json:"subtotal"`
	Subtotalwin int         `json:"subtotalwin"`
}
type ResponseReportWinlose struct {
	Status                  int         `json:"status"`
	Message                 string      `json:"message"`
	Record                  interface{} `json:"record"`
	Time                    string      `json:"time"`
	Subtotalturnover        int         `json:"subtotalturnover"`
	Subtotalwinlose         int         `json:"Subtotalwinlose"`
	Subtotalwinlose_company int         `json:"Subtotalwinlosecompany"`
}
type ResponseAdminManagement struct {
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	Record        interface{} `json:"record"`
	Time          string      `json:"time"`
	Listruleadmin interface{} `json:"listruleadmin"`
	Listiplist    interface{} `json:"listip"`
}
type ErrorResponse struct {
	Field string
	Tag   string
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
