package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("¿Qué datos deseas obtener desde Ripley.cl?")
	var query string
	fmt.Scanln(&query)

	fmt.Println("¿Cuántas páginas quieres leer?")
	var pages int
	_, err := fmt.Scanf("%d", &pages)
	if err != nil {
		log.Fatal(err)
	}

	parsed_data := make([][]string, 1)
	for i := 0; i < pages; i++ {
		url := "https://simple.ripley.cl/search/" + query + "?sort=score&page=" + strconv.Itoa(pages)
		doc := scrape_page(url)
		parsed_page_data := parse_data(doc)

		parsed_data = append(parsed_data, parsed_page_data...)
	}
	save_parsed_to_csv(query, parsed_data)
}

func scrape_page(url string) *goquery.Document {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func parse_data(doc *goquery.Document) [][]string {
	docs := doc.Find(".catalog-container").Children()

	data := make([][]string, 1)
	data[0] = append(data[0], "name")
	data[0] = append(data[0], "description")
	data[0] = append(data[0], "price")

	docs.Each(func(i int, s *goquery.Selection) {
		name := s.Find(".catalog-product-details__name").Text()

		dif_prices := s.Find(".catalog-product-details__prices").Text()
		price := strings.Split(dif_prices, "$")[1]
		brand := s.Find(".catalog-product-details").Find("span").First().Text()

		data = append(data, []string{name, brand, price})
	})

	return data
}

func save_parsed_to_csv(query string, data [][]string) {
	csvFile, err := os.Create(query + ".csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range data {
		_ = csvwriter.Write(empRow)
	}

	csvwriter.Flush()

	csvFile.Close()
}
