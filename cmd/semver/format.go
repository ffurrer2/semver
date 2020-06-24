// SPDX-License-Identifier: MIT
package main

import (
	"encoding/json"
	"os"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/ffurrer2/semver/pkg/semver"
)

type semVer struct {
	Major         uint   `json:"Major"`
	Minor         uint   `json:"Minor"`
	Patch         uint   `json:"Patch"`
	PreRelease    string `json:"PreRelease"`
	BuildMetadata string `json:"BuildMetadata"`
}

var formatCmd = &cobra.Command{
	Use:   "format [flag] format [semver]",
	Short: "Format and print semantic versions",
	Long: `This command formats and prints the given semantic versions according to the given format template.

The struct being passed to the template is:

type SemVer struct {
    Major         uint
    Minor         uint
    Patch         uint
    PreRelease    string
    BuildMetadata string
}`,
	Example: `  semver format '{{.Major}}.{{.Minor}}' 1.0.0
  semver format '{{.Major}}.{{.Minor}}.{{.Patch}}-{{.PreRelease}}.1' 1.0.0-alpha+001`,
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		text := args[0]
		tpl, err := template.New("semver").Parse(text)
		if err != nil {
			cli.PrintErrf("error: %v\n", err)
		}
		format := func(s string) {
			semver, err := semver.Parse(s)
			if err != nil {
				cli.PrintErrf("error: %v\n", err)
				os.Exit(1)
			}
			data := semVer{
				Major:         semver.Major,
				Minor:         semver.Minor,
				Patch:         semver.Patch,
				PreRelease:    semver.PreReleaseString(),
				BuildMetadata: semver.BuildMetadataString(),
			}
			err = tpl.Execute(os.Stdout, data)
			if err != nil {
				cli.PrintErrf("error: %v\n", err)
			}
			cli.Printf("\n")
		}
		cli.Apply(args[1:], format)
	},
}

func init() {
	formatCmd.SetUsageTemplate(usageTemplate)
	rootCmd.AddCommand(formatCmd)
}

func (s semVer) String() string {
	jsonStr, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(jsonStr)
}
