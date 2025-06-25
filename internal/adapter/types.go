package adapter

import (
	"google.golang.org/api/youtube/v3"

	"github.com/meesooqa/go-ytpl-custom-sort/internal/video"
)

// YtAdapter defines the interface for converting YouTube video items to a custom video type
type YtAdapter interface {
	ConvertItems([]*youtube.Video) ([]video.YtVideo, error)
}
