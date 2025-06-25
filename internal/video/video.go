package video

import "time"

// Video is a struct that implements the YtVideo interface for YouTube videos
type Video struct {
	youtubeVideoID string
	duration       time.Duration
}

// NewVideo creates a new instance of Video with the provided YouTube video ID and duration
func NewVideo(youtubeVideoID string, duration time.Duration) *Video {
	return &Video{
		youtubeVideoID: youtubeVideoID,
		duration:       duration,
	}
}

// ID returns the YouTube video ID of the Video instance
func (v *Video) ID() string {
	return v.youtubeVideoID
}

// Duration returns the duration of the Video instance
func (v *Video) Duration() time.Duration {
	return v.duration
}
