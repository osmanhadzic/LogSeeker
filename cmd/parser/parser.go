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

package parser

import (
	"bufio"
	"os"
	"regexp"
)

type LogEntry struct {
	DateTime string
	Level    string
	Source string
	Message  string
	Metadata string
}

func ParseLog(filePath string) ([]LogEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []LogEntry
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\[(.*?)\] \[(.*?)\] \[(.*?)\] \[(.*?)\] \[(.*?)\]`)

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
