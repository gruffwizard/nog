/*
Copyright © 2020 Steve Poole  gruff.wizard@yahoo.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var mvnLoc  string
var srcLoc  string
var Verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nog",
	Short: "Containerised Quarkus development environment",
	Long: `
Nog is an interface between local code and a containerised quarkus
development environment. Nog mounts local maven repos, source code etc
into the container running quarkus and allows you to develop in your
IDE and have Quarkus (in a container ) build and run it.

As an added bonus Nog also includes a browser based IDE to allow
for the full quarkus experience without having to install anything.

Well except for Nog - and you've already done that.
	`,

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nog.yaml)")
	rootCmd.PersistentFlags().StringVarP(&mvnLoc, "maven","m", "", "maven repo location")
	rootCmd.PersistentFlags().StringVarP(&srcLoc, "src", "s","", "src location")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v",false, "verbose")

	}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".nog" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".nog")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
