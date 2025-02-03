package model

type InteresCompuesto struct {
	CapitalInicial float64 `json:"capitalInicial"`
	TasaInteres    float64 `json:"tasaInteres"`
	Anios          int     `json:"anios"`
}