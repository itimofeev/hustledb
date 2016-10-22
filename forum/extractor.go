package forum

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"log"
	"net/http"
)

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

func GetTextFromHtml(body []byte) string {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	wholePage := ""

	// Find the review items
	doc.Find(".postcolor").Each(func(i int, s *goquery.Selection) {
		wholePage += getTextFromSelection(s)
	})

	return wholePage
}

// Text gets the combined text contents of each element in the set of matched
// elements, including their descendants.
func getTextFromSelection(s *goquery.Selection) string {
	var buf bytes.Buffer

	for _, n := range s.Nodes {
		buf.WriteString(getNodeText(n))
		buf.WriteString("\n")
	}
	return buf.String()
}

// Get the specified node's text content.
func getNodeText(node *html.Node) string {
	if node.Type == html.TextNode {
		// Keep newlines and spaces, like jQuery
		return node.Data
	} else if node.Type == html.ElementNode && node.Data == "br" {
		return "\n"
	} else if node.FirstChild != nil {
		var buf bytes.Buffer
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			buf.WriteString(getNodeText(c))
		}
		return buf.String()
	}

	return ""
}
