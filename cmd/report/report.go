package report

import (
    "fmt"
    "log-seeker/cmd/analyzer"
)

func GenerateReport(stats analyzer.LogStats) {
    fmt.Printf("Total Logs: %d\n", stats.TotalLogs)
    fmt.Printf("Info Logs: %d\n", stats.InfoLogs)
    fmt.Printf("Warning Logs: %d\n", stats.WarningLogs)
    fmt.Printf("Error Logs: %d\n", stats.ErrorLogs)
}
