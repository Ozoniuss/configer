package configer

type configTypeOption string

func (opt configTypeOption) apply(parser *configParser) {
	parser.changeConfigExtension(string(opt))
	parser.viper.SetConfigType(string(opt))
}

// WithConfigType allows specifying the type of the configuration file the
// parser looks for. Currently, the parser supports searching for a single
// file type.
//
// By default, the parser searches for yaml configuration files.
func WithConfigType(configType string) configTypeOption {
	return configTypeOption(configType)
}
