package entities

type Model_dashboardwinlose_parent struct {
	Dashboardwinlose_nmagen string      `json:"dashboardwinlose_nmagen"`
	Dashboardwinlose_detail interface{} `json:"dashboardwinlose_detail"`
}
type Model_dashboardwinlose_child struct {
	Dashboardwinlose_winlose int `json:"dashboardwinlose_winlose"`
}
type Model_dashboardagenpasaranwinlose_parent struct {
	Dashboardagenpasaran_nmpasaran string      `json:"dashboardagenpasaran_nmpasaran"`
	Dashboardagenpasaran_detail    interface{} `json:"dashboardagenpasaran_detail"`
}
type Model_dashboardagenpasaranwinlose_child struct {
	Dashboardagenpasaran_winlose int `json:"dashboardagenpasaran_winlose"`
}
type Controller_dashboardwinlose struct {
	Year string `json:"year" validate:"required"`
}
