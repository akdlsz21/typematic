//go:build darwin

package system

import (
    "fmt"

    mac "github.com/akdlsz21/typematic/internal/drivers/macos"
)

func Set(delayMS int, rateCPS float64) error {
    if err := mac.Set(delayMS, rateCPS); err != nil {
        if err == mac.ErrUnsupported {
            return fmt.Errorf("%w: %v", ErrUnsupported, err)
        }
        if err == mac.ErrOSCall {
            return fmt.Errorf("%w: %v", ErrOSCall, err)
        }
        return err
    }
    return nil
}

func Read() (State, error) {
    st, err := mac.Read()
    if err != nil {
        if err == mac.ErrUnsupported {
            return State{}, fmt.Errorf("%w: %v", ErrUnsupported, err)
        }
        if err == mac.ErrOSCall {
            return State{}, fmt.Errorf("%w: %v", ErrOSCall, err)
        }
        return State{}, err
    }
    return State{DelayMS: st.DelayMS, IntervalMS: st.IntervalMS, RateCPS: st.RateCPS}, nil
}
