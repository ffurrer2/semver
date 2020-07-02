// SPDX-License-Identifier: MIT
package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ffurrer2/semver/pkg/semver"
)

var sortCmd = &cobra.Command{
	Use:   "sort [flag] [semver]...",
	Short: "Sort semantic versions",
	Long:  `This command sorts semantic versions.`,
	Example: `  semver sort 1.1.1 1.0.0 1.0.1
  semver sort <versions.txt
  echo '1.1.1\n1.0.0\n1.0.1' | semver sort`,
	Args:                  cobra.ArbitraryArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		sort := func(args []string) {
			semverSlice := make([]semver.SemVer, 0)
			for _, s := range args {
				semVer, err := semver.Parse(s)
				if err != nil {
					cli.PrintErrf("error: %v\n", err)
					os.Exit(1)
				}
				semverSlice = append(semverSlice, *semVer)
			}
			semver.Sort(semverSlice)
			for _, s := range semverSlice {
				cli.Printf("%s\n", s)
			}
		}
		cli.Map(args, sort)
	},
}

func init() {
	sortCmd.SetUsageTemplate(usageTemplate)
	rootCmd.AddCommand(sortCmd)
}
