package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, price, change string
}

func main() {
	ticker := []string{
		"MSFT",
		"IBM",
		"GOOGL",
		"AAPL",
		"GE",
		"AMZN",
		"TSLA",
		"FB",
		"UNP",
		"MCD",
		"VZ",
		"INTC",
	}

	stocks := []Stock{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla 5.0")
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Request failed: ", err)
	})

	c.OnHTML("#main-content-wrapper", func(e *colly.HTMLElement) {
		stock := Stock{}

		stock.company = e.ChildText("h1")
		stock.price = e.ChildText("[data-testid='qsp-price']")
		stock.change = e.ChildText("[data-testid='qsp-price-change']")

		stocks = append(stocks, stock)
	})

	c.Wait() // Wait for all requests to finish

	for _, t := range ticker {
		c.Visit(("https://finance.yahoo.com/quote/" + t))
	}

	for _, stock := range stocks {
		fmt.Printf("Company: %s, Price: %s, Change: %s\n", stock.company, stock.price, stock.change)
	}

	log.Println("Scraping completed successfully.")
	log.Println("Total stocks scraped:", len(stocks))

	file, err := os.Create("stocks.csv")

	if err != nil {
		log.Fatalln("Failed to create file:", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{"Company", "Price", "Change"}

	writer.Write(headers)

	for _, stock := range stocks {
		record := []string{
			stock.company,
			stock.price,
			stock.change,
		}

		writer.Write(record)
		defer writer.Flush()
	}
}
