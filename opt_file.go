package configer

type configFileOption string

func (opt configFileOption) apply(parser *configParser) {
	parser.viper.SetConfigFile(string(opt))
}

// WithConfigFile allows providing a specific configuration file for the parser.
// This will overwrite all the options that have been specified through
// WithConfigPath, including the default ones.
//
// This option is not set by default.
func WithConfigFile(configFile string) configFileOption {
	return configFileOption(configFile)
}
