// SPDX-License-Identifier: MIT
package semver_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSemver(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	suiteConfig, reporterConfig := GinkgoConfiguration()
	reporterConfig.Verbose = true
	RunSpecs(t, "Semver Suite", suiteConfig, reporterConfig)
}
