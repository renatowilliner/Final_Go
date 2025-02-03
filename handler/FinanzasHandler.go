package handler

import (
	"examen/model"
	"examen/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FinanzasHandler struct {
	finanzasService service.FinanzasInterface
}

func NewFinanzasHandler(finanzasService service.FinanzasInterface) *FinanzasHandler {
	return &FinanzasHandler{finanzasService: finanzasService}
}


func (handler *FinanzasHandler) CalcularPresupuestoMensual(c *gin.Context) {
	var balance model.Balance

	if handler.finanzasService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Servicio no configurado"})
		return
	}
	
	
	err := c.ShouldBindJSON(&balance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		err,resultado := handler.finanzasService.CalcularPresupuestoMensual(&balance)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"resultado": resultado})
		}

	}

}

func (handler *FinanzasHandler) CalcularInteresCompuesto(c *gin.Context) {
	var interesCompuesto model.InteresCompuesto
	err := c.ShouldBindJSON(&interesCompuesto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		resultado,err := handler.finanzasService.CalcularInteresCompuesto(interesCompuesto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"respuesta": resultado})
		}

	}

}

func (handler *FinanzasHandler) ObtenerTablaAmortizacion(c *gin.Context) {
	capitalInicialStr := c.Query("capitalInicial")
	tasaInteresStr := c.Query("tasaInteres")
	aniosStr := c.Query("anios")

	capitalInicial, err := strconv.ParseFloat(capitalInicialStr, 64)
	tasaInteres, err := strconv.ParseFloat(tasaInteresStr, 64)
	anios, err := strconv.Atoi(aniosStr)
	

	var interesesCompuestos = model.InteresCompuesto{
		CapitalInicial: capitalInicial,
		TasaInteres:    tasaInteres,
		Anios:          anios,
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		resultado,err := handler.finanzasService.ObtenerTablaAmortizacion(interesesCompuestos)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"respuesta": resultado})
		}

	}

}








