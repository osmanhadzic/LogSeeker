package parser

import (
	"bufio"
	"os"
	"regexp"
)

type LogEntry struct {
	DateTime string
	Level    string
	Message  string
}

func ParseLog(filePath string) ([]LogEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []LogEntry
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\[(.*?)\] \[(.*?)\] (.*)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 4 {
			entries = append(entries, LogEntry{
				DateTime: matches[1],
				Level:    matches[2],
				Message:  matches[3],
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
