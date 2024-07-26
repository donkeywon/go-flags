package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/donkeywon/go-flags"
)

type EditorOptions struct {
	Input  flags.Filename `flag-short:"i" flag-long:"input" flag-description:"Input file" flag-default:"-"`
	Output flags.Filename `flag-short:"o" flag-long:"output" flag-description:"Output file" flag-default:"-"`
}

type Point struct {
	X, Y int
}

func (p *Point) UnmarshalFlag(value string) error {
	parts := strings.Split(value, ",")

	if len(parts) != 2 {
		return errors.New("expected two numbers separated by a ,")
	}

	x, err := strconv.ParseInt(parts[0], 10, 32)

	if err != nil {
		return err
	}

	y, err := strconv.ParseInt(parts[1], 10, 32)

	if err != nil {
		return err
	}

	p.X = int(x)
	p.Y = int(y)

	return nil
}

func (p Point) MarshalFlag() (string, error) {
	return fmt.Sprintf("%d,%d", p.X, p.Y), nil
}

type Options struct {
	// Example of verbosity with level
	Verbose []bool `flag-short:"v" flag-long:"verbose" flag-description:"Verbose output"`

	// Example of optional value
	User string `flag-short:"u" flag-long:"user" flag-description:"User name" flag-optional:"yes" flag-optional-value:"pancake"`

	// Example of map with multiple default values
	Users map[string]string `flag-long:"users" flag-description:"User e-mail map" flag-default:"system:system@example.org" flag-default:"admin:admin@example.org"`

	// Example of option group
	Editor EditorOptions `flag-group:"Editor Options"`

	// Example of custom type Marshal/Unmarshal
	Point Point `flag-long:"point" flag-description:"A x,y point" flag-default:"1,2"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default, flags.FlagTagPrefix("flag-"))

func main() {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}
