//go:build darwin

package macos

import (
    "bytes"
    "fmt"
    "os/exec"
    "strconv"
    "strings"
)

// State represents current typematic settings on macOS.
// macOS stores InitialKeyRepeat and KeyRepeat as integer ticks (15ms units).
type State struct {
    DelayMS    int
    IntervalMS int
    RateCPS    float64
}

// Read queries defaults for InitialKeyRepeat (delay) and KeyRepeat (interval in 15ms units).
func Read() (State, error) {
    if _, err := lookPath("defaults"); err != nil {
        return State{}, fmt.Errorf("%w: defaults not found in PATH", ErrUnsupported)
    }

    // KeyRepeat: repeat interval in 15ms units (lower is faster)
    kr, err := readIntDefault("-g", "KeyRepeat")
    if err != nil {
        return State{}, fmt.Errorf("%w: %v", ErrOSCall, err)
    }
    // InitialKeyRepeat: delay until repeat starts in 15ms units
    ikr, err := readIntDefault("-g", "InitialKeyRepeat")
    if err != nil {
        return State{}, fmt.Errorf("%w: %v", ErrOSCall, err)
    }

    interval := kr * 15
    delay := ikr * 15
    cps := 0.0
    if interval > 0 {
        cps = 1000.0 / float64(interval)
    }
    return State{DelayMS: delay, IntervalMS: interval, RateCPS: cps}, nil
}

func readIntDefault(domainFlag, key string) (int, error) {
    var stdout, stderr bytes.Buffer
    cmd := exec.Command("defaults", "read", domainFlag, key)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    if err := cmd.Run(); err != nil {
        s := strings.TrimSpace(stderr.String())
        if s != "" {
            return 0, fmt.Errorf("defaults read error: %s", s)
        }
        return 0, err
    }
    s := strings.TrimSpace(stdout.String())
    n, err := strconv.Atoi(s)
    if err != nil {
        return 0, fmt.Errorf("unexpected output: %q", s)
    }
    return n, nil
}
