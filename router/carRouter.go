package router

import (
	controllers "simple-api/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/cars", controllers.GetAllCar)
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carId", controllers.UpdateCar)
	router.GET("/cars/:carId", controllers.GetCar)
	router.DELETE("/cars/:carId", controllers.DeleteCar)
	return router
}
