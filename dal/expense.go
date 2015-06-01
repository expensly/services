package expensly

import (
	"time"
)

type Expense struct {
	Payee Payee `json:"payee"`
	// Price in pence
	Amount   int       `json:"amount"`
	Category Category  `json:"category"`
	Date     time.Time `json:"date"`
	// Image path
	Image   string `json:"image"`
	Cleared bool   `json:"cleared"`
	Note    string `json:"note"`
}

type RepeatExpense struct {
	Expense
	Frequency Frequency `json:"frequency"`
}

type Frequency struct {
	Days   int `json:"days"`
	Months int `json:"months"`
	Years  int `json:"years"`
}
