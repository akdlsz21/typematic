//go:build linux

package gnome

import (
    "bytes"
    "fmt"
    "os"
    "os/exec"
    "strings"

    "github.com/akdlsz21/typematic/internal/units"
)

// ready performs a minimal readiness check for using gsettings:
// - DBUS_SESSION_BUS_ADDRESS must be set
// - gsettings must be present in PATH
func ready() (bool, string) {
    if strings.TrimSpace(os.Getenv("DBUS_SESSION_BUS_ADDRESS")) == "" {
        return false, "DBUS_SESSION_BUS_ADDRESS is empty"
    }
    if _, err := lookPath("gsettings"); err != nil {
        return false, "gsettings not found in PATH"
    }
    return true, ""
}

// Set applies the typematic settings via gsettings on GNOME Wayland.
func Set(delayMS int, rateCPS float64) error {
    intervalMS, err := units.Validate(delayMS, rateCPS)
    if err != nil {
        return err
    }

    ok, why := ready()
    if !ok {
        return fmt.Errorf("%w: %s", ErrUnsupported, why)
    }

    schema := "org.gnome.desktop.peripherals.keyboard"
    if err := gsettingsSet(schema, "delay", fmt.Sprintf("%d", delayMS)); err != nil {
        return err
    }
    if err := gsettingsSet(schema, "repeat-interval", fmt.Sprintf("%d", intervalMS)); err != nil {
        return err
    }
    return nil
}

// gsettingsSet is a tiny helper to run a single gsettings set command with
// consistent error mapping.
func gsettingsSet(schema, key, value string) error {
    var stdout, stderr bytes.Buffer
    cmd := exec.Command("gsettings", "set", schema, key, value)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("%w: %s", ErrOSCall, strings.TrimSpace(stderr.String()))
    }
    return nil
}
