//go:build darwin

package system

import (
    mac "github.com/akdlsz21/typematic/internal/drivers/macos"
)

// MacOSDetail exposes macOS-specific typematic details for --show.
type MacOSDetail struct {
    KeyRepeatTicks        int
    InitialKeyRepeatTicks int
    PressAndHoldEnabled   bool
    PressAndHoldSet       bool
}

func ReadMacOSDetail() (MacOSDetail, error) {
    d, err := mac.ReadDetail()
    if err != nil {
        if err == mac.ErrOSCall {
            return MacOSDetail{}, ErrOSCall
        }
        if err == mac.ErrUnsupported {
            return MacOSDetail{}, ErrUnsupported
        }
        return MacOSDetail{}, err
    }
    return MacOSDetail{
        KeyRepeatTicks:        d.KeyRepeatTicks,
        InitialKeyRepeatTicks: d.InitialKeyRepeatTicks,
        PressAndHoldEnabled:   d.PressAndHoldEnabled,
        PressAndHoldSet:       d.PressAndHoldSet,
    }, nil
}

