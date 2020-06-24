// SPDX-License-Identifier: MIT
package app

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CLI struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

func NewCLI() CLI {
	return CLI{
		stdin:  os.Stdin,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
}

func (c CLI) Stdin() io.Reader {
	return c.stdin
}

func (c CLI) Stdout() io.Writer {
	return c.stdout
}

func (c CLI) Stderr() io.Writer {
	return c.stderr
}

func (c CLI) Printf(format string, a ...interface{}) {
	_, err := fmt.Fprintf(c.stdout, format, a...)
	if err != nil {
		panic(err)
	}
}

func (c CLI) PrintErrf(format string, a ...interface{}) {
	_, err := fmt.Fprintf(c.stderr, format, a...)
	if err != nil {
		panic(err)
	}
}

func (c CLI) Apply(args []string, f func(s string)) {
	if len(args) > 0 {
		for _, s := range args {
			f(s)
		}
	} else {
		reader := bufio.NewReader(c.stdin)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			f(scanner.Text())
		}
	}
	os.Exit(0)
}

func (c CLI) Map(args []string, f func(s []string)) {
	if len(args) == 0 {
		args = make([]string, 0)
		reader := bufio.NewReader(c.stdin)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			args = append(args, scanner.Text())
		}
	}
	f(args)
	os.Exit(0)
}
