//go:build windows

package system

import (
    win "github.com/akdlsz21/typematic/internal/drivers/windows"
)

// WindowsDetail exposes Windows-specific typematic details for --show.
type WindowsDetail struct {
    Enabled       bool
    Flags         uint32
    WaitMSec      uint32
    DelayMSec     uint32
    RepeatMSec    uint32
    BounceMSec    uint32
    KeyboardSpeed uint32
    KeyboardDelay uint32
}

func ReadWindowsDetail() (WindowsDetail, error) {
    d, err := win.ReadDetail()
    if err != nil {
        if err == win.ErrOSCall {
            return WindowsDetail{}, ErrOSCall
        }
        return WindowsDetail{}, err
    }
    return WindowsDetail{
        Enabled:       d.Enabled,
        Flags:         d.Flags,
        WaitMSec:      d.WaitMSec,
        DelayMSec:     d.DelayMSec,
        RepeatMSec:    d.RepeatMSec,
        BounceMSec:    d.BounceMSec,
        KeyboardSpeed: d.KeyboardSpeed,
        KeyboardDelay: d.KeyboardDelay,
    }, nil
}

