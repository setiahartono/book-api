package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Title   string
	Authors []string
	Edition string
}

// List of books and necessary info to schedule lending
var books = map[string]Book{
	"OL4731539M": {
		Title:   "Gödel, Escher, Bach",
		Edition: "OL4731539M",
		Authors: []string{"Douglas R. Hofstadter"},
	},
	"OL22994526M": {
		Title:   "Just for fun",
		Edition: "OL22994526M",
		Authors: []string{"Linus Torvalds", "David Diamond"},
	},
	"OL2847892M": {
		Title:   "Digital computer fundamentals",
		Edition: "OL2847892M",
		Authors: []string{"Thomas C. Bartee"},
	},
	"OL3161962M": {
		Title:   "The Fifth Generation",
		Edition: "OL3161962M",
		Authors: []string{"Edward A. Feigenbaum", "Pamela McCorduck"},
	},
	"OL2043881M": {
		Title:   "Disaster recovery planning",
		Edition: "OL2043881M",
		Authors: []string{"Jon William Toigo", "Jon Toigo"},
	},
	"OL3184315M": {
		Title:   "The second self",
		Edition: "OL3184315M",
		Authors: []string{"Sherry Turkle"},
	},
	"OL19609345M": {
		Title:   "Digital design",
		Edition: "OL19609345M",
		Authors: []string{"M. Morris Mano"},
	},
	"OL1709032M": {
		Title:   "Reliable computer systems",
		Edition: "OL1709032M",
		Authors: []string{"Daniel P. Siewiorek"},
	},
	"OL3418615M": {
		Title:   "The principles of computer hardware",
		Edition: "OL3418615M",
		Authors: []string{"Clements, Alan"},
	},
	"OL2539425M": {
		Title:   "Mind Over Machine",
		Edition: "OL2539425M",
		Authors: []string{"Hubert L. Dreyfus"},
	},
}

func GetBookList() map[string]Book {
	return books
}

func GetBookByEdition(edition string) Book {
	bookInstance, exists := books[edition]
	if !exists {
		return Book{}
	}
	return bookInstance
}

func GetBookListResponse(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
