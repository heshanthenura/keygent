package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValidKey(key string) bool {
	if key == "" {
		return false
	}

	for _, ch := range key {
		if !(ch == '_' ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= 'a' && ch <= 'z') ||
			(ch >= '0' && ch <= '9')) {
			return false
		}
	}

	return true
}

func isValidEnvLine(line string) bool {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return false
	}

	key := strings.TrimSpace(parts[0])
	if !isValidKey(key) {
		return false
	}

	value := strings.TrimSpace(parts[1])

	_ = value

	return true
}

func ParseEnvFile(filePath string) ([]string, error) {
	var envs []string

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("env file does not exist: %s", filePath)
	}
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, fmt.Errorf("env file is a directory: %s", filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open env file: %s", filePath)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if !isValidEnvLine(line) {
			return nil, fmt.Errorf("invalid env line: %q", line)
		}

		envs = append(envs, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %s", filePath)
	}

	return envs, nil
}
