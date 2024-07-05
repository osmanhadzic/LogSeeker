package main

import (
	"fmt"
	"log-seeker/cmd/analyzer"
	"log-seeker/cmd/parser"
	"log-seeker/cmd/report"
)

func main() {
	entries, err := parser.ParseLog("logs/sample.log")
	if err != nil {
		fmt.Printf("Error parsing log file: %v\n", err)
		return
	}

	stats := analyzer.AnalyzeLogs(entries)
	report.GenerateReport(stats)
}
