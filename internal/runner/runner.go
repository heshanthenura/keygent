package runner

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Run(envVars []string, commands []string) error {
	if len(commands) == 0 {
		return fmt.Errorf("no command provided")
	}

	cmd := exec.Command(commands[0], commands[1:]...)

	envMap := map[string]string{}

	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}

	for _, e := range envVars {
		e = strings.TrimSpace(e)

		if !strings.Contains(e, "=") {
			return fmt.Errorf("invalid env var: %q", e)
		}

		parts := strings.SplitN(e, "=", 2)
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		value = strings.Trim(value, "\"'")

		envMap[key] = value
	}

	env := make([]string, 0, len(envMap))
	for k, v := range envMap {
		env = append(env, k+"="+v)
	}

	cmd.Env = env

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
