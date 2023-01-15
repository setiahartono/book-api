package schedules

import (
	"book-api/api/books"

	"github.com/go-playground/validator/v10"
)

var bookable validator.Func = func(fl validator.FieldLevel) bool {
	schedule_id, ok := fl.Field().Interface().(string)
	if ok {
		_, scheduleExists := schedules[schedule_id]
		if !scheduleExists {
			return false
		}
	}
	return true
}

var quotaAvailable validator.Func = func(fl validator.FieldLevel) bool {
	schedule_id, ok := fl.Field().Interface().(string)
	if ok {
		val := schedules[schedule_id]
		if val.Quota <= 0 {
			return false
		}
	}
	return true
}

var notBooked validator.Func = func(fl validator.FieldLevel) bool {
	schedule_id, ok := fl.Field().Interface().(string)
	if ok {
		_, exists := mySchedules[schedule_id]
		if exists {
			return false
		}
	}
	return true
}

var validBooks validator.Func = func(fl validator.FieldLevel) bool {
	editionNumbers, ok := fl.Field().Interface().([]BookItem)
	if ok {
		validBookList := books.GetBookList()
		for _, edition := range editionNumbers {
			_, validBook := validBookList[edition.EditionNumber]
			if !validBook {
				return false
			}
		}
	}
	return true
}

func BookableValidator() validator.Func {
	return bookable
}

func NotBookedYetValidator() validator.Func {
	return notBooked
}

func QuotaAvailableValidator() validator.Func {
	return quotaAvailable
}

func ValidBookValidator() validator.Func {
	return validBooks
}
