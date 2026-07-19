package isbnverifier

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	valid := strings.ReplaceAll(isbn, "-", "")

	if len(valid) != 10 {
		return false
	}

	sum := 0

	for i := 0; i < 10; i++ {
		char := string(valid[i])
		var current int

		if (char == "X" || char == "x") && i == 9 {
			current = 10
		} else {
			var err error
			current, err = strconv.Atoi(char)
			if err != nil {
				return false
			}
		}
		idx := 10 - i
		sum += current * idx
	}
	return sum%11 == 0
}
