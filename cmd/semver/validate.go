// SPDX-License-Identifier: MIT

package main

import (
	"os"

	"github.com/ffurrer2/semver/v2/internal/pkg/cli"
	"github.com/ffurrer2/semver/v2/pkg/semver"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate [flag] [string]...",
	Short: "Validate semantic versions",
	Long: `This command checks if the given strings are valid semantic versions.

The exit code will be:
  0 if all string are valid semantic versions
  1 if at least one string is invalid`,
	Example: `  semver validate 1.0.0-alpha+001
  semver validate 1.0.0 1.0.0-alpha
  semver validate < file.txt
  echo '1.0.0-alpha\n1.0.0-alpha+001' | semver validate`,
	Args:                  cobra.ArbitraryArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		isValid := func(s string) {
			if !semver.IsValid(s) {
				os.Exit(1)
			}
		}
		cli.Apply(args, cmd.InOrStdin(), isValid)
	},
}

func init() {
	validateCmd.SetUsageTemplate(usageTemplate)
	rootCmd.AddCommand(validateCmd)
}
