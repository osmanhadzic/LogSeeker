/*
 *
 * Copyright 2024 OCode
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package analyzer

import (
	"fmt"
	"github.com/spf13/cobra"
	"log-seeker/cmd/parser"
	"log-seeker/cmd/report"
	"strings"
	"time"
)

type LogStats struct {
	TotalLogs   int
	DebugLogs   int
	InfoLogs    int
	WarningLogs int
	ErrorLogs   int
	FatalLogs   int
}

var fileReportPath string

var AnalyzeLogFormFileCmd = &cobra.Command{
	Use:     "analyze <file_path>",
	Short:   "Analyze Log File",
	Long:    "Analyze Log File.",
	Example: `log-seeker file-path`,
	Args:    cobra.ExactArgs(1),
	RunE:    analyzeLogFile,
}

var AnalyzeLogByErrorCodeCmd = &cobra.Command{
	Use:     "analyze-type-log <file_path> <ERROR>",
	Short:   "Analyze Log File error log.",
	Long:    "Analyze Log File by error log. [ERROR]",
	Example: `log-seeker file-path error-code`,
	Args:    cobra.ExactArgs(2),
	RunE:    analyzeLogFileByLogFile,
}

func analyzeLogFile(command *cobra.Command, args []string) error {
	file_path := args[0]

	enterers, err := parser.ParseLog(file_path)

	if err != nil {
		return err
	}

	sates := AnalyzeLogs(enterers)
	
	fmt.Printf("Total Logs: %d\n", sates.TotalLogs)
	fmt.Printf("Debug Logs: %d\n", sates.DebugLogs)
	fmt.Printf("Info Logs: %d\n", sates.InfoLogs)
	fmt.Printf("Warning Logs: %d\n", sates.WarningLogs)
	fmt.Printf("Error Logs: %d\n", sates.ErrorLogs)
	fmt.Printf("Fatal Logs: %d\n", sates.FatalLogs)
	return nil
}

func analyzeLogFileByLogFile(command *cobra.Command, args []string) error {
	file_path := args[0]
	error_code := args[1]

	enterers, err := parser.ParseLog(file_path)

	if err != nil {
		return err
	}

	logs := AnalyzeLogsByErrorCode(enterers, error_code)
	if fileReportPath != "" {
        err = report.PrintLogResult(logs, fileReportPath)
        if err != nil {
            return err
        }
    }else{
		for _, log := range logs {
			fmt.Println(log)
		}
	}
	return nil
}

var AnalyzeLogByDateCmd = &cobra.Command{
	Use:     "analyze-date-log <file_path> <date_time_from> <date_time_to>",
	Short:   "Analyze Log File by date range.",
	Long:    "Analyze Log File by date range. [date_time_from] to [date_time_to]",
	Example: `log-seeker file-path "2024-01-01T00:00:00" "2024-01-02T00:00:00"`,
	Args:    cobra.ExactArgs(3),
	RunE:    analyzeLogFileByDate,
}

func init() {
    AnalyzeLogFormFileCmd.Flags().StringVar(&fileReportPath, "report", "", "Path to save the report file")
    AnalyzeLogByErrorCodeCmd.Flags().StringVar(&fileReportPath, "report", "", "Path to save the report file")
    AnalyzeLogByDateCmd.Flags().StringVar(&fileReportPath, "report", "", "Path to save the report file")
}

func analyzeLogFileByDate(command *cobra.Command, args []string) error {
	file_path := args[0]
	date_time_from := args[1]
	date_time_to := args[2]

	entries, err := parser.ParseLog(file_path)
	if err != nil {
		return err
	}

	logs, err := AnalyzeLogsByDate(entries, date_time_from, date_time_to)
	if fileReportPath != "" {
        err = report.PrintLogResult(logs, fileReportPath)
        if err != nil {
            return err
        }
    }else{
		if err != nil {
			return err
		}
		for _, log := range logs {
			fmt.Println(log)
			fmt.Print("\n")
		}
	}
	return nil
}

func AnalyzeLogsByDate(logs []parser.LogEntry, date_time_from string, date_time_to string) ([]parser.LogEntry, error) {
	from, err := time.Parse(time.RFC3339, date_time_from)
	if err != nil {
		return nil, fmt.Errorf("parsing time from: %w", err)
	}
	to, err := time.Parse(time.RFC3339, date_time_to)
	if err != nil {
		return nil, fmt.Errorf("parsing time to: %w", err)
	}

	filtered_logs := []parser.LogEntry{}
	for _, log := range logs {
		log_time, err := time.Parse(time.RFC3339, log.DateTime)
		if err != nil {
			continue
		}
		if log_time.After(from) && log_time.Before(to) {
			filtered_logs = append(filtered_logs, log)
		}
	}
	return filtered_logs, nil
}

func AnalyzeLogsByErrorCode(logs []parser.LogEntry, error_code string) []parser.LogEntry {
	error_logs := []parser.LogEntry{}
	for _, log := range logs {
		if log.Level == error_code {
			error_logs = append(error_logs, log)
		}
	}
	return error_logs
}
func AnalyzeLogs(entries []parser.LogEntry) LogStats {
	stats := LogStats{}
	for _, entry := range entries {
		stats.TotalLogs++
		switch strings.ToUpper(entry.Level) {
		case "DEBUG":
			stats.DebugLogs++
		case "INFO":
			stats.InfoLogs++
		case "WARNING":
			stats.WarningLogs++
		case "ERROR":
			stats.ErrorLogs++
		case "FATAL":
			stats.FatalLogs++
		}
	}
	return stats
}
