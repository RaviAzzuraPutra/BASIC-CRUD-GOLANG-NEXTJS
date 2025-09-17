package routes

import (
	"backend/controller/travel_controller"

	"github.com/gin-gonic/gin"
)

func TravelRouter(app *gin.Engine) {

	route := app

	route.POST("/add-travel", travel_controller.AddTravel)
	route.GET("/", travel_controller.GetAllTravel)
	route.GET("/:id", travel_controller.GetTravelByID)
	route.DELETE("/:id", travel_controller.DeleteTravel)
	route.PUT("/update-travel/:id", travel_controller.UpdateTravel)
}
