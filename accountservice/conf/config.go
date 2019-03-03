/**
 * Create by Le Trong on 02/Mar/2019
 */

package conf

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ENVPREFIX Environment prefix
const ENVPREFIX = "MJ"

// DEFAULTCONFIGFILE Default config file name
const DEFAULTCONFIGFILE = "config"

// DEFAULTCONFIGPATH Default config path
const DEFAULTCONFIGPATH = "./"

// DEFAULTCONFIGHOME Default config home
const DEFAULTCONFIGHOME = "$HOME/.config"

// Config the application's config struct
type Config struct {
	Port      int64
	Config    string
	LogConfig LoggingConfig
}

// LoadConfig load configuration
func LoadConfig(cmd *cobra.Command) (*Config, error) {

	// From command line
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return nil, err
	}

	// From environment
	viper.SetEnvPrefix(ENVPREFIX)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// From config file
	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName(DEFAULTCONFIGFILE)
		viper.AddConfigPath(DEFAULTCONFIGPATH)
		viper.AddConfigPath(DEFAULTCONFIGHOME)
	}

	// NOTE: this will require that you have config file somewhere in the paths specified.
	// It can be reading from JSON, TOML, YAML, HCL, and Java properties files.
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return populateConfig(new(Config))
}
