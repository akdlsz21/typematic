//go:build windows

package windows

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/akdlsz21/typematic/internal/units"
	"github.com/zzl/go-win32api/v2/win32"
)

var (
	// ErrOSCall indicates a failure from the Windows API.
	ErrOSCall = errors.New("os call failed")
)

// Set applies auto-repeat settings via SystemParametersInfoW(SPI_SETFILTERKEYS).
func Set(delayMS int, rateCPS float64) error {
    intervalMS, err := units.Validate(delayMS, rateCPS)
    if err != nil {
        return err
    }

	// Read existing FilterKeys to preserve user's other flags, but always enable FilterKeys
	// so the configured typematic settings take effect without manual toggling.
	var fk win32.FILTERKEYS
	fk.CbSize = uint32(unsafe.Sizeof(fk))
	if r, _ := win32.SystemParametersInfoW(win32.SPI_GETFILTERKEYS, fk.CbSize, unsafe.Pointer(&fk), 0); r == 0 {
		// If reading fails, start from a minimal struct with FilterKeys on.
		fk.DwFlags = 0
	}
	// Force-enable FilterKeys regardless of prior state; preserve other bits.
	fk.DwFlags |= win32.FKF_FILTERKEYSON
	// Update only the values we manage.
	fk.IDelayMSec = uint32(delayMS)
	fk.IRepeatMSec = uint32(intervalMS)

	// Persist settings and broadcast change via SystemParametersInfoW.
    ret, winErr := win32.SystemParametersInfoW(win32.SPI_SETFILTERKEYS, fk.CbSize, unsafe.Pointer(&fk), win32.SPIF_UPDATEINIFILE|win32.SPIF_SENDCHANGE)
    if ret == 0 {
        return fmt.Errorf("%w: %v", ErrOSCall, winErr)
    }
    return nil
}
