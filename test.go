package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	// Create a new Collector instance
	c := colly.NewCollector()

	// On every an element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
	})

	// Visit http://www.baidu.com
	c.Visit("https://www.baidu.com")
}
