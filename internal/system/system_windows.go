//go:build windows

package system

import (
    "fmt"

    win "github.com/akdlsz21/typematic/internal/drivers/windows"
)

func Set(delayMS int, rateCPS float64) error {
    if err := win.Set(delayMS, rateCPS); err != nil {
        if err == win.ErrOSCall {
            return fmt.Errorf("%w: %v", ErrOSCall, err)
        }
        return err
    }
    return nil
}

func Read() (State, error) {
    st, err := win.Read()
    if err != nil {
        if err == win.ErrOSCall {
            return State{}, fmt.Errorf("%w: %v", ErrOSCall, err)
        }
        return State{}, err
    }
    return State{DelayMS: st.DelayMS, IntervalMS: st.IntervalMS, RateCPS: st.RateCPS}, nil
}
