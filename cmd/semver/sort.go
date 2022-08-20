// SPDX-License-Identifier: MIT

package main

import (
	"os"
	"sort"

	"github.com/ffurrer2/semver/v2/internal/pkg/cli"
	"github.com/ffurrer2/semver/v2/pkg/semver"
	"github.com/spf13/cobra"
)

var (
	reverseFlag bool
	sortCmd     = &cobra.Command{
		Use:   "sort [flag] [semver]...",
		Short: "Sort semantic versions",
		Long:  `This command sorts semantic versions.`,
		Example: `  semver sort 1.1.1 1.0.0 1.0.1
  semver sort --reverse 1.1.1 1.0.0 1.0.1
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
						cmd.PrintErrf("error: %v\n", err)
						os.Exit(1)
					}
					semverSlice = append(semverSlice, *semVer)
				}
				if reverseFlag {
					sort.Sort(sort.Reverse(semver.BySemVer(semverSlice)))
				} else {
					sort.Sort(semver.BySemVer(semverSlice))
				}
				for _, s := range semverSlice {
					cmd.Printf("%s\n", s.String())
				}
			}
			cli.Map(args, cmd.InOrStdin(), sort)
		},
	}
)

func init() {
	sortCmd.SetUsageTemplate(usageTemplate)
	sortCmd.Flags().BoolVarP(&reverseFlag, "reverse", "r", false, "sort in reverse order")
	rootCmd.AddCommand(sortCmd)
}
