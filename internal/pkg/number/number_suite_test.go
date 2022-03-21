// SPDX-License-Identifier: MIT
package number_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNumber(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	suiteConfig, reporterConfig := GinkgoConfiguration()
	reporterConfig.Verbose = true
	RunSpecs(t, "Number Suite", suiteConfig, reporterConfig)
}
