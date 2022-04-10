package config

func IsLocal() bool {
	return GetString("app.env") == "local"
}

func IsProduction() bool {
	return GetString("app.env") == "production"
}

func IsTesting() bool {
	return GetString("app.env") == "testing"
}
