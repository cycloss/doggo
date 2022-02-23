# Doggo

Docker logging go.

A simple level based logging library for go built to work with the docker `journald` logging driver so that logs can be easily filtered by priority.

## log.Print

- `log.Print` and `log.Printf` will now log to stdout, allowing them to be interpreted as priority level 6 (notice) in `journalctl`.
- This causes messages to appear in white.
- `journalctl -p 6..6 CONTAINER_ID=<ID>` to filter for only notice logs.

## log.Error

- `log.Error` and `log.Errorf` have the behaviour of default logging and log to stderr. This is interpreted as level 3 (error) in `journalctl`.
- This causes messasges to appear in red.
- `journalctl -p 3..3 CONTAINER_ID=<ID>` to filter for only error logs.
