// SPDX-License-Identifier: MIT

package main

import (
	"os"

	"github.com/ffurrer2/semver/v2/internal/pkg/cli"
	"github.com/ffurrer2/semver/v2/pkg/semver"
	"github.com/spf13/cobra"
)

var majorCmd = &cobra.Command{
	Use:                   "major [flag] [semver]...",
	Short:                 "Increment semantic versions to the next major version",
	Long:                  `This command increments a given semantic version to the next major version.`,
	Example:               `  semver next major 1.0.0-alpha+001`,
	Args:                  cobra.ArbitraryArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		nextMajor := func(s string) {
			semver, err := semver.Parse(s)
			if err != nil {
				cmd.PrintErrf("error: %v\n", err)
				os.Exit(1)
			}
			cmd.Printf("%s\n", semver.NextMajor().String())
		}
		cli.Apply(args, cmd.InOrStdin(), nextMajor)
	},
}

func init() {
	majorCmd.SetUsageTemplate(usageTemplate)
	nextCmd.AddCommand(majorCmd)
}
