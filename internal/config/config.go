package config

type HTTPServer struct {
	Addr string
}

// Configuration holds the application configuration settings.

// env-default:"production"
type Configuration struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server" env-required:"true"`
}
