package configer

// ConfigOption defines the values for each coniguration option of a
// microservice.
type ConfigOption struct {
	// The name of the flag associated with the configuration option. May be
	// empty if the option is associated with no flag.
	FlagName string
	// The shorthand of the flag. Must be empty if a flagname is not specified.
	Shorthand string
	// The value of the configuration option.
	Value any
	// The description of the configuration option. May be displayed in the
	// command line using the --help flag.
	Usage string
	// The key associated with the configuration option. If desired to associate
	// the key with the field of a struct, should be named the same as the
	// struct field (case insensitive). If the field is also a struct, the "."
	// separator should be used for the config key.
	//
	// E.g. if binding to the struct below,
	//
	// =========================
	// | type Config struct {  |
	// | 	Server Server      |
	// | }                     |
	// |                       |
	// | type Server struct {  |
	// | 	Address int        |
	// | }                     |
	// =========================
	//
	// the value of the config option with key "server.address" would get
	// unmarshaled to the server's address in the struct.
	ConfigKey string
}
