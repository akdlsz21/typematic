package system

// State is a platform-agnostic snapshot of typematic settings.
type State struct {
    DelayMS    int
    IntervalMS int
    RateCPS    float64
}

