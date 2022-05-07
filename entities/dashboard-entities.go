package entities

type Model_dashboardwinlose struct {
	Winlose int `json:"winlose"`
}
type Controller_dashboardwinlose struct {
	Year string `json:"year" validate:"required"`
}
