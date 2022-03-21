package config

// EasyvereinCfg struct that represents all required easyverin configurations
type EasyvereinCfg struct {
	Endpoint string               `mapstructure:"endpoint"`
	Token    string               `mapstructure:"token"`
	Options  EasyvereinAPIOptions `mapstructure:"options"`
}

// EasyvereinAPIOptions that represents all easyverein api options
type EasyvereinAPIOptions struct {
	ResultsPerSite int `mapstructure:"results-per-site"`
}
