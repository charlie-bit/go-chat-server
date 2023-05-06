package config

var Cfg *Config

// Config is the configuration
type Config struct {
	Debug bool `yaml:"debug"`

	Env  string `yaml:"-"`
	File string `yaml:"-"`

	Stat `yaml:",inline"`
}

// Stat config
type Stat struct {
	Addr      string `yaml:"addr"`
	HealthURI string `yaml:"health_uri"`
}
