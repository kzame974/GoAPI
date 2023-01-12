package models

// Temperature est la structure qui contient les données de la temperature
type Temperature struct {
	Time     string  `json:"time"`
	Location string  `json:"location"`
	Value    float32 `json:"value"`
}
