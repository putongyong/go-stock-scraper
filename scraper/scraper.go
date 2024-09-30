package scraper

import (
	"fmt"
	"strings"
	"github.com/gocolly/colly"
)

// ScrapeStockData is a function that scrapes the stock data for a given ticker symbol
func ScrapeStockData(tickerSymbol string) map[string]string {
	// Initialize the collector
	c := colly.NewCollector()

	// Create a map to store the scraped data
	stock := make(map[string]string)

	// Set up the selectors for regular market data
	regularMarketSelectors := map[string]string{
		"regular_market_price":           fmt.Sprintf(`[data-symbol="%s"][data-field="regularMarketPrice"]`, tickerSymbol),
		"regular_market_change":          fmt.Sprintf(`[data-symbol="%s"][data-field="regularMarketChange"]`, tickerSymbol),
		"regular_market_change_percent":  fmt.Sprintf(`[data-symbol="%s"][data-field="regularMarketChangePercent"]`, tickerSymbol),
	}

	// Scraping the regular market data
	for key, selector := range regularMarketSelectors {
		c.OnHTML(selector, func(e *colly.HTMLElement) {
			value := strings.TrimSpace(e.Text)

			// Clean up percent values by removing parentheses
			if key == "regular_market_change_percent" {
				value = strings.ReplaceAll(value, "(", "")
				value = strings.ReplaceAll(value, ")", "")
			}

			// Avoid overwriting existing data (debugging log if duplicate detected)
			if _, exists := stock[key]; exists {
				return
			}

			// Store the data in the map
			stock[key] = value
		})
	}

	// Visit the page and handle any potential errors
	err := c.Visit("https://finance.yahoo.com/quote/" + tickerSymbol)
	if err != nil {
		fmt.Printf("Error visiting page for %s: %v\n", tickerSymbol, err)
	}

	return stock
}
