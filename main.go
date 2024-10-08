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

package main

import (
	// "fmt"
	// "log-seeker/cmd/analyzer"
	// "log-seeker/cmd/parser"
	// "log-seeker/cmd/report"
	"log-seeker/cmd"
)

func main() {
	// entries, err := parser.ParseLog("logs/sample.log")
	// if err != nil {
	// 	fmt.Printf("Error parsing log file: %v\n", err)
	// 	return
	// }

	// stats := analyzer.AnalyzeLogs(entries)
	// report.GenerateReport(stats)

	cmd.Execute()
}
