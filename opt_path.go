package config

type configPathOption string

func (opt configPathOption) apply(parser *configParser) {
	parser.viper.AddConfigPath(string(opt))
}

// WithConfigPath allows specifying paths where the parser should search for
// a configuration options file. The format of this configuration file may be
// defined using the functions WithConfigType and WithConfigName.
//
// By default, the parser searches for configuration files in the project
// directory.
func WithConfigPath(path string) configPathOption {
	return configPathOption(path)
}
