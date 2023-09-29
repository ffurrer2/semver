// SPDX-License-Identifier: MIT

package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ffurrer2/semver/v2/pkg/semver"
)

const compareCmdExactArgs = 2

var compareCmd = &cobra.Command{
	Use:   "compare [flag] SEMVER1 SEMVER2",
	Short: "Compare semantic versions",
	Long: `This command compares two semantic versions.

The output will be:
 -1 if SEMVER1 < SEMVER2
  0 if SEMVER1 == SEMVER2
  1 if SEMVER1 > SEMVER2`,
	Example:               `  semver compare 1.0.0 1.0.0-alpha.1`,
	Args:                  cobra.ExactArgs(compareCmdExactArgs),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		semver1, err := semver.Parse(args[0])
		if err != nil {
			cmd.PrintErrf("error: %v\n", err)
			os.Exit(1)
		}
		semver2, err := semver.Parse(args[1])
		if err != nil {
			cmd.PrintErrf("error: %v\n", err)
			os.Exit(1)
		}
		result := semver1.CompareTo(*semver2)
		cmd.Printf("%d\n", result)
	},
}

func init() {
	compareCmd.SetUsageTemplate(usageTemplate)
	rootCmd.AddCommand(compareCmd)
}
