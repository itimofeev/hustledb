package util

import (
	"log"
	"strconv"
)

// CheckErr check error is nil and if not panic with message
func CheckErr(err error, msg string) {
	if err != nil {
		log.Panicln(msg, err)
	}
}

// CheckOk check ok
func CheckOk(ok bool, msg string) {
	if !ok {
		log.Panicln(msg)
	}
}

func Atoi(s string) int {
	r, err := strconv.Atoi(s)
	CheckErr(err, "unable to parse string to int"+s)
	return r
}
