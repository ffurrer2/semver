// SPDX-License-Identifier: MIT
package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ffurrer2/semver/pkg/semver"
)

var compareCmd = &cobra.Command{
	Use:   "compare [flag] semver1 semver2",
	Short: "Compare semantic versions",
	Long: `This command compares two semantic versions.

The output will be:
 -1 if semver1 < semver2
  0 if semver1 == semver2
  1 if semver1 > semver2`,
	Example:               `  semver compare 1.0.0 1.0.0-alpha.1`,
	Args:                  cobra.ExactArgs(2),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		semver1, err := semver.Parse(args[0])
		if err != nil {
			cli.PrintErrf("error: %v\n", err)
			os.Exit(1)
		}
		semver2, err := semver.Parse(args[1])
		if err != nil {
			cli.PrintErrf("error: %v\n", err)
			os.Exit(1)
		}
		result := semver1.CompareTo(*semver2)
		cli.Printf("%d\n", result)
	},
}

func init() {
	compareCmd.SetUsageTemplate(usageTemplate)
	rootCmd.AddCommand(compareCmd)
}
