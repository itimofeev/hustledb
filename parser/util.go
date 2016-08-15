package parser

import "log"

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
