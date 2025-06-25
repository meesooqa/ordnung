package video

import "time"

// YtVideo is an interface that defines methods for interacting with YouTube videos
type YtVideo interface {
	ID() string
	Duration() time.Duration
}
