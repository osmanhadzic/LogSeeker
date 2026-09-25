/*
 *
 * Copyright 2026 OCode
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

package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type LogEntry struct {
	DateTime string
	Level    string
	Source   string
	Message  string
	Metadata string
}

var standardPatterns = map[string]string{
	"bracketed": `\[(.*?)\] \[(.*?)\] \[(.*?)\] \[(.*?)\] \[(.*?)\]`,
	"csv":       `^([^,]*),([^,]*),([^,]*),([^,]*),(.*)$`,
	"pipe":      `^([^|]*)\|([^|]*)\|([^|]*)\|([^|]*)\|(.*)$`,
	"kv":        `^datetime=(\S+)\s+level=(\S+)\s+source=(\S+)\s+message="(.*?)"\s+metadata=(.*)$`,
}

func ParseLog(filePath string) ([]LogEntry, error) {
	defaultPattern := standardPatterns["bracketed"]
	return ParseLogWithPattern(filePath, defaultPattern)
}

func ParseLogWithFormat(filePath string, format string) ([]LogEntry, error) {
	pattern, err := PatternForFormat(format)
	if err != nil {
		return nil, err
	}

	return ParseLogWithPattern(filePath, pattern)
}

func PatternForFormat(format string) (string, error) {
	key := strings.ToLower(strings.TrimSpace(format))
	pattern, ok := standardPatterns[key]
	if !ok {
		return "", fmt.Errorf("unknown log format %q. available formats: %s", format, strings.Join(StandardFormatNames(), ", "))
	}

	return pattern, nil
}

func StandardFormatNames() []string {
	names := make([]string, 0, len(standardPatterns))
	for name := range standardPatterns {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func ParseLogWithPattern(filePath string, pattern string) ([]LogEntry, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %w", err)
	}

	if re.NumSubexp() != 5 {
		return nil, fmt.Errorf("invalid regex pattern: expected 5 capture groups, got %d", re.NumSubexp())
	}

	return parseLogWithRegex(filePath, re)
}

func parseLogWithRegex(filePath string, re *regexp.Regexp) ([]LogEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []LogEntry
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 6 {
			entries = append(entries, LogEntry{
				DateTime: matches[1],
				Level:    matches[2],
				Source:   matches[3],
				Message:  matches[4],
				Metadata: matches[5],
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
