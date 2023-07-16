// SPDX-License-Identifier: MIT

package main

import (
	"github.com/ffurrer2/semver/v2/internal/pkg/app"
	"github.com/spf13/cobra"
)

const versionFormatShort = `%s`

const versionFormat = `semver version: %s
git commit:     %s
git tree state: %s
`

var (
	json       bool
	short      bool
	versionCmd = &cobra.Command{
		Use:   "version [flag]",
		Short: "Print version information",
		Long:  `This command prints the version information.`,
		Example: `  semver version
  semver version --json
  semver version --short`,
		Args:                  cobra.NoArgs,
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			if json {
				cmd.Println(app.BuildInfoJSON())
			} else if short {
				cmd.Printf(versionFormatShort, app.Version())
			} else {
				cmd.Printf(versionFormat, app.Version(), app.CommitHash(), app.TreeState())
			}
		},
	}
)

func init() {
	versionCmd.SetUsageTemplate(usageTemplate)
	versionCmd.Flags().BoolVarP(&short, "short", "", false, "print only the version number")
	versionCmd.Flags().BoolVarP(&json, "json", "", false, "print complete build info as json")
	versionCmd.MarkFlagsMutuallyExclusive("short", "json")
	rootCmd.AddCommand(versionCmd)
}
