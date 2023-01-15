package schedules

import (
	"book-api/api/books"
)

type Schedule struct {
	Timeslot   string
	ScheduleID string
	Quota      int
}

type MySchedule struct {
	ScheduleID string
	Timeslot   string
	Books      []books.Book
}

// Binding for request
type BookItem struct {
	EditionNumber string `json:"edition_number" binding:"required"`
}

type SchedulePostRequest struct {
	ScheduleID string     `json:"schedule_id" binding:"required,bookable,quota_available,not_booked_yet"`
	Books      []BookItem `json:"books" binding:"required,unique,valid_books"`
}

// Simulating data representation from DB Record
var schedules = map[string]Schedule{
	"0001": {
		Timeslot:   "2023-12-31 09:00:00 to 2023-12-31 09:45:00",
		ScheduleID: "0001",
		Quota:      20,
	},
	"0002": {
		Timeslot:   "2023-12-31 09:45:00 to 2023-12-31 10:30:00",
		ScheduleID: "0002",
		Quota:      0,
	},
	"0003": {
		Timeslot:   "2023-12-31 10:30:00 to 2023-12-31 11:00:00",
		ScheduleID: "0003",
		Quota:      10,
	},
}

var mySchedules = map[string]MySchedule{}
