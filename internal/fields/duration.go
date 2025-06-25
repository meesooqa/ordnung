package fields

import (
	"fmt"
	"slices"

	"google.golang.org/api/youtube/v3"

	"github.com/meesooqa/ordnung/internal/tools"
	"github.com/meesooqa/ordnung/internal/video"
)

// DURATION is the name of the duration field used for sorting YouTube videos
const DURATION = "duration"

// Duration represents a field for sorting YouTube videos by their duration.
type Duration struct{}

// NewDuration creates a new Duration field instance
func NewDuration() Field {
	return &Duration{}
}

// Value extracts the duration value from a YouTube video
func (Duration) Value(v *youtube.Video) (any, error) {
	value, err := tools.ParseYtDuration(v.ContentDetails.Duration)
	if err != nil {
		return nil, fmt.Errorf("unable to parse duration: %v", err)
	}
	return value, nil
}

// Sort sorts a slice of YouTube videos by their duration
func (Duration) Sort(items []video.YtVideo) {
	slices.SortFunc(items, func(a, b video.YtVideo) int {
		aDur := a.Duration()
		bDur := b.Duration()

		switch {
		case aDur < bDur:
			return -1
		case aDur > bDur:
			return 1
		default:
			return 0
		}
	})
}
