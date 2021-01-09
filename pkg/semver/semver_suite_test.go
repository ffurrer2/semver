// SPDX-License-Identifier: MIT
package semver_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
)

func TestSemver(t *testing.T) {
	RegisterFailHandler(Fail)
	config.DefaultReporterConfig.Verbose = true
	RunSpecs(t, "Semver Suite")
}
