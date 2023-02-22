package configer

type readFlagOption bool

func (opt readFlagOption) apply(parser *configParser) {
	parser.readFlag = bool(opt)
}

// WithReadFlag defines a flag that can be used to explicitly set the location
// from where the configuration file is read. It ignores the options that
// were set through WithConfigPath.
//
// If the specified location is a directory, the parser will search for a
// config file with the name and type specified via WithConfigName and
// WithConfigType, or the default values if those were not set. Otherwise,
// if the location is a file, it will try to read the config from that file.
// If no config file is found, the parser will throw an error.
//
// By default, this flag will not be defined.
func WithReadFlag() readFlagOption {
	return readFlagOption(true)
}
