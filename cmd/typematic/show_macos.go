//go:build darwin

package main

import (
    "fmt"

    "github.com/akdlsz21/typematic/internal/system"
)

func printShow() error {
    st, err := system.Read()
    if err != nil {
        return err
    }
    // Summary line consistent with other platforms.
    fmt.Printf("delay_ms=%d interval_ms=%d rate_cps=%.2f\n", st.DelayMS, st.IntervalMS, st.RateCPS)

    // macOS specifics.
    det, err := system.ReadMacOSDetail()
    if err != nil {
        // If detailed read fails, still return OK after summary.
        return nil
    }
    fmt.Printf("defaults.keyrepeat_ticks=%d initial_keyrepeat_ticks=%d\n", det.KeyRepeatTicks, det.InitialKeyRepeatTicks)
    if det.PressAndHoldSet {
        fmt.Printf("apple.press_and_hold_enabled=%t\n", det.PressAndHoldEnabled)
    } else {
        fmt.Printf("apple.press_and_hold_enabled=<unset>\n")
    }
    return nil
}

