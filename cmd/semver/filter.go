// SPDX-License-Identifier: MIT

package main

import (
	"github.com/ffurrer2/semver/internal/pkg/cli"
	"github.com/ffurrer2/semver/pkg/semver"
	"github.com/spf13/cobra"
)

var (
	invalid   bool
	filterCmd = &cobra.Command{
		Use:   "filter [flag] [semver]...",
		Short: "Filter semantic versions",
		Long:  `This command prints either valid or invalid semantic versions.`,
		Example: `  semver filter 1.0.0 1.0 v1.0.0
  semver filter --invalid 1.0.0 1.0 v1`,
		Args:                  cobra.ArbitraryArgs,
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			filter := func(s string) {
				ok := semver.IsValid(s)
				if ok != invalid {
					cmd.Printf("%s\n", s)
				}
			}
			cli.Apply(args, cmd.InOrStdin(), filter)
		},
	}
)

func init() {
	filterCmd.SetUsageTemplate(usageTemplate)
	filterCmd.Flags().BoolVarP(&invalid, "invalid", "i", false, "print invalid semantic versions")
	rootCmd.AddCommand(filterCmd)
}
