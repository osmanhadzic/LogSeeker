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
 
package report

import (
	"log-seeker/cmd/parser"
	"fmt"
	"os"
)

// PrintLogResult prints the log result to a file
func PrintLogResult(result []parser.LogEntry, filePath string) error {
	// Create or open the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	for _, entry := range result {
		_, err := file.WriteString(fmt.Sprintf("%s\n", entry))
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}
	

	fmt.Println("Log result written to file successfully")
	return nil
}
