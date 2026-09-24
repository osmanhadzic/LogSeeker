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

package analyzer

import (
	"os"
	"strings"
	"testing"
)

func writeAnalyzerTempLogFile(t *testing.T, line string) string {
	t.Helper()

	file, err := os.CreateTemp(t.TempDir(), "analyzer-log-*.log")
	if err != nil {
		t.Fatalf("failed to create temp log file: %v", err)
	}

	if _, err := file.WriteString(line + "\n"); err != nil {
		_ = file.Close()
		t.Fatalf("failed to write temp log file: %v", err)
	}

	if err := file.Close(); err != nil {
		t.Fatalf("failed to close temp log file: %v", err)
	}

	return file.Name()
}

func withParseFlags(pattern, format string) func() {
	oldPattern := logPattern
	oldFormat := logFormat

	logPattern = pattern
	logFormat = format

	return func() {
		logPattern = oldPattern
		logFormat = oldFormat
	}
}

func TestParseLogFileWithConfiguredPatternConflict(t *testing.T) {
	restore := withParseFlags(`^(.*)$`, "csv")
	defer restore()

	_, err := parseLogFileWithConfiguredPattern("irrelevant.log")
	if err == nil {
		t.Fatal("expected error when both --pattern and --format are set")
	}

	if !strings.Contains(err.Error(), "use either --pattern or --format") {
		t.Fatalf("unexpected error message: %v", err)
	}
}

func TestParseLogFileWithConfiguredPatternUsesFormat(t *testing.T) {
	restore := withParseFlags("", "csv")
	defer restore()

	path := writeAnalyzerTempLogFile(t, "2026-09-24T10:00:00Z,INFO,api,ready,{\"ok\":true}")

	entries, err := parseLogFileWithConfiguredPattern(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(entries) != 1 {
		t.Fatalf("unexpected number of entries: got %d, want 1", len(entries))
	}

	if entries[0].Level != "INFO" {
		t.Fatalf("unexpected parsed level: got %q, want %q", entries[0].Level, "INFO")
	}
}

func TestParseLogFileWithConfiguredPatternUsesCustomPattern(t *testing.T) {
	pattern := `^(\S+)\s+(\S+)\s+(\S+)\s+(.*?)\s+(\{.*\})$`
	restore := withParseFlags(pattern, "")
	defer restore()

	path := writeAnalyzerTempLogFile(t, "2026-09-24T10:00:00Z ERROR worker failed_job {\"code\":500}")

	entries, err := parseLogFileWithConfiguredPattern(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(entries) != 1 {
		t.Fatalf("unexpected number of entries: got %d, want 1", len(entries))
	}

	if entries[0].Source != "worker" {
		t.Fatalf("unexpected parsed source: got %q, want %q", entries[0].Source, "worker")
	}
}
