package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/akdlsz21/typematic/internal/system"
)

const (
	exitOK          = 0
	exitInvalid     = 2
	exitUnsupported = 3
	exitOSCall      = 4
)

func main() {
	if err := Execute(); err != nil {
		switch {
		case errors.Is(err, system.ErrUnsupported):
			fmt.Fprintln(os.Stderr, "unsupported environment:", err)
			os.Exit(exitUnsupported)
		case errors.Is(err, system.ErrOSCall):
			fmt.Fprintln(os.Stderr, "os call failed:", err)
			os.Exit(exitOSCall)
		default:
			// Treat CLI/validation errors as invalid input (exit 2)
			fmt.Fprintln(os.Stderr, err)
			os.Exit(exitInvalid)
		}
	}
}
