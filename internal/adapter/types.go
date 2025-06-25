package adapter

import (
	"google.golang.org/api/youtube/v3"

	"github.com/meesooqa/go-ytpl-custom-sort/internal/video"
)

type YtAdapter interface {
	ConvertItems([]*youtube.Video) ([]video.YtVideo, error)
}
