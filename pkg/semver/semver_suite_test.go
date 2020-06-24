// SPDX-License-Identifier: MIT
package semver_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSemver(t *testing.T) {
	RegisterFailHandler(Fail)
	resultDir, ok := os.LookupEnv("TEST_RESULT_DIR")
	if !ok {
		resultDir = "../../artifacts/test-results"
	}
	if err := os.MkdirAll(resultDir, os.ModePerm); err != nil {
		panic(err)
	}
	junitReporter := reporters.NewJUnitReporter(filepath.Join(resultDir, "semver-tests-junit-report.xml"))
	config.DefaultReporterConfig.Verbose = true
	RunSpecsWithDefaultAndCustomReporters(t, "Semver Suite", []Reporter{junitReporter})
}
