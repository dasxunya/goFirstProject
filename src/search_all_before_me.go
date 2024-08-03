package main

import (
	"fmt"
	"github.com/gocolly/colly"
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
	var answer string

	xPath := "//*[contains(@class, 'RatingPage_table__position')]"
	url := "abit.itmo.ru/rating/master/budget/1905"
	c := colly.NewCollector()

	c.Visit(url)

	c.OnHTML(xPath, func(collyElement *colly.HTMLElement) {
		answer = collyElement.Text
	})

	return answer
}
