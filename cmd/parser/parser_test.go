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
	"os"
	"reflect"
	"strings"
	"testing"
)

func writeTempLogFile(t *testing.T, line string) string {
	t.Helper()

	file, err := os.CreateTemp(t.TempDir(), "log-*.log")
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

func TestStandardFormatNames(t *testing.T) {
	got := StandardFormatNames()
	want := []string{"bracketed", "csv", "kv", "pipe"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected format names: got %v, want %v", got, want)
	}
}

func TestPatternForFormatUnknown(t *testing.T) {
	_, err := PatternForFormat("unknown")
	if err == nil {
		t.Fatal("expected error for unknown format")
	}

	if !strings.Contains(err.Error(), "unknown log format") {
		t.Fatalf("unexpected error message: %v", err)
	}
}

func TestParseLogWithFormat(t *testing.T) {
	tests := []struct {
		name   string
		format string
		line   string
		want   LogEntry
	}{
		{
			name:   "bracketed",
			format: "bracketed",
			line:   "[2026-09-24T10:00:00Z] [INFO] [api] [request completed] [{\"status\":200}]",
			want: LogEntry{
				DateTime: "2026-09-24T10:00:00Z",
				Level:    "INFO",
				Source:   "api",
				Message:  "request completed",
				Metadata: "{\"status\":200}",
			},
		},
		{
			name:   "csv",
			format: "csv",
			line:   "2026-09-24T10:00:00Z,ERROR,worker,job failed,{\"job_id\":42}",
			want: LogEntry{
				DateTime: "2026-09-24T10:00:00Z",
				Level:    "ERROR",
				Source:   "worker",
				Message:  "job failed",
				Metadata: "{\"job_id\":42}",
			},
		},
		{
			name:   "pipe",
			format: "pipe",
			line:   "2026-09-24T10:00:00Z|WARNING|scheduler|retrying task|{\"attempt\":2}",
			want: LogEntry{
				DateTime: "2026-09-24T10:00:00Z",
				Level:    "WARNING",
				Source:   "scheduler",
				Message:  "retrying task",
				Metadata: "{\"attempt\":2}",
			},
		},
		{
			name:   "kv",
			format: "kv",
			line:   "datetime=2026-09-24T10:00:00Z level=DEBUG source=cache message=\"cache miss\" metadata={\"key\":\"u:1\"}",
			want: LogEntry{
				DateTime: "2026-09-24T10:00:00Z",
				Level:    "DEBUG",
				Source:   "cache",
				Message:  "cache miss",
				Metadata: "{\"key\":\"u:1\"}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := writeTempLogFile(t, tt.line)

			entries, err := ParseLogWithFormat(path, tt.format)
			if err != nil {
				t.Fatalf("ParseLogWithFormat returned error: %v", err)
			}

			if len(entries) != 1 {
				t.Fatalf("unexpected number of entries: got %d, want 1", len(entries))
			}

			if !reflect.DeepEqual(entries[0], tt.want) {
				t.Fatalf("unexpected entry: got %+v, want %+v", entries[0], tt.want)
			}
		})
	}
}

func TestParseLogWithPatternInvalidGroupCount(t *testing.T) {
	path := writeTempLogFile(t, "anything")

	_, err := ParseLogWithPattern(path, `^(.*)$`)
	if err == nil {
		t.Fatal("expected error for invalid capture group count")
	}

	if !strings.Contains(err.Error(), "expected 5 capture groups") {
		t.Fatalf("unexpected error message: %v", err)
	}
}
