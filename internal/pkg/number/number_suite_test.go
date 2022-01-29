// SPDX-License-Identifier: MIT
package number_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
)

func TestNumber(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	config.DefaultReporterConfig.Verbose = true
	RunSpecs(t, "Number Suite")
}
