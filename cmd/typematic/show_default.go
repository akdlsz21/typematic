//go:build !windows && !darwin

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
    fmt.Printf("delay_ms=%d interval_ms=%d rate_cps=%.2f\n", st.DelayMS, st.IntervalMS, st.RateCPS)
    return nil
}
