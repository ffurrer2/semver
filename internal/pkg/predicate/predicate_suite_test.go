// SPDX-License-Identifier: MIT

package predicate_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestPredicate(t *testing.T) {
	t.Parallel()
	gomega.RegisterFailHandler(Fail)
	suiteConfig, reporterConfig := GinkgoConfiguration()
	reporterConfig.Verbose = true
	RunSpecs(t, "Predicate Suite", suiteConfig, reporterConfig)
}
