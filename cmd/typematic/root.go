package main

import (
	"fmt"

	"github.com/akdlsz21/typematic/internal/system"
	"github.com/akdlsz21/typematic/internal/units"
	"github.com/spf13/cobra"
)

var (
	delayMS int
	rateCPS float64
	show    bool
)

var rootCmd = &cobra.Command{
    Use:     "typematic",
    Short:   "Set keyboard auto-repeat delay and rate",
    Long:    longDesc,
    Example: `typematic --delay-ms=250 --rate-cps=25`,
    RunE: func(cmd *cobra.Command, args []string) error {
        // Show help when no flags are provided (idiomatic Cobra UX).
        if cmd.Flags().NFlag() == 0 {
            _ = cmd.Help()
			return nil
		}
		if show {
			return printShow()
		}
		if _, err := units.Validate(delayMS, rateCPS); err != nil {
			return fmt.Errorf("invalid input: %w", err)
		}
        if err := system.Set(delayMS, rateCPS); err != nil {
            return err
        }
        return nil
    },
	SilenceUsage:  false,
	SilenceErrors: false,
}

func init() {
	f := rootCmd.Flags()
	f.IntVarP(&delayMS, "delay-ms", "d", 0, "auto-repeat delay in ms [100..2000]")
	f.Float64VarP(&rateCPS, "rate-cps", "r", 0, "auto-repeat rate in characters per second (0.5..50.0]")
	f.BoolVar(&show, "show", false, "print current settings; on Windows includes FilterKeys and SPI details")
}

// Execute runs the Cobra command.
func Execute() error { return rootCmd.Execute() }
