package domain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func readJSONL[T any](
	path string,
	handle func(line int, value T) error,
) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line++
		if err := scanJSONLLine(scanner.Text(), line, handle); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func scanJSONLLine[T any](
	text string,
	line int,
	handle func(line int, value T) error,
) error {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return nil
	}
	var value T
	if err := json.Unmarshal([]byte(trimmed), &value); err != nil {
		return fmt.Errorf("line %d: %w", line, err)
	}
	return handle(line, value)
}
