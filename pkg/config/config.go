package config

// EasySyncConfig that represents the configuration for the cli
type EasySyncConfig struct {
	Easyverein EasyvereinCfg `mapstructure:"easyverein"`
	Wordpress  WordpressCfg  `mapstructure:"wordpress"`
}
