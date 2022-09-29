package models

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type LocationRequest struct {
	UserId  string
	History []*Location
}
