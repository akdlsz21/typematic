//go:build windows

package windows

import (
    "fmt"
    "unsafe"

    "github.com/zzl/go-win32api/v2/win32"
)

// State represents current typematic settings on Windows (FilterKeys).
type State struct {
    DelayMS    int
    IntervalMS int
    RateCPS    float64
}

// Read queries the system via SystemParametersInfoW(SPI_GETFILTERKEYS) and returns
// the current FilterKeys delay and repeat interval. RateCPS is derived as 1000/interval.
func Read() (State, error) {
    var fk win32.FILTERKEYS
    fk.CbSize = uint32(unsafe.Sizeof(fk))
    if ret, winErr := win32.SystemParametersInfoW(win32.SPI_GETFILTERKEYS, fk.CbSize, unsafe.Pointer(&fk), 0); ret == 0 {
        return State{}, fmt.Errorf("%w: %v", ErrOSCall, winErr)
    }
    delay := int(fk.IDelayMSec)
    interval := int(fk.IRepeatMSec)
    cps := 0.0
    if interval > 0 {
        cps = 1000.0 / float64(interval)
    }
    return State{DelayMS: delay, IntervalMS: interval, RateCPS: cps}, nil
}

// Detail includes full FilterKeys fields and classic SPI keyboard settings.
type Detail struct {
    Enabled       bool
    Flags         uint32
    WaitMSec      uint32
    DelayMSec     uint32
    RepeatMSec    uint32
    BounceMSec    uint32
    KeyboardSpeed uint32 // 0..31 as per SPI_GETKEYBOARDSPEED
    KeyboardDelay uint32 // 0..3 as per SPI_GETKEYBOARDDELAY
}

// ReadDetail returns extended Windows-specific typematic information.
func ReadDetail() (Detail, error) {
    var fk win32.FILTERKEYS
    fk.CbSize = uint32(unsafe.Sizeof(fk))
    if ret, winErr := win32.SystemParametersInfoW(win32.SPI_GETFILTERKEYS, fk.CbSize, unsafe.Pointer(&fk), 0); ret == 0 {
        return Detail{}, fmt.Errorf("%w: %v", ErrOSCall, winErr)
    }
    var speed uint32
    if ret, _ := win32.SystemParametersInfoW(win32.SPI_GETKEYBOARDSPEED, 0, unsafe.Pointer(&speed), 0); ret == 0 {
        speed = 0
    }
    var kdelay uint32
    if ret, _ := win32.SystemParametersInfoW(win32.SPI_GETKEYBOARDDELAY, 0, unsafe.Pointer(&kdelay), 0); ret == 0 {
        kdelay = 0
    }
    return Detail{
        Enabled:       (fk.DwFlags & win32.FKF_FILTERKEYSON) != 0,
        Flags:         fk.DwFlags,
        WaitMSec:      fk.IWaitMSec,
        DelayMSec:     fk.IDelayMSec,
        RepeatMSec:    fk.IRepeatMSec,
        BounceMSec:    fk.IBounceMSec,
        KeyboardSpeed: speed,
        KeyboardDelay: kdelay,
    }, nil
}
