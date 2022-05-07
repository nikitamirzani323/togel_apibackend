package entities

type Model_dashboardwinlose_parent struct {
	Dashboardwinlose_nmagen string      `json:"dashboardwinlose_nmagen"`
	Dashboardwinlose_detail interface{} `json:"dashboardwinlose_detail"`
}
type Model_dashboardwinlose_child struct {
	Dashboardwinlose_winlose int `json:"dashboardwinlose_winlose"`
}
type Controller_dashboardwinlose struct {
	Year string `json:"year" validate:"required"`
}
