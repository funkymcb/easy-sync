package config

// WordpressCfg struct that contains all required wordpress api configurations
type WordpressCfg struct {
	Host        string              `mapstructure:"host"`
	Endpoint    string              `mapstructure:"endpoint"`
	APIUsername string              `mapstructure:"user"`
	APIPassword string              `mapstructure:"pass"`
	DefaultPass string              `mapstructure:"default-password"`
	Options     WordpressAPIOptions `mapstructure:"options"`
}

// WordpressAPIOptions struct that contains wordpress api options
type WordpressAPIOptions struct {
	ResultsPerSite int    `mapstructure:"results-per-site"`
	Context        string `mapstructure:"context"`
}
