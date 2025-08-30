//go:build darwin

package macos

import (
    "bytes"
    "fmt"
    "os/exec"
    "strconv"
    "strings"
)

// Detail exposes macOS-specific settings behind --show.
// KeyRepeat/InitialKeyRepeat are stored in 15ms ticks. ApplePressAndHoldEnabled
// may be unset; PressAndHoldSet indicates presence.
type Detail struct {
    KeyRepeatTicks          int
    InitialKeyRepeatTicks   int
    PressAndHoldEnabled     bool
    PressAndHoldSet         bool
}

// ReadDetail returns macOS-specific typematic details.
func ReadDetail() (Detail, error) {
    if _, err := lookPath("defaults"); err != nil {
        return Detail{}, fmt.Errorf("%w: defaults not found in PATH", ErrUnsupported)
    }
    kr, err := readIntDefault("-g", "KeyRepeat")
    if err != nil {
        return Detail{}, fmt.Errorf("%w: %v", ErrOSCall, err)
    }
    ikr, err := readIntDefault("-g", "InitialKeyRepeat")
    if err != nil {
        return Detail{}, fmt.Errorf("%w: %v", ErrOSCall, err)
    }
    val, set, err := readOptionalBoolDefault("-g", "ApplePressAndHoldEnabled")
    if err != nil {
        return Detail{}, fmt.Errorf("%w: %v", ErrOSCall, err)
    }
    return Detail{KeyRepeatTicks: kr, InitialKeyRepeatTicks: ikr, PressAndHoldEnabled: val, PressAndHoldSet: set}, nil
}

// readOptionalBoolDefault reads a boolean defaults key and reports whether it was set.
func readOptionalBoolDefault(domainFlag, key string) (value bool, set bool, err error) {
    var stdout, stderr bytes.Buffer
    cmd := exec.Command("defaults", "read", domainFlag, key)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    if e := cmd.Run(); e != nil {
        serr := strings.TrimSpace(stderr.String())
        // If key does not exist, signal unset without error.
        if strings.Contains(serr, "does not exist") {
            return false, false, nil
        }
        if serr != "" {
            return false, false, fmt.Errorf("defaults read error: %s", serr)
        }
        return false, false, e
    }
    s := strings.TrimSpace(stdout.String())
    // Accept true/false or 1/0
    switch strings.ToLower(s) {
    case "true", "1":
        return true, true, nil
    case "false", "0":
        return false, true, nil
    default:
        // Fallback: try numeric parse
        if n, e := strconv.Atoi(s); e == nil {
            return n != 0, true, nil
        }
        return false, true, fmt.Errorf("unexpected output for bool: %q", s)
    }
}
