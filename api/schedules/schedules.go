package schedules

import (
	"book-api/api/books"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetScheduleListResponse(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, schedules)
}

func GetMyScheduleResponse(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mySchedules)
}

func PostBookSchedule(c *gin.Context) {
	var postData SchedulePostRequest
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scheduleID := postData.ScheduleID

	// Reduce Quota
	selectedSchedule := schedules[scheduleID]
	selectedSchedule.Quota -= 1
	schedules[scheduleID] = selectedSchedule

	response := fmt.Sprintf("Timeslot booked at %s", selectedSchedule.Timeslot)

	// Adding books detail to my schedule
	bookList := []books.Book{}
	for _, book := range postData.Books {
		bookInstance := books.GetBookByEdition(book.EditionNumber)
		bookList = append(bookList, bookInstance)
	}

	newSchedule := MySchedule{
		ScheduleID: scheduleID,
		Timeslot:   schedules[scheduleID].Timeslot,
		Books:      bookList,
	}
	mySchedules[scheduleID] = newSchedule

	c.IndentedJSON(http.StatusOK, map[string]string{
		"message": response,
	})
}
