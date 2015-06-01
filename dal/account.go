package expensly

import (
	"time"
)

type Account struct {
	Name           string    `json:"name"`
	StartBalance   int       `json:"start_balance"`
	DatetimeOpened time.Time `json:"datetime_opened"`
}

type AccountType struct {
	AbstractSlug
}
