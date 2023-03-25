package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarId string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price string `json:"price"`
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newCar.CarId = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)
	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

func UpdateCar(ctx *gin.Context) {
	carId := ctx.Param("carId")
	condition := false
	var updateCar Car
	if err := ctx.ShouldBindJSON(&updateCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas {
		if carId == car.CarId {
			condition = true
			CarDatas[i] = updateCar
			CarDatas[i].CarId = carId
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", carId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully updated", carId),
	})
}

func GetCar(ctx *gin.Context) {
	carId := ctx.Param("carId")
	condition := false
	var carData Car
	for i, car := range CarDatas {
		if carId == car.CarId {
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", carId),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

func DeleteCar(ctx *gin.Context) {
	carId := ctx.Param("carId")
	condition := false
	var carIndex int

	for i, car := range CarDatas {
		if carId == car.CarId {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", carId),
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", carId),
	})
}

func GetAllCar(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"car": CarDatas,
	})
}
