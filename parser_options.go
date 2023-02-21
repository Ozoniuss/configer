package config

// ParserOption is an interface which allows implementing the functional
// options pattern for the config parser. By default, the package comes with
// some default options, and offers the possibility to overwrite them in
// every project using this package by using the WithXXX public functions
// exported below.
type ParserOption interface {
	// apply changes the internal parser options.
	apply(*configParser)
}
