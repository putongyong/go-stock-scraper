package main

import (
	"fmt"
	"strings"
	"github.com/gocolly/colly"
)

func scrapeStockData(tickerSymbol string) map[string]string {
	// Initialize the collector
	c := colly.NewCollector()

	// Create a map to store the scraped data
	stock := make(map[string]string)
	stock["ticker"] = tickerSymbol

	// Set up the selectors for regular and post-market data
	regularMarketSelectors := map[string]string{
		"regular_market_price":           fmt.Sprintf(`[data-symbol="%s"][data-field="regularMarketPrice"]`, tickerSymbol),
		"regular_market_change":          fmt.Sprintf(`[data-symbol="%s"][data-field="regularMarketChange"]`, tickerSymbol),
		"regular_market_change_percent":  fmt.Sprintf(`[data-symbol="%s"][data-field="regularMarketChangePercent"]`, tickerSymbol),
		"post_market_price":              fmt.Sprintf(`[data-symbol="%s"][data-field="postMarketPrice"]`, tickerSymbol),
		"post_market_change":             fmt.Sprintf(`[data-symbol="%s"][data-field="postMarketChange"]`, tickerSymbol),
		"post_market_change_percent":     fmt.Sprintf(`[data-symbol="%s"][data-field="postMarketChangePercent"]`, tickerSymbol),
	}

	// Selectors for summary table
	summaryTableSelectors := map[string]string{
		"previous_close":    `[data-test="PREV_CLOSE-value"]`,
		"open_value":        `[data-test="OPEN-value"]`,
		"bid":               `[data-test="BID-value"]`,
		"ask":               `[data-test="ASK-value"]`,
		"days_range":        `[data-test="DAYS_RANGE-value"]`,
		"week_range":        `[data-test="FIFTY_TWO_WK_RANGE-value"]`,
		"volume":            `[data-test="TD_VOLUME-value"]`,
		"avg_volume":        `[data-test="AVERAGE_VOLUME_3MONTH-value"]`,
		"market_cap":        `[data-test="MARKET_CAP-value"]`,
		"beta":              `[data-test="BETA_5Y-value"]`,
		"pe_ratio":          `[data-test="PE_RATIO-value"]`,
		"eps":               `[data-test="EPS_RATIO-value"]`,
		"earnings_date":     `[data-test="EARNINGS_DATE-value"]`,
		"dividend_yield":    `[data-test="DIVIDEND_AND_YIELD-value"]`,
		"ex_dividend_date":  `[data-test="EX_DIVIDEND_DATE-value"]`,
		"year_target_est":   `[data-test="ONE_YEAR_TARGET_PRICE-value"]`,
	}

	// Scraping the regular and post market data
	for key, selector := range regularMarketSelectors {
		c.OnHTML(selector, func(e *colly.HTMLElement) {
			value := strings.TrimSpace(e.Text)
			if key == "regular_market_change_percent" || key == "post_market_change_percent" {
				value = strings.ReplaceAll(value, "(", "")
				value = strings.ReplaceAll(value, ")", "")
			}
			stock[key] = value
		})
	}

	// Scraping the summary table data
	for key, selector := range summaryTableSelectors {
		c.OnHTML(selector, func(e *colly.HTMLElement) {
			stock[key] = strings.TrimSpace(e.Text)
		})
	}

	// Visit the page (replace the URL with the correct one)
	c.Visit("https://finance.yahoo.com/quote/"+ tickerSymbol) // Update with the actual URL

	// Wait for asynchronous scraping to finish
	c.Wait()

	return stock
}

func main() {
	// Replace this with the actual ticker symbol
	tickers := []string{"AAPL", "GOOGL", "AMZN", "MSFT", "TSLA", "FB", "NVDA", "PYPL", "ADBE", "NFLX"}
	for _, tickerSymbol := range tickers {
		stockData := scrapeStockData(tickerSymbol)

		// Print the scraped stock data
		for key, value := range stockData {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

}