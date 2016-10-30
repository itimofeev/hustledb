package util

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// CheckErr check error is nil and if not panic with message
func CheckErr(err error, msg ...interface{}) {
	if err != nil {
		log.Panicln(err, msg)
	}
}

// CheckOk check ok
func CheckOk(ok bool, msg ...interface{}) {
	if !ok {
		log.Panicln(msg)
	}
}

func CheckMatchesRegexp(regexpStr string, str string) {
	re, err := regexp.Compile(regexpStr)
	CheckErr(err, "Unable to compile "+regexpStr)

	CheckOk(re.MatchString(str), fmt.Sprintf("String '%s' not matched by regExp '%s'", str, regexpStr))
}

func Atoi(s string) int {
	r, err := strconv.Atoi(s)
	CheckErr(err, "unable to parse string to int"+s)
	return r
}
