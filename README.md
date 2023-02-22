work in progress...

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

To be continued.

Personal notes
--------------

When working with configs, I often found myself repeating the same patterns for coding a basic configuration setup on my projects. For this reason, I abstracted away that pattern to a custom package that I find useful for most of my projects.

However, this works slightly different that how viper would have worked by default using a similar setup. I find that in some aspects what viper is doing is not intuitive, at least for me, and for this reason the usage of this package is slightly different than viper's usage. For example:

- Files set via `viper.SetConfigFile` throw a different error if not found than if set via `viper.AddConfigPath` combined with `viper.SetConfigName` and `viper.SetConfigType`; 
- Writing files using `viper.WriteXXX` and `viper.SafeWriteXXX` is a challenge. Additionally, these same functions behave differently if the config file was set with one of the two methods above;
- The behaviour of the config files is just so different between the two approaches listed above it's just unbelieveable. I often get lost working with config files with viper;
- `viper.SetConfigFile` doesn't seem to actually overwrite what was set via `viper.AddConfigPath` as mentioned in the docs.

The comments I currently left to the viper repository can be found on my [open source contribution](https://github.com/Ozoniuss) list. I will likely make some more contributions in the future.