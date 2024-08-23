package cmd

import (
	"fmt"
	"log-seeker/cmd/analyzer"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "log-seeker",
	Short: "log-seeker",
	Long:  `Log seeker cli app`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			os.Exit(1)
		}
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(analyzer.AnalyzeLogFormFileCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("log-seeker")
	viper.AutomaticEnv() // read in environment variables that match
}
