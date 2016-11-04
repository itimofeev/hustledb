package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

func PrintJson(i interface{}) {
	j, err := json.Marshal(i)
	CheckErr(err)
	fmt.Println("JSON: ", string(j))
}

func IsFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
