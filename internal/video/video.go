package video

import "time"

type Video struct {
	youtubeVideoID string
	duration       time.Duration
}

func NewVideo(youtubeVideoID string, duration time.Duration) *Video {
	return &Video{
		youtubeVideoID: youtubeVideoID,
		duration:       duration,
	}
}

func (v *Video) Duration() time.Duration {
	return v.duration
}
