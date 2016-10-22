package main

import (
	"encoding/json"
	"fmt"
	"github.com/itimofeev/hustlesa/forum"
	"github.com/itimofeev/hustlesa/util"
	"io/ioutil"
)

const url = "http://hustle-sa.ru/forum/index.php?showtopic=3761"

func main() {
	//data := forum.GetUrlContent(url)

	data, err := ioutil.ReadFile("/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustlesa/tools/forum/3761.html")
	util.CheckErr(err, "")

	res := forum.GetTextFromHtml(data)

	results := forum.ParseForum([]byte(res))

	jsonData, err := json.Marshal(results)
	util.CheckErr(err, "")

	fmt.Println("!!!", string(jsonData)) //TODO remove
}

//err := ioutil.WriteFile("/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustlesa/tools/forum/3761.html", data, 0644)
