package config

import "os"

type Config struct {
	Addr     string
	CertFile string
	KeyFile  string
	DSN      string
}

// envOrDefault возвращает значение переменной окружения или значение по умолчанию.
func envOrDefault(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func New() Config {
	return Config{
		Addr:     envOrDefault("SERVER_ADDR", ":8443"),
		CertFile: envOrDefault("CERT_FILE", "certs/server.crt"),
		KeyFile:  envOrDefault("KEY_FILE", "certs/server.key"),
		DSN:      envOrDefault("DATABASE_DSN", "postgres://postgres:postgres@localhost:5434/study_security?sslmode=disable"),
	}
}