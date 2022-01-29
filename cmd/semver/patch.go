// SPDX-License-Identifier: MIT

package main

import (
	"os"

	"github.com/ffurrer2/semver/internal/pkg/cli"
	"github.com/ffurrer2/semver/pkg/semver"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:                   "patch [flag] [semver]...",
	Short:                 "Increment semantic versions to the next patch version",
	Long:                  `This command increments a given semantic version to the next patch version.'`,
	Example:               `  semver next patch 1.0.0-alpha+001`,
	Args:                  cobra.ArbitraryArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		nextPatch := func(s string) {
			semver, err := semver.Parse(s)
			if err != nil {
				cmd.PrintErrf("error: %v\n", err)
				os.Exit(1)
			}
			cmd.Printf("%s\n", semver.NextPatch().String())
		}
		cli.Apply(args, cmd.InOrStdin(), nextPatch)
	},
}

func init() {
	patchCmd.SetUsageTemplate(usageTemplate)
	nextCmd.AddCommand(patchCmd)
}
