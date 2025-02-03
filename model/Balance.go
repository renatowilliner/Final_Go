package model

type Balance struct {
	Ingresos []float64 `json:"ingresos"`
	Gastos   []float64 `json:"gastos"`
}
