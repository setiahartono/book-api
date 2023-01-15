package testcases

import (
	"testing"
)

func TestBookList(t *testing.T) {
	makeRequest(t, "GET", "/books", nil, 200)
}
