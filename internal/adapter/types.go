package adapter

import (
	"google.golang.org/api/youtube/v3"

	"github.com/meesooqa/go-ytpl-custom-sort/internal/video"
)

type YtAdapter interface {
	Convert(*youtube.Video) (video.YtVideo, error)
	ConvertItems([]*youtube.Video) ([]video.YtVideo, error)
}
