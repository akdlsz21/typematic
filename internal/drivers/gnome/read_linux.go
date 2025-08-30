//go:build linux

package gnome

import (
    "bytes"
    "fmt"
    "os/exec"
    "strconv"
    "strings"
)

// State represents current typematic settings on Linux GNOME.
type State struct {
    DelayMS    int
    IntervalMS int
    RateCPS    float64
}

// Read queries gsettings for current delay and repeat interval.
func Read() (State, error) {
    ok, why := ready()
    if !ok {
        return State{}, fmt.Errorf("%w: %s", ErrUnsupported, why)
    }

    delay, err := getInt("org.gnome.desktop.peripherals.keyboard", "delay")
    if err != nil {
        return State{}, fmt.Errorf("%w: %v", ErrOSCall, err)
    }
    interval, err := getInt("org.gnome.desktop.peripherals.keyboard", "repeat-interval")
    if err != nil {
        return State{}, fmt.Errorf("%w: %v", ErrOSCall, err)
    }
    cps := 0.0
    if interval > 0 {
        cps = 1000.0 / float64(interval)
    }
    return State{DelayMS: delay, IntervalMS: interval, RateCPS: cps}, nil
}

func getInt(schema, key string) (int, error) {
    var stdout, stderr bytes.Buffer
    cmd := exec.Command("gsettings", "get", schema, key)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    if err := cmd.Run(); err != nil {
        s := strings.TrimSpace(stderr.String())
        if s != "" {
            return 0, fmt.Errorf("gsettings get error: %s", s)
        }
        return 0, err
    }
    s := strings.TrimSpace(stdout.String())
    fields := strings.Fields(s)
    for _, f := range fields {
        if n, e := strconv.Atoi(f); e == nil {
            return n, nil
        }
    }
    return 0, fmt.Errorf("unexpected output: %q", s)
}

