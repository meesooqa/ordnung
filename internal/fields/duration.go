package fields

import (
	"fmt"
	"slices"

	"google.golang.org/api/youtube/v3"

	"github.com/meesooqa/go-ytpl-custom-sort/internal/tools"
	"github.com/meesooqa/go-ytpl-custom-sort/internal/video"
)

const DURATION = "duration"

type Duration struct{}

func NewDuration() Field {
	return &Duration{}
}

func (Duration) Value(v *youtube.Video) (any, error) {
	value, err := tools.ParseYtDuration(v.ContentDetails.Duration)
	if err != nil {
		return nil, fmt.Errorf("unable to parse duration: %v", err)
	}
	return value, nil
}

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
