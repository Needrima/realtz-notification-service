package helpers

import (
	"strconv"
)

// GetSkipAndLimit takes the amount of products to be queried and the page number
// and return the documents to skip in database, the limit of the next documents to get and the current page number.
// If amount is an empty string, or a string less than 1 or a non numeric value string, it defaults to 10 and if pageNo is an empty string,
// a string less than 1 or a non numeric value string it defaults to 1 .
func GetSkipAndLimit(amount, pageNo string) (int, int, int) {

	amountInt, err := strconv.Atoi(amount)
	if err != nil || amountInt < 1 {
		amountInt = 10
	}

	pageNoInt, err := strconv.Atoi(pageNo)
	if err != nil || pageNoInt < 1 {
		pageNoInt = 1
	}

	skip := amountInt * (pageNoInt - 1)
	limit := amountInt

	return skip, limit, pageNoInt
}
