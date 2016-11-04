package forum

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"log"
	"strings"
)

func GetMainContentFromForumHtml(body []byte) string {
	return getTextOfNodesBySelector(body, ".postcolor")
}

func GetMainTitleFromForumHtml(body []byte) string {
	mainTitle := getTextOfNodesBySelector(body, ".maintitle")
	mainTitle = strings.Replace(mainTitle, "Скрыть опции темы", "", 1)
	return strings.TrimSpace(mainTitle)
}

func getTextOfNodesBySelector(body []byte, selector string) string {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	wholePage := ""

	// Find the review items
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
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
