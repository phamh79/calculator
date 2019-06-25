package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	ErrorInvalidInput = errors.New("invalid input")
	ErrorUnknownCommand = errors.New("unknown command")
	ErrUsage = errors.New("usage")
)

func main() {
	m := NewMain()
	if err := m.Run(os.Args[1:]...); err == ErrUsage {
		os.Exit(2)
	} else if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}


type Main struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func NewMain() *Main {
	return &Main {
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}


// command args[0]
// x as args[1]
// x as args[2]
func (m *Main) Run(args ...string) error {
	// Require a command at the beginning.
	if len(args) == 0 {
		fmt.Fprintln(m.Stderr, m.Usage())
		return ErrUsage
	}
	return ErrorUnknownCommand
}

func (m *Main) Usage() string {
	return strings.TrimLeft(`
Calculator is a simple cli program to do mathematics calculations.
Usage:
	calculator command [arguments]
The commands are:
    add       add two int numbers
    subtract  first number is subtracted by second number
    multiply  multiply two int numbers
    divide    divide first number by second number time
`, "\n")
}
