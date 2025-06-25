package fields

import (
	"google.golang.org/api/youtube/v3"

	"github.com/meesooqa/go-ytpl-custom-sort/internal/video"
)

// Field is an interface that defines methods for extracting values from a YouTube video and sorting a list of videos
type Field interface {
	Value(v *youtube.Video) (any, error)
	Sort(items []video.YtVideo)
}
