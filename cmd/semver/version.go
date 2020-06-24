// SPDX-License-Identifier: MIT
package main

import (
	"github.com/spf13/cobra"

	"github.com/ffurrer2/semver/internal/pkg/app"
)

const versionFormat = `semver version %s
commit: %s
built at: %s
`

var versionCmd = &cobra.Command{
	Use:                   "version [flag]",
	Short:                 "Print version information",
	Long:                  `This command prints the version information.`,
	Example:               `  semver version`,
	Args:                  cobra.NoArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Printf(versionFormat, app.BuildVersion(), app.CommitHash(), app.BuildDate())
	},
}

func init() {
	versionCmd.SetUsageTemplate(usageTemplate)
	rootCmd.AddCommand(versionCmd)
}
