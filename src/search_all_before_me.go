package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	var identityNumber int
	var answer string

	fmt.Scan(&identityNumber)

	answer = searchByIdentity(identityNumber)

	if len(answer) == 0 {
		fmt.Print("Answer is empty")
	}
}

func searchByIdentity(identityNumber int) string {
	xPath := "//*[contains(@class, 'RatingPage_table__position')]"
	url := "abit.itmo.ru/rating/master/budget/1905"
	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Load the HTML content into an XPath queryable context
		doc, err := htmlquery.Parse(e.Response.Body)
		if err != nil {
			log.Fatal(err)
		}

		nodes, err := htmlquery.QueryAll(doc, xPath)
		if err != nil {
			log.Fatal(err)
		}

		for _, node := range nodes {
			fmt.Println(htmlquery.SelectAttr(node, "span"))
		}
	})

	// Start scraping
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
