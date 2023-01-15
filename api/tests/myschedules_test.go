package testcases

import (
	"testing"
)

func TestMySchedule(t *testing.T) {
	makeRequest(t, "GET", "/myschedules", nil, 200)
}
