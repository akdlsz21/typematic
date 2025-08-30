package system

import "errors"

var (
    ErrUnsupported = errors.New("unsupported/undetected environment")
    ErrOSCall      = errors.New("os call failed")
)

