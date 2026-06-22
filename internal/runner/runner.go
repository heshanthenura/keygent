package runner

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func stripQuotes(v string) string {
	v = strings.TrimSpace(v)

	if len(v) >= 2 {
		if (v[0] == '"' && v[len(v)-1] == '"') ||
			(v[0] == '\'' && v[len(v)-1] == '\'') {
			return v[1 : len(v)-1]
		}
	}

	return v
}

func Run(envVars []string, commands []string) error {
	if len(commands) == 0 {
		return fmt.Errorf("no command provided")
	}

	cmd := exec.Command(commands[0], commands[1:]...)

	envMap := make(map[string]string)

	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}

	for _, e := range envVars {
		e = strings.TrimSpace(e)

		parts := strings.SplitN(e, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid env var: %q", e)
		}

		key := strings.TrimSpace(parts[0])
		value := stripQuotes(parts[1])

		if key == "" {
			return fmt.Errorf("invalid env key in: %q", e)
		}

		envMap[key] = value
	}

	keys := make([]string, 0, len(envMap))
	for k := range envMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	env := make([]string, 0, len(envMap))
	for _, k := range keys {
		env = append(env, k+"="+envMap[k])
	}

	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
