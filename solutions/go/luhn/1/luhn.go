package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	cleanText := strings.ReplaceAll(id, " ", "")
	if len(cleanText) < 2 {
		return false
	}
    sum := 0
	double := false
	for i := len(cleanText) - 1; i >= 0; i-- {
		current, err := strconv.Atoi(string(cleanText[i]))
		if err != nil {
			return false
		}
		if double {
			current = 2 * current
			if current > 9 {
				current = current - 9
			}
		}
		sum += current
		double = !double
	}
	return sum%10 == 0
}
