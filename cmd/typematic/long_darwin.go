//go:build darwin

package main

const longDesc = "typematic sets the keyboard auto-repeat delay and rate on macOS using defaults (KeyRepeat/InitialKeyRepeat). macOS quantizes to 15 ms ticks; values are clamped accordingly. --show prints delay_ms, interval_ms, rate_cps, plus macOS-specific details."

