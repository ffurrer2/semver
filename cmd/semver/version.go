// SPDX-License-Identifier: MIT

package main

import (
	"github.com/ffurrer2/semver/v2/internal/pkg/app"
	"github.com/spf13/cobra"
)

const versionFormatShort = `%s`

const versionFormat = `semver version: %s
built at:       %s
git commit:     %s
git tree state: %s
`

var (
	short      bool
	versionCmd = &cobra.Command{
		Use:   "version [flag]",
		Short: "Print version information",
		Long:  `This command prints the version information.`,
		Example: `  semver version
  semver version --short`,
		Args:                  cobra.NoArgs,
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			if short {
				cmd.Printf(versionFormatShort, app.BuildVersion())
			} else {
				cmd.Printf(versionFormat, app.BuildVersion(), app.BuildDate(), app.CommitHash(), app.TreeState())
			}
		},
	}
)

func init() {
	versionCmd.SetUsageTemplate(usageTemplate)
	versionCmd.Flags().BoolVarP(&short, "short", "", false, "print only the version number")
	rootCmd.AddCommand(versionCmd)
}
