package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
  doc, err := goquery.NewDocument("https://ja.wikipedia.org/wiki/NLP")
	if err != nil {
		fmt.Println(err)
	}
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}
