// SPDX-License-Identifier: MIT

package main

import (
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Increment semantic versions",
	Long:  `This command increments a given semantic version to the next major, minor or patch version.`,
	Args:  cobra.NoArgs,
}

func init() {
	nextCmd.SetUsageTemplate(usageTemplate)
	rootCmd.AddCommand(nextCmd)
}
