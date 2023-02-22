package configer

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type configParser struct {
	viper      *viper.Viper
	readFlag   bool
	writeFlag  bool
	configName string
}

// newParser initializes a project parser with some default options.
func newParser() *configParser {
	parser := &configParser{
		viper:      viper.New(),
		writeFlag:  false,
		readFlag:   false,
		configName: "config.yml",
	}

	return parser
}

// setDefaultParserOptions sets the default parser options that are often good
// enough for most projects.
func (p *configParser) setDefaultParserOptions() {
	// Looks for config.yml in the current directory by default.
	p.viper.SetConfigName("config")
	p.viper.SetConfigType("yml")

	p.viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	p.viper.AutomaticEnv()
}

// applyOptions applies the parser options that were supplied to the parser.
func (p *configParser) applyOptions(options ...ParserOption) {
	for _, opt := range options {
		opt.apply(p)
	}
}

// setDefaultValues sets the default values of the config, specified through
// ConfigOptions. These values may be overwritten, in this order of precedence,
// by flags, environment variables and configuration files. Any configuration
// option must provide a default value.
func (p *configParser) setDefaultValues(opts []ConfigOption) error {
	for _, opt := range opts {

		// "Special" configuration options that can only be set through flags.
		if opt.ConfigKey == "" {
			continue
		}
		p.viper.SetDefault(opt.ConfigKey, opt.Value)

		// Bind to the defined flags. Flags may be left empty.
		//
		// Flags must be configured individually. viper comes with a function
		// called BindPFlags to bind all flags at once, but that cannot be used
		// since it uses each flag's full name as the config key
		// (see function documentation).
		if f := pflag.Lookup(opt.FlagName); f != nil {
			err := p.viper.BindPFlag(opt.ConfigKey, f)
			if err != nil {
				return fmt.Errorf("could not bind to flag %s: %w", opt.FlagName, err)
			}
		}
	}
	return nil
}

// changeConfigName is a helper method that changes the internal config file
// name stored by the parser.
func (p *configParser) changeConfigName(name string) {
	parts := strings.Split(p.configName, ".")
	if len(parts) != 2 {
		panic("internal parser error: invalid config name")
	}
	parts[0] = name
	p.configName = strings.Join(parts, ".")
}

// changeConfigExtension is a helper method that changes the internal config
// file extension stored by the parser.
func (p *configParser) changeConfigExtension(extension string) {
	parts := strings.Split(p.configName, ".")
	if len(parts) != 2 {
		panic("internal parser error: invalid config name")
	}
	parts[1] = extension
	p.configName = strings.Join(parts, ".")
}
