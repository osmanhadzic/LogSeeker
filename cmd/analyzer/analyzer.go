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
