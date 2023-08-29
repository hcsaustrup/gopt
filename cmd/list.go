/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hcsaustrup/gopt/app"
	"github.com/jedib0t/go-pretty/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var listAsTable *bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all found packages",
	Long: `List all found packages in $GOPT_PACKAGES. Specify one of more packages
to list available versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := app.Config.GetRepository()
		if err != nil {
			logrus.Fatal(err)
		}

		asTable := *listAsTable == true
		var t table.Writer = nil

		if asTable {
			t = table.NewWriter()
			t.SetOutputMirror(os.Stdout)
		}

		if len(args) == 0 {
			if asTable {
				t.AppendHeader(table.Row{"Package", "Path", "Latest Version"})
			}

			for _, p := range repo.Packages {
				if asTable {
					t.AppendRow(table.Row{p.Name, p.Path, p.Versions.Latest().Version})
				} else {
					fmt.Printf("%s (%s)\n", p.Name, p.Path)
				}
			}
		} else {
			if asTable {
				t.AppendHeader(table.Row{"Package", "Version", "Path"})
			}

			for _, arg := range args {
				p, ok := repo.PackageMap[arg]
				if !ok {
					continue
				}
				for _, v := range p.Versions {

					if asTable {
						t.AppendRow(table.Row{p.Name, v.Version, v.Path})
					} else {
						fmt.Printf("%s %s (%s)\n", p.Name, v.Version, v.Path)
					}

				}
			}
		}

		if asTable {
			t.Render()
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	listAsTable = listCmd.PersistentFlags().BoolP("table", "t", false, "Output in pretty table")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
