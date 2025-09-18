package config

import (
	"log"
	"os"
	"strconv"
)

// helper to get env with default value
func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func getEnvAsInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
		log.Printf("Invalid value for %s: %s, using default %d", key, value, defaultVal)
	}
	return defaultVal
}

// helper to split comma-separated env vars
func splitEnv(val string) []string {
	var result []string
	for _, s := range splitAndTrim(val, ",") {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}

// split string by sep and trim spaces
func splitAndTrim(s, sep string) []string {
	raw := make([]string, 0)
	for _, part := range Split(s, sep) {
		raw = append(raw, TrimSpace(part))
	}
	return raw
}

// minimal wrappers to avoid extra imports
func Split(s, sep string) []string {
	var result []string
	curr := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == sep {
			result = append(result, curr)
			curr = ""
		} else {
			curr += string(s[i])
		}
	}
	result = append(result, curr)
	return result
}

func TrimSpace(s string) string {
	start := 0
	end := len(s) - 1
	for start <= end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n') {
		start++
	}
	for end >= start && (s[end] == ' ' || s[end] == '\t' || s[end] == '\n') {
		end--
	}
	if start > end {
		return ""
	}
	return s[start : end+1]
}
