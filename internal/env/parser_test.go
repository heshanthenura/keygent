package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsValidEnvLine(t *testing.T) {
	tests := []struct {
		line     string
		expected bool
	}{
		{"ENV=VALUE", true},
		{"ENV=\"VALUE\"", true},
		{"ENV=", true},
		{"ENV", false},
	}

	for _, test := range tests {
		result := isValidEnvLine(test.line)
		if result != test.expected {
			t.Errorf("isValidEnvLine(%q) = %v, want %v", test.line, result, test.expected)
		}
	}
}

func TestParseEnvFile_WrongFilePath(t *testing.T) {
	_, err := ParseEnvFile("this-file-should-never-exist.env")

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseEnvFile_IsDirectory(t *testing.T) {
	dir := t.TempDir()
	_, err := ParseEnvFile(dir)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseEnvFile_IgnoreComment(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, ".env")

	content := " # NAME=Heshan\n # AGE=22"

	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	env, err := ParseEnvFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if len(env) != 0 {
		t.Fatal("expected no environment variables")

	}
}

func TestParseEnvFile_InvalidLine(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, ".env")

	content := " # NAME+=Heshan\n # AGE=\n age"

	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	if err != nil {
		t.Fatal(err)
	}

	_, err = ParseEnvFile(path)

	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
