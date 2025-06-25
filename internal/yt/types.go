package yt

import "github.com/meesooqa/go-ytpl-custom-sort/internal/video"

// Yt is an interface for interacting with YouTube data
type Yt interface {
	PlaylistVideo(id string) ([]video.YtVideo, error)
}
