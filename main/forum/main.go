package main

import (
	"fmt"
	"github.com/itimofeev/hustlesa/forum"
)

const url = "http://hustle-sa.ru/forum/index.php?showtopic=3761"

func main() {
	data := forum.GetUrlContent(url)

	res := forum.GetTextFromHtml(data)

	fmt.Println("!!!", res) //TODO remove
}

//err := ioutil.WriteFile("/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustlesa/tools/forum/3761.html", data, 0644)
