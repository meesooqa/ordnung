package video

import "time"

type YtVideo interface {
	Duration() time.Duration
}
