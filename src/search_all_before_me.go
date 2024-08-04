package main

import (
	"bytes"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"log"
	"strconv"
)

func main() {
	var identityNumber string
	var url string

	fmt.Print("Введите ссылку на рейтинговый список (пример: https://abit.itmo.ru/rating/master/budget/1905): ")
	_, errUrl := fmt.Scan(&url)
	if errUrl != nil {
		log.Fatal(errUrl)
	}

	fmt.Print("Введите СНИЛС: ")
	_, errSnils := fmt.Scan(&identityNumber)
	if errSnils != nil {
		log.Fatal(errSnils)
	}

	var p, r = searchByIdentity(url, identityNumber)
	fmt.Printf("\nВаш порядковый номер: %s\nВаш рейтинговый номер: %d", p, r)
}

func searchByIdentity(url string, identityNumber string) (string, int) {
	var d string
	var rating = 0

	positionXPath := "//*[contains(@class, 'RatingPage_table__position')]"
	itemXPath := "//*[contains(@class, 'RatingPage_table__item')]"
	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		r := bytes.NewReader(e.Response.Body)
		doc, err := htmlquery.Parse(r)
		throwFatal(err)

		positionNodes, err := htmlquery.QueryAll(doc, positionXPath)
		throwFatal(err)

		itemNodes, err := htmlquery.QueryAll(doc, itemXPath)
		throwFatal(err)

		for _, node := range positionNodes {
			if spanData := node.LastChild.FirstChild.Data; spanData == identityNumber {
				d = node.FirstChild.Data
				break
			}
		}

		for _, node := range itemNodes {
			pData := node.FirstChild.FirstChild.FirstChild.FirstChild.Data
			isHasOriginal := node.LastChild.LastChild.LastChild.LastChild.LastChild.LastChild.Data
			isFirstPriority := node.FirstChild.NextSibling.FirstChild.FirstChild.FirstChild.LastChild.FirstChild.Data
			intPData, err1 := strconv.Atoi(pData)
			throwFatal(err1)

			intD, err2 := strconv.Atoi(d)
			throwFatal(err2)

			if intPData <= intD && isHasOriginal == "да" && isFirstPriority == "1" {
				rating++
			}
		}
	})

	// Start scraping
	err := c.Visit(url)
	throwFatal(err)
	return d, rating
}

func throwFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
