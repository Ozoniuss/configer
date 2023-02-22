package configer

type writeFlagOption bool

func (opt writeFlagOption) apply(parser *configParser) {
	parser.writeFlag = bool(opt)
}

// WithWriteFlag defines a flag that can be used to write the project
// configuration to a file.
//
// If the specified location is a directory, the parser will search for a
// config file with the name and type specified via WithConfigName and
// WithConfigType, or the default values if those were not set. Otherwise,
// if the location is a file, it will try to read the config from that file.
//
// If there already is a configuration file at the specified location, this
// option will overwrite that file.
//
// By default, this flag will not be defined.
func WithWriteFlag() writeFlagOption {
	return writeFlagOption(true)
}
