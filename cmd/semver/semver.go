// SPDX-License-Identifier: MIT
package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ffurrer2/semver/internal/pkg/app"
)

const usageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

var cli app.CLI

var rootCmd = &cobra.Command{
	Use:  "semver",
	Long: "The semantic version utility",
}

func init() {
	cli = app.NewCLI()
	rootCmd.SetUsageTemplate(usageTemplate)
	rootCmd.Flags().FlagUsages()
}

func main() {
	execute()
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
