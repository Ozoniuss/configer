package configer

import "strings"

type envKeyReplacerOption string

func (r envKeyReplacerOption) apply(parser *configParser) {
	replacer := strings.NewReplacer(".", string(r))
	parser.viper.SetEnvKeyReplacer(replacer)
}

// WithEnvKeyReplacer allows to specify a string replacer in order to use a
// different separator in environment variables than the configuration option
// key. The separator used for configuration keys is "."
//
// e.g. if the replacer is set to "_", the flag bound to "env.variable" can be
// set via the environment variable "ENV_VARIABLE".
//
// By default, the parser uses the "_" separator for environment variables.
func WithEnvKeyReplacer(replacer string) envKeyReplacerOption {
	return envKeyReplacerOption(replacer)
}

type envPrefixOptions string

func (p envPrefixOptions) apply(parser *configParser) {
	parser.viper.SetEnvPrefix(string(p))
}

// WithEnvPrefix allows using a specific prefix for the environment variables
// associated with the project.
//
// e.g. if the prefix is set to "DEMO", for the option parser.option the parser
// associates the "DEMO_PARSER_OPTION" environment variable, assuming the
// replacer is set to "_"
//
// By default, the parser doesn't look for any prefix.
func WithEnvPrefix(prefix string) envPrefixOptions {
	return envPrefixOptions(prefix)
}
