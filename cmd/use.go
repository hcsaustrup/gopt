/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hcsaustrup/gopt/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var usePrepend *bool

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Build and output PATH statement that includes the selected versions",
	Long: `Build and output PATH statement suitable for evaluation by a BASH
compatible shell.

Specify as many packages to include in the format "package:version". Omitting ":version"
will get you the latest version. Omitting parameters will get you all latest versions of
all found packages.`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := app.Config.GetRepository()
		if err != nil {
			logrus.Fatal(err)
		}

		path := os.Getenv("PATH")

		if len(args) > 0 {
			for _, arg := range args {
				v, err := repo.FindPackageVersion(arg)
				if err != nil {
					logrus.Warn(err)
					continue
				}
				path = v.AddToPath(path, *usePrepend)
			}
		} else {
			for _, p := range repo.Packages {
				path = p.Versions.Latest().AddToPath(path, *usePrepend)
			}
		}

		fmt.Printf("export PATH=\"%s\"\n", path)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	usePrepend = useCmd.Flags().BoolP("prepend", "p", false, "Insert in beginning of path instead of at the end")
}
