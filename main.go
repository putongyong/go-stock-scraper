package main

import (
	"fmt"
	"log"
	"sync"
	"time" // Import time package for scheduling
	"github.com/putongyong/go-stock-scraper/scraper"
	"github.com/putongyong/go-stock-scraper/utils"
)

const interval int = 10

func main() {
	// Create a schedule that triggers every 30 seconds
	schedule := time.NewTicker(time.Duration(interval) * time.Second)
	defer schedule.Stop() // Stop the schedule when done

	// Run the task immediately at the start
	fmt.Println("-------------START---------------")
	executeTask()

	// Loop to run the task every time the schedule ticks
	for range schedule.C {
		fmt.Printf("-------------%d seconds---------------\n", interval)
		executeTask()
	}
}

func executeTask() {
	// Read tickers from the file
	tickers, err := utils.ReadTickersFromFile("tickers.txt")
	if err != nil {
		log.Fatalf("Error reading tickers: %v", err)
	}

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	for _, tickerSymbol := range tickers {
		// Increment the WaitGroup counter
		wg.Add(1)

		// Launch Printtickers as a goroutine
		go func(ticker string) {
			defer wg.Done() // Decrement the counter when the goroutine completes
			Printtickers(ticker)
		}(tickerSymbol)
	}

	// Wait for all goroutines to complete
	wg.Wait()
}

func Printtickers(tickerSymbol string) {
	// Get the current timestamp from the utility package
	currentTime := utils.GetCurrentTimestamp()

	// Scrape the stock data using the scraper package
	stockData := scraper.ScrapeStockData(tickerSymbol)

	// Print the timestamp and stock data
	fmt.Printf("Stock Data for %s:\n%s\n", tickerSymbol, currentTime)
	for key, value := range stockData {
		fmt.Printf("%s: %s\n", key, value)
	}
	fmt.Println("----------------------------")
}
