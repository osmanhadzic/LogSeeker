package analyzer

import (
    "log-seeker/cmd/parser"
    "strings"
)

type LogStats struct {
    TotalLogs   int
    InfoLogs    int
    WarningLogs int
    ErrorLogs   int
}

func AnalyzeLogs(entries []parser.LogEntry) LogStats {
    stats := LogStats{}
    for _, entry := range entries {
        stats.TotalLogs++
        switch strings.ToUpper(entry.Level) {
        case "INFO":
            stats.InfoLogs++
        case "WARNING":
            stats.WarningLogs++
        case "ERROR":
            stats.ErrorLogs++
        }
    }
    return stats
}
