//go:build darwin

package macos

import (
    "errors"
    "os/exec"
)

var (
    // ErrUnsupported indicates the current environment is not supported
    // (e.g., required command not found).
    ErrUnsupported = errors.New("unsupported or undetected environment")
    // ErrOSCall indicates an underlying system command failed.
    ErrOSCall = errors.New("os call failed")
)

// Allow tests to stub lookPath.
var lookPath = exec.LookPath

