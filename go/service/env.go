package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Environment structure responsible for storing the info of a .env file
type Environment struct {
	Path string
	Map  map[string]string
}

func checkEnvFile(path string) ([]string, error) {
	var lines []string
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, "=")

		if len(splitted) == 0 || len(splitted) != 2 {
			break
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func buildEnvironment(lines []string) map[string]string {
	m := make(map[string]string)

	for _, line := range lines {
		info := strings.Split(line, "=")
		m[info[0]] = info[1]
	}

	return m
}

// NewEnv constructs a new environment from a path
func NewEnv(path string) (*Environment, error) {
	lines, err := checkEnvFile(path)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("that's not a valid env file")
	}

	env := Environment{
		Path: path,
		Map:  buildEnvironment(lines),
	}

	return &env, nil
}

// GetEnvVariable returns an available env variable
func (e *Environment) GetEnvVariable(variable string) (string, error) {
	var env string
	var ok bool
	if env, ok = e.Map[variable]; !ok {
		return "", fmt.Errorf("'%s' is not an environment variable", variable)
	}
	return env, nil
}
