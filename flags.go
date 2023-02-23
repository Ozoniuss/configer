package configer

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
)

// defineFlags defines all the flags that have been introduced through
// configuration options.
func defineFlags(configOptions []ConfigOption) error {

	for _, opt := range configOptions {

		if opt.FlagName == readFlagName() || opt.FlagName == writeFlagName() {
			return fmt.Errorf("cannot use reserved flag name: %s", opt.FlagName)
		}

		if opt.FlagName == "" && opt.Shorthand != "" {
			return fmt.Errorf("defined shorthand %s for flag with no name", opt.Shorthand)
		}

		if opt.FlagName != "" {
			switch opt.Value.(type) {
			case bool:
				pflag.BoolP(opt.FlagName, opt.Shorthand, opt.Value.(bool), opt.Usage)
			case string:
				pflag.StringP(opt.FlagName, opt.Shorthand, opt.Value.(string), opt.Usage)
			case int:
				pflag.IntP(opt.FlagName, opt.Shorthand, opt.Value.(int), opt.Usage)
			case int32:
				pflag.Int32P(opt.FlagName, opt.Shorthand, opt.Value.(int32), opt.Usage)
			case time.Duration:
				pflag.DurationP(opt.FlagName, opt.Shorthand, opt.Value.(time.Duration), opt.Usage)
			default:
				return fmt.Errorf("Invalid flag value provided for option %s", opt.FlagName)
			}
		}
	}
	return nil
}

// defineWriteFlag defines the flag that can be used to write the project
// configuration to the path supplied via the flag value. If an empty path
// is supplied, the working directory is used.
func defineWriteFlags() *string {
	// Do not use a shorthand option to minimize programmer limitations.
	return pflag.String(writeFlagName(), "",
		// Helps with formatting to the console.
		`If supplied, the project configuration is written at the specified
location, and returns an error if the operation is not possible. If 
empty, uses the working directory. If the location is a directory, 
the parser will attempt to read from a file inside that directory, 
with the configured named and type (or the default values if not 
configured)`)
}

// defineReadFlag defines the flag that can be used to explicitly specify
// where the configuration is read from. If the specified location doesn't
// contain a configuration file, the parser returns an error.
//
// If the specified location is a directory, the parser will search for a
// config file with the name and type specified via WithConfigName and
// WithConfigType, or the default values if those were not set. Otherwise,
// if the location is a file, it will try to read the config from that file.
func defineReadFlag() *string {
	// Do not use a shorthand option to minimize programmer limitations.
	return pflag.String(readFlagName(), "",
		// Helps with formatting to the console.
		`If supplied, the parser attempts to read the config from that specified
location, and returns an error if no config file is found there. If empty,
uses the working directory. If the location is a file, the parser tries to
read that file. If the location is a directory, the parser will attempt to 
read from a file inside that directory, with the configured named and type 
(or the default values if not configured)`)
}

func writeFlagName() string {
	return "write-config"
}

func readFlagName() string {
	return "read-config"
}
