package common

import (
	"log"
	"regexp"
)

const MobileNumberPattern string = `^0[689]\d{8}$`

func MobileNumberValidate(mobileNumber string) bool {
	res, err := regexp.MatchString(MobileNumberPattern, mobileNumber)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}
