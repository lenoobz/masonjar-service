/**
 * Create by Le Trong on 02/Mar/2019
 */

package cmd

import (
	"log"

	"github.com/letrong/masonjar-service/accountservice/conf"
	"github.com/letrong/masonjar-service/accountservice/service"
	"github.com/letrong/masonjar-service/utilities"
	"github.com/spf13/cobra"
)

// RootCommand set up and return root command
func RootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "app",
		Run: run,
	}

	rootCmd.PersistentFlags().StringP("config", "c", "", "An explicit config file to use")
	rootCmd.Flags().IntP("port", "p", 0, "Web server port")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	config, err := conf.LoadConfig(cmd)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}

	logger, err := conf.ConfigureLogging(&config.LogConfig)
	if err != nil {
		log.Fatal("Failed to configure logging: " + err.Error())
	}

	prettyConfig, err := utilities.PrettyPrint(config)
	if err != nil {
		log.Fatal("Failed to parse config: " + err.Error())
	}

	logger.Infof("Starting with config: \n%s", prettyConfig)
	service.StartWebServer(config, logger)
}
