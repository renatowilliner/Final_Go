package service

import (
	"errors"
	"examen/model"
)

type FinanzasInterface interface {
	CalcularPresupuestoMensual(*model.Balance) (error,float64)
	CalcularInteresCompuesto(model.InteresCompuesto) (float64, error)
	ObtenerTablaAmortizacion(model.InteresCompuesto) ([]model.InteresCompuesto, error)
}
type FinanzasService struct{}



func NewFinanzasService() *FinanzasService {
	return &FinanzasService{}
}

func (f *FinanzasService)CalcularPresupuestoMensual(balance *model.Balance) (error,float64) {
	if len(balance.Ingresos) == 0 && len(balance.Gastos) == 0 {
		return errors.New("la lista de ingresos y egresos está vacía"), 0
	}

	totalIngresos := 0.0
	for _, ingreso := range balance.Ingresos {
		if ingreso < 100 && ingreso > 0 {
			totalIngresos += float64(ingreso)
		}
	}

	totalEgresos := 0.0
	for _, egreso := range balance.Gastos {
		if egreso < 100 && egreso > 0 {
			totalEgresos += float64(egreso)
		}
	}

	balanceNeto := totalIngresos - totalEgresos
	return nil, balanceNeto
}

func (f *FinanzasService)CalcularInteresCompuesto(interesCompuesto model.InteresCompuesto) (float64, error) {
	if interesCompuesto.CapitalInicial <= 0 || interesCompuesto.TasaInteres <= 0 || interesCompuesto.Anios <= 0 {
		return 0, errors.New("capital inicial, tasa de interés y años deben ser mayores a cero")
	}

	return interesCompuesto.CapitalInicial * (1 + interesCompuesto.TasaInteres) * float64(interesCompuesto.Anios), nil
}

func(f *FinanzasService) ObtenerTablaAmortizacion(interesesCompuestos model.InteresCompuesto) ([]model.InteresCompuesto, error) {
	if interesesCompuestos.CapitalInicial <= 0 || interesesCompuestos.TasaInteres <= 0 || interesesCompuestos.Anios <= 0 {
		return nil, errors.New("capital inicial, tasa de interés y años deben ser mayores a cero")
	}
	var tabla []model.InteresCompuesto
	saldo := interesesCompuestos.CapitalInicial
	for i := 1; i <= interesesCompuestos.Anios; i++ {
		interes := saldo * (interesesCompuestos.CapitalInicial / 100)
		saldoFinal := saldo + interes
		tabla = append(tabla, model.InteresCompuesto{
			Anios:          i,
			TasaInteres:    interes,
			CapitalInicial: saldoFinal,
		})
		saldo = saldoFinal
	}

	return tabla, nil
}
