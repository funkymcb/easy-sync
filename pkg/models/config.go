package models

var Cfg EasySyncConfig

// EasySyncConfig that represents the configuration for the cli
type EasySyncConfig struct {
	Easyverein EasyvereinCfg `mapstructure:"easyverein"`
	Wordpress  WordpressCfg  `mapstructure:"wordpress"`
}

// EasyvereinCfg struct that represents all required easyverin configurations
type EasyvereinCfg struct {
	URL      string               `mapstructure:"url"`
	Endpoint string               `mapstructure:"endpoint"`
	Token    string               `mapstructure:"token"`
	Options  EasyvereinAPIOptions `mapstructure:"options"`
}

// EasyvereinAPIOptions that represents all easyverein api options
type EasyvereinAPIOptions struct {
	ResultsPerSite int `mapstructure:"results-per-site"`
}

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

// SetConfig sets the config for global access
func SetConfig(config EasySyncConfig) {
	Cfg = config
}

// GetConfig returns the config
func GetConfig() EasySyncConfig {
	return Cfg
}
