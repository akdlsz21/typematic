//go:build linux

package system

import (
    "fmt"

    gnome "github.com/akdlsz21/typematic/internal/drivers/gnome"
)

func Set(delayMS int, rateCPS float64) error {
    if err := gnome.Set(delayMS, rateCPS); err != nil {
        if err == gnome.ErrUnsupported {
            return fmt.Errorf("%w: %v", ErrUnsupported, err)
        }
        if err == gnome.ErrOSCall {
            return fmt.Errorf("%w: %v", ErrOSCall, err)
        }
        return err
    }
    return nil
}

func Read() (State, error) {
    st, err := gnome.Read()
    if err != nil {
        if err == gnome.ErrUnsupported {
            return State{}, fmt.Errorf("%w: %v", ErrUnsupported, err)
        }
        if err == gnome.ErrOSCall {
            return State{}, fmt.Errorf("%w: %v", ErrOSCall, err)
        }
        return State{}, err
    }
    return State{DelayMS: st.DelayMS, IntervalMS: st.IntervalMS, RateCPS: st.RateCPS}, nil
}
