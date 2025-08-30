# Repository Guidelines

## Project Structure & Module Organization
- `cmd/typematic/`: CLI entrypoint (`main.go`, Cobra `root.go`).
- `internal/platform/`: Cross‑platform facade (`Set`, `Read`) with build tags for `linux` and `windows`.
- `internal/linux/`, `internal/windows/`: OS integrations (GNOME gsettings; Windows SystemParametersInfo).
- `internal/units/`: Validation and conversions (CPS↔interval).
- `internal/run/`: Command runner with `Real` and test `Fake`.
- Tests live next to code as `*_test.go`.

## Build, Test, and Development Commands
- Build: `go build -o typematic ./cmd/typematic`
- Run (Linux example): `./typematic --delay-ms=250 --rate-cps=25`
- Show current settings (Linux): `./typematic --show`
- Tests: `go test ./...`
- Coverage: `go test -cover ./...`
- Vet (sanity checks): `go vet ./...`

## Coding Style & Naming Conventions
- Go formatting: run `gofmt -s -w .` before pushing.
- Use standard Go idioms; keep packages lower_snakecase (`internal/units`).
- Platform code uses build tags (`//go:build linux`, `//go:build windows`) and mirrors file names across OSes where practical.
- Tests: table‑driven where helpful; name tests `TestXxx` matching exported behavior.

## Testing Guidelines
- Framework: standard `go test` with co‑located tests.
- Prefer testing via `internal/run.Fake` for command execution and capturing stdout where needed.
- Target: keep/raise coverage for changed packages; include happy‑path and error cases (e.g., unsupported env, OS call failures).
- Quick smoke: `go test ./internal/... ./cmd/...`

## Commit & Pull Request Guidelines
- Commits: concise, imperative subject (e.g., "linux: validate rate before set").
- Group related changes; keep refactors separate from behavior changes.
- PRs must include: summary, rationale, test coverage notes, platforms tested (Linux/Windows), and sample CLI output (e.g., from `--show`).
- Link issues where applicable; screenshots not required, but include command output for UX changes.

## Platform & Security Notes
- Linux requires GNOME on Wayland with `gsettings` and a session bus (`DBUS_SESSION_BUS_ADDRESS`).
- Windows uses `SystemParametersInfo` for FilterKeys; dry‑run is a no‑op.
- Avoid shelling out in new code except via `internal/run.Runner` to keep logic testable.
