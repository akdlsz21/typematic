//go:build darwin

package macos

import (
    "bytes"
    "fmt"
    "math"
    "os/exec"
    "strings"

    "github.com/akdlsz21/typematic/internal/units"
)

// Set applies the typematic settings via `defaults write -g` on macOS.
// KeyRepeat and InitialKeyRepeat use 15ms ticks; settings typically apply for new apps.
func Set(delayMS int, rateCPS float64) error {
    intervalMS, err := units.Validate(delayMS, rateCPS)
    if err != nil {
        return err
    }

    if _, err := lookPath("defaults"); err != nil {
        return fmt.Errorf("%w: defaults not found in PATH", ErrUnsupported)
    }

    // Convert to 15ms ticks using nearest rounding (closest possible to requested).
    // Clamp to reasonable bounds [1..120].
    toTicks := func(ms int) int {
        t := int(math.Round(float64(ms) / 15.0))
        if t < 1 {
            t = 1
        }
        if t > 120 {
            t = 120
        }
        return t
    }
    kr := toTicks(intervalMS)  // repeat interval ticks (15ms each)
    ikr := toTicks(delayMS)    // delay ticks (15ms each)

    // Write only to the global domain (-g).
    if err := defaultsWrite("-g", "KeyRepeat", kr); err != nil {
        return err
    }
    if err := defaultsWrite("-g", "InitialKeyRepeat", ikr); err != nil {
        return err
    }

    // Report the effective values applied after macOS clamping/granularity.
    effInterval := kr * 15
    effDelay := ikr * 15
    effCPS := 0.0
    if effInterval > 0 {
        effCPS = 1000.0 / float64(effInterval)
    }
    // Concise report: show only the effective, clamped values applied.
    fmt.Printf("applied (macos): delay_ms=%d interval_ms=%d rate_cps=%.2f\n", effDelay, effInterval, effCPS)
    return nil
}

func defaultsWrite(domainFlag, key string, value int) error {
    var stdout, stderr bytes.Buffer
    cmd := exec.Command("defaults", "write", domainFlag, key, "-int", fmt.Sprintf("%d", value))
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("%w: %s", ErrOSCall, strings.TrimSpace(stderr.String()))
    }
    return nil
}
