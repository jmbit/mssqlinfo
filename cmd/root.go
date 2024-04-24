/*
Copyright © 2024 Johannes Bülow <johannes.buelow@jmbit.de>

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mssqlinfo",
	Short: "Get information about an MSSQL server",
	Long: `This application bundles common queries for information about MSSQL servers into a single cli application.
  Username, Password and Server address can either be entered via flags, environment variables or interactively`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/mssqlinfo/mssqlinfo.yaml)")
	rootCmd.PersistentFlags().StringP("user", "u", "", "user for authentication")
	rootCmd.PersistentFlags().StringP("password", "p", "", "password to use for authentication")
	rootCmd.PersistentFlags().StringP("server", "s", "", "address of the server")
	rootCmd.PersistentFlags().StringP("db", "d", "", "database to connect to")
	rootCmd.PersistentFlags().IntP("port", "P", 1433, "Port to connect to")
	viper.BindPFlags(rootCmd.PersistentFlags())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix("MSSQLINFO")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(path.Join(home, ".comfig/mssqlinfo"))
		viper.SetConfigType("yaml")
		viper.SetConfigName("mssqlinfo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
