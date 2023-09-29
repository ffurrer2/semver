// SPDX-License-Identifier: MIT

package main

import (
	"github.com/spf13/cobra"

	"github.com/ffurrer2/semver/v2/internal/pkg/cli"
	"github.com/ffurrer2/semver/v2/internal/pkg/predicate"
	"github.com/ffurrer2/semver/v2/pkg/semver"
)

var (
	printInvalidFlag       bool
	printReleasesFlag      bool
	printPreReleasesFlag   bool
	printBuildMetadataFlag bool
	filterCmd              = &cobra.Command{
		Use:   "filter [flag] [semver]...",
		Short: "Filter semantic versions",
		Long:  `This command filters (in)valid semantic versions, pre-release versions and versions containing build metadata.`,
		Example: `  semver filter 1.0.0 1.0 v1.0.0
  semver filter --invalid 1.0.0 1.0 v1,
  semver filter --pre-releases 1.0.0-alpha 1.0.0
  semver filter --releases=false 1.0.0-alpha 1.0.0,
  semver filter --build-metadata 1.0.0+exp.sha.5114f85 1.0.0`,
		Args:                  cobra.ArbitraryArgs,
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			filter := func(s string) {
				parsedSemVer, err := semver.Parse(s)
				if invalidFilter(err) || (validFilter(err) && predicate.Or(printReleasesFilter(), printPreReleasesFilter(), printBuildMetadataFilter())(parsedSemVer)) {
					cmd.Printf("%s\n", s)
				}
			}
			cli.Apply(args, cmd.InOrStdin(), filter)
		},
	}
)

func init() {
	filterCmd.SetUsageTemplate(usageTemplate)
	filterCmd.Flags().BoolVarP(&printInvalidFlag, "invalid", "i", false, "print only invalid semantic versions")
	filterCmd.Flags().BoolVarP(&printReleasesFlag, "releases", "", true, "print release versions")
	filterCmd.Flags().BoolVarP(&printPreReleasesFlag, "pre-releases", "", false, "print pre-release versions")
	filterCmd.Flags().BoolVarP(&printBuildMetadataFlag, "build-metadata", "", false, "print versions containing build metadata")
	filterCmd.MarkFlagsMutuallyExclusive("invalid", "releases")
	filterCmd.MarkFlagsMutuallyExclusive("invalid", "pre-releases")
	filterCmd.MarkFlagsMutuallyExclusive("invalid", "build-metadata")
	rootCmd.AddCommand(filterCmd)
}

func validFilter(err error) bool {
	return err == nil && !printInvalidFlag
}

func invalidFilter(err error) bool {
	return err != nil && printInvalidFlag
}

func printReleasesFilter() func(s *semver.SemVer) bool {
	return func(s *semver.SemVer) bool {
		return s.IsRelease() && printReleasesFlag
	}
}

func printPreReleasesFilter() func(s *semver.SemVer) bool {
	return func(s *semver.SemVer) bool {
		return s.IsPreRelease() && printPreReleasesFlag
	}
}

func printBuildMetadataFilter() func(s *semver.SemVer) bool {
	return func(s *semver.SemVer) bool {
		return s.HasBuildMetadata() && printBuildMetadataFlag
	}
}
