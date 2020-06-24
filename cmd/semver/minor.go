// SPDX-License-Identifier: MIT
package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ffurrer2/semver/pkg/semver"
)

var minorCmd = &cobra.Command{
	Use:                   "minor [flag] [semver]...",
	Short:                 "Increment semantic versions to the next minor version",
	Long:                  `This command increments a given semantic version to the next minor version.`,
	Example:               `  semver next minor 1.0.0-alpha+001`,
	Args:                  cobra.ArbitraryArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		nextMinor := func(s string) {
			semver, err := semver.Parse(s)
			if err != nil {
				cli.PrintErrf("error: %v\n", err)
				os.Exit(1)
			}
			cli.Printf("%s\n", semver.NextMinor().String())
		}
		cli.Apply(args, nextMinor)
	},
}

func init() {
	minorCmd.SetUsageTemplate(usageTemplate)
	nextCmd.AddCommand(minorCmd)
}
