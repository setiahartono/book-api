package testcases

import (
	"testing"
)

func TestSchedulesList(t *testing.T) {
	makeRequest(t, "GET", "/schedules", nil, 200)
}

func TestCreateSchedule(t *testing.T) {
	data := []byte(`{"schedule_id": "0001", "books": [{"edition_number": "OL1709032M"},{"edition_number": "OL19609345M"}]}`)
	makeRequest(t, "POST", "/schedules", data, 200)
}

func TestInvalidScheduleRequests(t *testing.T) {
	// Booked schedule ID "0003" for initial data
	data := []byte(`{"schedule_id": "0003", "books": [{"edition_number": "OL1709032M"},{"edition_number": "OL19609345M"}]}`)
	makeRequest(t, "POST", "/schedules", data, 200)

	requests := [][]byte{
		//Invalid Schedule ID
		[]byte(`{"schedule_id": "999", "books": [{"edition_number": "OL1709032M"},{"edition_number": "OL19609345M"}]}`),

		//Already Booked Schedule ID
		[]byte(`{"schedule_id": "0003", "books": [{"edition_number": "OL1709032M"},{"edition_number": "OL19609345M"}]}`),

		//Zero quota schedule
		[]byte(`{"schedule_id": "0002", "books": [{"edition_number": "OL1709032M"},{"edition_number": "OL19609345M"}]}`),

		//Trying to book non-existent book
		[]byte(`{"schedule_id": "0002", "books": [{"edition_number": "OL1708382832M"},{"edition_number": "OL19609345M"}]}`),
	}

	for _, invalidRequest := range requests {
		makeRequest(t, "POST", "/schedules", invalidRequest, 400)
	}
}
