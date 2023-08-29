/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var debugEnabled bool
var infoEnabled bool
var traceEnabled bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gopt",
	Short: "Go OPTional Package Helper",
	Long: `This tool looks for software packages in $GOPT_PACKAGES, organized in name/semantic-version
folders. It can list the versions available and manipulate the PATH.`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if traceEnabled {
			logrus.SetLevel(logrus.TraceLevel)
		} else if debugEnabled {
			logrus.SetLevel(logrus.DebugLevel)
		} else if infoEnabled {
			logrus.SetLevel(logrus.InfoLevel)
		} else {
			logrus.SetLevel(logrus.ErrorLevel)
		}
	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gopt.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&debugEnabled, "debug", "", false, "Set log level to 'debug'")
	rootCmd.PersistentFlags().BoolVarP(&traceEnabled, "trace", "", false, "Set log level to 'trace'")
	rootCmd.PersistentFlags().BoolVarP(&infoEnabled, "info", "", false, "Set log level to 'info'")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
