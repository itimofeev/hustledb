package util

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
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

func GetJson(i interface{}) string {
	j, err := json.Marshal(i)
	CheckErr(err)
	return string(j)
}

func IsFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetUrlContent(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	utf8, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(utf8)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func DownloadUrlToFile(url, path string) []byte {
	data := GetUrlContent(url)

	err := ioutil.WriteFile(path, data, 0644)

	CheckErr(err, "")
	return data
}
func DownloadUrlToFileIfNotExists(url, path string) []byte {
	if !IsFileExists(path) {
		return DownloadUrlToFile(url, path)
	}

	data, err := ioutil.ReadFile(path)
	CheckErr(err, "")

	return data
}

// 2014-05-18kshdfjkhsdf
func ParseForumDate(dateStr string) time.Time {
	year := dateStr[:4]
	monthStr := dateStr[5:7]
	day := dateStr[8:10]

	month := time.Month(Atoi(monthStr))

	return time.Date(Atoi(year), month, Atoi(day), 0, 0, 0, 0, time.UTC)
}
