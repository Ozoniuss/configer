package configer

import (
	"fmt"
	"os"
	"path"
	"strings"

	log "github.com/Ozoniuss/stdlog"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// NewConfig generates a new configuration setting for the project, based on
// the provided config options. It unmarshals the options to the provided
// struct, which can then be used in the project to read those options.
func NewConfig(appOptions []ConfigOption, configStruct interface{}, parserOptions ...ParserOption) error {

	parser := newParser()
	parser.setDefaultParserOptions()
	parser.applyOptions(parserOptions...)

	// Define the flags after applying the options to allow defining special
	// flags as well.
	err := defineFlags(appOptions)
	if err != nil {
		return fmt.Errorf("unable to define flags: %w", err)
	}

	var readFlag, writeFlag *string
	if parser.readFlag {
		readFlag = defineReadFlag()
	}
	if parser.writeFlag {
		writeFlag = defineWriteFlags()
	}
	pflag.Parse()

	err = parser.setDefaultValues(appOptions)
	if err != nil {
		return fmt.Errorf("could not set default values: %w", err)
	}

	// Config path was supplied explicitly via flag.
	if readFlag != nil && pflag.Lookup(readFlagName()).Changed {
		if *readFlag == "" {
			*readFlag = "."
		}
		stat, err := os.Stat(*readFlag)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("no config file found at location %s", *readFlag)
			}
			return fmt.Errorf("could not retrieve stats for %s: %w", *readFlag, err)
		}

		configpath := *readFlag

		// If the path supplied is a directory, append the config name at the
		// end.
		if stat.IsDir() {
			configpath = path.Join(configpath, parser.configName)
		}

		configfile, err := os.Open(configpath)
		defer configfile.Close()
		if err != nil {
			return fmt.Errorf("could not open config file: %w", err)
		}

		if err = parser.viper.ReadConfig(configfile); err != nil {
			return fmt.Errorf("could not parse config from file %s: %w", configpath, err)
		}

		// No explicit config path set, use the values provided via
		// WithConfigPath.
	} else {
		if err := parser.viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				additional := ""
				if writeFlag != nil {
					additional += " Use --write-config to create it."
				}
				log.Warnf("Config file not found.%s\n", additional)
			} else {
				return fmt.Errorf("could not read config: %w", err)
			}
		}

	}

	if parser.writeFlag && pflag.Lookup(writeFlagName()).Changed {
		log.Infoln("Writing configuration file.")

		if *writeFlag == "" {
			*writeFlag = "."
		}

		configpath := *writeFlag

		stat, err := os.Stat(*writeFlag)
		if err != nil {
			// If the specified path doesn't exist as a file or directory, attempt
			// to write as a file.
			if os.IsNotExist(err) {
				if strings.HasSuffix(configpath, "/") {
					err := os.MkdirAll(configpath, os.ModeDir)
					if err != nil {
						return fmt.Errorf("could not create directory %s provided via write flag: %w", configpath, err)
					}
					configpath = path.Join(configpath, parser.configName)
				} else {
					_, err := os.Create(configpath)
					if err != nil {
						return fmt.Errorf("could not create file %s provided via write flag: %w", configpath, err)
					}
				}
			} else {
				return fmt.Errorf("could not get stats for path %s provided via write flag: %w", configpath, err)
			}
		} else {
			if stat.IsDir() {
				configpath = path.Join(configpath, parser.configName)
			}
		}

		fmt.Println(configpath)
		if err := parser.viper.WriteConfigAs(configpath); err != nil {
			return fmt.Errorf("could not write viper config at path %s provided via write flag: %w", configpath, err)
		}
	}
	//parser.viper.WriteConfigAs("../test2/config.yml")

	return parser.viper.Unmarshal(&configStruct)
}
