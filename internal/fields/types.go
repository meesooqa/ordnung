package fields

import (
	"google.golang.org/api/youtube/v3"

	"github.com/meesooqa/go-ytpl-custom-sort/internal/video"
)

type Field interface {
	Value(v *youtube.Video) (any, error)
	Sort(items []video.YtVideo)
}
