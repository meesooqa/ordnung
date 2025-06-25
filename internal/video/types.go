package video

import "time"

type YtVideo interface {
	ID() string
	Duration() time.Duration
}
