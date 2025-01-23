package main

import "os"

type config struct {
	addr   string
	dsn    string
	jwtKey string
}

func getEnvString(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}
