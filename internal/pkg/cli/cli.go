// SPDX-License-Identifier: MIT

package cli

import (
	"bufio"
	"io"
	"os"
)

func Apply(args []string, rd io.Reader, f func(s string)) {
	if len(args) > 0 {
		for _, s := range args {
			f(s)
		}
	} else {
		reader := bufio.NewReader(rd)

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			f(scanner.Text())
		}
	}

	os.Exit(0)
}

func Map(args []string, rd io.Reader, f func(s []string)) {
	if len(args) == 0 {
		args = make([]string, 0)
		reader := bufio.NewReader(rd)

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			args = append(args, scanner.Text())
		}
	}

	f(args)
	os.Exit(0)
}
