//go:build windows

package main

import (
    "fmt"
    "strings"

    "github.com/akdlsz21/typematic/internal/system"
    "github.com/zzl/go-win32api/v2/win32"
)

func printShow() error {
    st, err := system.Read()
    if err != nil {
        return err
    }
    det, err := system.ReadWindowsDetail()
    if err != nil {
        // If detailed read fails, at least print the generic triple.
        fmt.Printf("delay_ms=%d interval_ms=%d rate_cps=%.2f\n", st.DelayMS, st.IntervalMS, st.RateCPS)
        return nil
    }

    // First line: same summary as other platforms.
    fmt.Printf("delay_ms=%d interval_ms=%d rate_cps=%.2f\n", st.DelayMS, st.IntervalMS, st.RateCPS)

    // Windows details.
    fmt.Printf("filterkeys.enabled=%t flags=0x%X [%s]\n", det.Enabled, det.Flags, decodeFKFlags(det.Flags))
    fmt.Printf("filterkeys.wait_ms=%d delay_ms=%d repeat_ms=%d bounce_ms=%d\n", det.WaitMSec, det.DelayMSec, det.RepeatMSec, det.BounceMSec)
    fmt.Printf("spi.keyboard.speed=%d delay=%d\n", det.KeyboardSpeed, det.KeyboardDelay)
    return nil
}

func decodeFKFlags(flags uint32) string {
    var names []string
    if flags&win32.FKF_FILTERKEYSON != 0 {
        names = append(names, "FILTERKEYSON")
    }
    if flags&win32.FKF_AVAILABLE != 0 {
        names = append(names, "AVAILABLE")
    }
    if flags&win32.FKF_HOTKEYACTIVE != 0 {
        names = append(names, "HOTKEYACTIVE")
    }
    if flags&win32.FKF_CONFIRMHOTKEY != 0 {
        names = append(names, "CONFIRMHOTKEY")
    }
    if flags&win32.FKF_HOTKEYSOUND != 0 {
        names = append(names, "HOTKEYSOUND")
    }
    if flags&win32.FKF_INDICATOR != 0 {
        names = append(names, "INDICATOR")
    }
    if flags&win32.FKF_CLICKON != 0 {
        names = append(names, "CLICKON")
    }
    if len(names) == 0 {
        return ""
    }
    return strings.Join(names, ",")
}

