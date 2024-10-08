package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// GetCurrentTimestamp returns the current timestamp in a standard format
func GetCurrentTimestamp() string {
	// Format as YYYY-MM-DD HH:MM:SS
	return time.Now().Format("2006-01-02 15:04:05")
}

// ReadTickersFromFile reads a list of tickers from a file and returns a slice of tickers
func ReadTickersFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open ticker file: %v", err)
	}
	defer file.Close()

	var tickers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			tickers = append(tickers, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading ticker file: %v", err)
	}

	return tickers, nil
}