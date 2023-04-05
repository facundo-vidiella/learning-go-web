package product

import (
	"regexp"
	"strings"
	"time"
)

type ProductHelpers interface {
	validateDate() bool
}

func ValidateDate(date string) bool {
	re := regexp.MustCompile(`\d{2}/\d{2}/\d{4}`)
	currentTime := time.Now().Format("2006-01-02")
	splitedCurrentTime := strings.Split(currentTime, "-")
	splitedDate := strings.Split(date, "/")
	if splitedCurrentTime[0] > splitedDate[2] {
		return false
	}
	return re.MatchString(date)
}
