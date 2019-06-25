package main

import (
	calculator "cmd_project"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
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

	// Execute command.
	switch args[0] {
	case "help":
		fmt.Fprintln(m.Stderr, m.Usage())
		return ErrUsage
	case "add":
		return newAddCommand(m).Run(args...)
	case "subtract":
		return nil
	case "divide":
		return nil
	case "multiply":
		return nil
	default:
		return ErrorUnknownCommand

	}
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

// AddCommand represents the "add" command execution.
type AddCommand struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// NewAddCommand returns a AddCommand.
func newAddCommand(m *Main) *AddCommand {
	return &AddCommand{
		Stdin:  m.Stdin,
		Stdout: m.Stdout,
		Stderr: m.Stderr,
	}
}

// Run executes the command.
func (cmd *AddCommand) Run(args ...string) error {
	// Parse flags.
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	help := fs.Bool("h", false, "")
	if err := fs.Parse(args); err != nil {
		return err
	} else if *help {
		fmt.Fprintln(cmd.Stderr, cmd.Usage())
		return ErrUsage
	}

	x, err := strconv.Atoi(args[1])
	if err != nil {
		return ErrorInvalidInput
	}

	y, err := strconv.Atoi(args[2])
	if err != nil {
		return ErrorInvalidInput
	}

	fmt.Fprintf(cmd.Stdout, "(%d) adds (%d) equals %d\n", x, y, calculator.Add(x, y))
	return nil
}

// Usage returns the help message for AddCommand
func (cmd *AddCommand) Usage() string {
	return strings.TrimLeft(`
usage: add two numbers
return numbers of two integers
`, "\n")
}

