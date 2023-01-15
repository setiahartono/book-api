package router

import (
	"book-api/api/books"
	"book-api/api/schedules"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func initCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookable", schedules.BookableValidator())
		v.RegisterValidation("quota_available", schedules.QuotaAvailableValidator())
		v.RegisterValidation("not_booked_yet", schedules.NotBookedYetValidator())
		v.RegisterValidation("valid_books", schedules.ValidBookValidator())
	}
}

func InitRouter(release bool) *gin.Engine {
	router := gin.Default()

	if release {
		gin.SetMode(gin.ReleaseMode)
	}

	initCustomValidator()
	router.GET("/books", books.GetBookListResponse)
	router.GET("/schedules", schedules.GetScheduleListResponse)
	router.POST("/schedules", schedules.PostBookSchedule)
	router.GET("/myschedules", schedules.GetMyScheduleResponse)
	return router
}
