package analyzer

import (
    "log-seeker/cmd/parser"
    "strings"
    "fmt"
    "github.com/spf13/cobra"
)

type LogStats struct {
    TotalLogs   int
    InfoLogs    int
    WarningLogs int
    ErrorLogs   int
}

var AnalyzeLogFormFileCmd = &cobra.Command{
	Use:     "analyze <file_path>",
	Short:   "Analyze Log File",
	Long:    "Analyze Log File.",
	Example: `log-seeker file-path`,
	Args:    cobra.ExactArgs(1),
	RunE:    analyzeLogFile,
}

func analyzeLogFile(command *cobra.Command, args []string) error {
	file_path := args[0]

	enterers, err := parser.ParseLog(file_path);

	if err != nil {
		return err;
	}

	sates := AnalyzeLogs(enterers)
    fmt.Printf("Total Logs: %d\n", sates.TotalLogs)
    fmt.Printf("Info Logs: %d\n", sates.InfoLogs)
    fmt.Printf("Warning Logs: %d\n", sates.WarningLogs)
    fmt.Printf("Error Logs: %d\n", sates.ErrorLogs)
	return nil
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
