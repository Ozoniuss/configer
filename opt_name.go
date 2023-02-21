package config

type configNameOption string

func (opt configNameOption) apply(parser *configParser) {
	parser.changeConfigName(string(opt))
	parser.viper.SetConfigName(string(opt))
}

// WithConfigName allows specifying the name of the configuration file the
// parser looks for. Currently, the parser supports searching for a single
// file name.
//
// By default, the parser searches for a configuration file named "config".
func WithConfigName(configName string) configNameOption {
	return configNameOption(configName)
}
