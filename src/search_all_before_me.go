package main

import (
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	var identityNumber string

	_, err := fmt.Scan(&identityNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Порядковый номер: %s", searchByIdentity(identityNumber))
}

func searchByIdentity(identityNumber string) string {
	var d string

	xPath := "//*[contains(@class, 'RatingPage_table__position')]"
	url := "https://abit.itmo.ru/rating/master/budget/1905"
	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		r := bytes.NewReader(e.Response.Body)
		doc, err := htmlquery.Parse(r)
		if err != nil {
			log.Fatal(err)
		}

		nodes, err := htmlquery.QueryAll(doc, xPath)
		if err != nil {
			log.Fatal(err)
		}

		for _, node := range nodes {
			if spanData := node.LastChild.FirstChild.Data; spanData == identityNumber {
				d = node.FirstChild.Data
			}
		}
	})

	// Start scraping
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
