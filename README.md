# configer

I wrote this package to have a quick and easy way to configure my personal projects. It uses [viper](https://github.com/spf13/viper) as the configuration package under the hood, and the [pflag](https://github.com/spf13/pflag) library to work with flags.

Installation
------------

To get the latest version, run

```
go get github.com/Ozoniuss/configer
```

Usage
-----

Start by defining a structure for storing the project's overall configuration. The runtime configurations will be unmarshaled into an object of this struct when the project is started, which is passed on via dependency injection.

```go
// Config holds the project's runtime configuration.
type Config struct {
	Server   Server
	Key      int
	Insecure bool
}

// Server defines the configuration options for the server.
type Server struct {
	Address string
	Port    int32
}
```

Use the configer `ConfigOptions` struct to specify the project's configuration options. Through the `ConfigKey` field, specify the field of the config struct the option is attached to:

```go
// getProjectOpts returns all the configuration options enabled for the project.
func getProjectOpts() []cfg.ConfigOption {
	return []cfg.ConfigOption{
		{FlagName: "server-address", Shorthand: "", Value: "Mirel", ConfigKey: "server.address",
			Usage: "The address on which the server is listening for connections."},
		{FlagName: "server-port", Shorthand: "", Value: 8080, ConfigKey: "server.port",
			Usage: "The server port opened for incoming connections."},
		{FlagName: "key", Shorthand: "k", Value: 123456, ConfigKey: "key",
			Usage: "The key required to access the server."},
		{FlagName: "insecure", Shorthand: "", Value: true, ConfigKey: "insecure",
			Usage: "Specify whether or not TLS is enabled."},
	}
}
```

Create the config:

```go
import (
	"fmt"
	cfg "github.com/Ozoniuss/configer"
)

func main() {
	config := NewConfig()

	parserOptions := []cfg.ParserOption{
		cfg.WithConfigName("config"),
		cfg.WithConfigType("yml"),
		cfg.WithConfigPath("."),
		cfg.WithConfigPath("./config"),
		cfg.WithEnvPrefix("DEMO"),
		cfg.WithEnvKeyReplacer("_"),
	}

	err := cfg.NewConfig(&config, getProjectOpts(), parserOptions...)
	if err != nil {
		fmt.Println("could not initialize config: %w", err)
		return
	}
	
	fmt.Printf("config: %+v", config)
}
```

Output:

```
config: {Server:{Address:localhost Port:8080} Key:123456 Insecure:true}
```

See the available config options:

```
$ ./main --help
Usage of main:
      --insecure                Specify whether or not TLS is enabled. (default true)
  -k, --key int                 The key required to access the server. (default 123456)
      --server-address string   The address on which the server is listening for connections. (default "localhost")
      --server-port int         The server port opened for incoming connections. (default 8080)
```

And you're all set. Specify config options via flags, environment variables, configuration files and default values, in this order of precedence.

```
$ DEMO_PORT=9090
$ ./main --key 100 --server-address 0.0.0.0
{Server:{Address:0.0.0.0 Port:9090} Key:100 Insecure:true}
```

See the [documentation](DOCUMENTATION.md) for more details and advanced usage.

Personal notes
--------------

When working with configs, I often found myself repeating the same patterns for coding a basic configuration setup on my projects. For this reason, I abstracted away that pattern to a custom package that I find useful for most of my projects.

However, this works slightly different that how viper would have worked by default using a similar setup. I find that in some aspects what viper is doing is not intuitive, at least for me, and for this reason the usage of this package is slightly different than viper's usage. For example:

- Files set via `viper.SetConfigFile` throw a different error if not found than if set via `viper.AddConfigPath` combined with `viper.SetConfigName` and `viper.SetConfigType`; 
- Writing files using `viper.WriteXXX` and `viper.SafeWriteXXX` is a challenge. Additionally, these same functions behave differently if the config file was set with one of the two methods above;
- The behaviour of the config files is just so different between the two approaches listed above it's just unbelieveable. I often get lost working with config files with viper;
- `viper.SetConfigFile` doesn't seem to actually overwrite what was set via `viper.AddConfigPath` as mentioned in the docs.

The comments I currently left to the viper repository can be found on my [open source contribution](https://github.com/Ozoniuss) list. I will likely make some more contributions in the future.