package config

type HTTPServer struct {
	Address string
}

// Configuration holds the application configuration settings.
type Configuration struct {
	Env         string
	StoragePath string
	HTTPServer
}