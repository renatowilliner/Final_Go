package main

import (
	"examen/handler"
	"examen/service"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine

	finanzasHandler    *handler.FinanzasHandler
)

func main() {

	router = gin.Default()

	dependencies()

	mappingRoutes()

	router.Run(":8080")
}

func mappingRoutes() {
	// authClient := &clients.AuthClient{}
	// authMiddleware := middlewares.NewAuthMiddleware(authClient)

	// router.Use(middlewares.CORSMiddleware())
	// router.Use(authMiddleware.ValidateToken)

	groupFinanzas := router.Group("/finanzas")
	// groupFinanzas.Use(authMiddleware.ValidateToken)
	groupFinanzas.GET("/", finanzasHandler.ObtenerTablaAmortizacion)
	groupFinanzas.POST("/interesCompuesto", finanzasHandler.CalcularInteresCompuesto)
	groupFinanzas.POST("/", finanzasHandler.CalcularPresupuestoMensual)



}

func dependencies() {

	
	finanzasService := service.NewFinanzasService()

	// Pasa esta instancia al handler
	finanzasHandler = handler.NewFinanzasHandler(finanzasService)


}
