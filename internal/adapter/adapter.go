package adapter

import (
	"fmt"
	"time"

	"google.golang.org/api/youtube/v3"

	"github.com/meesooqa/go-ytpl-custom-sort/internal/fields"
	"github.com/meesooqa/go-ytpl-custom-sort/internal/video"
)

type Adapter struct {
	ff map[string]fields.Field
}

func NewAdapter(ff map[string]fields.Field) YtAdapter {
	return &Adapter{
		ff: ff,
	}
}

func (a *Adapter) ConvertItems(items []*youtube.Video) ([]video.YtVideo, error) {
	var result []video.YtVideo
	for _, item := range items {
		videoItem, err := a.convert(item)
		if err != nil {
			return nil, err
		}
		result = append(result, videoItem)
	}
	return result, nil
}

func (a *Adapter) convert(item *youtube.Video) (video.YtVideo, error) {
	duration, err := convertField[time.Duration](a, fields.DURATION, item)
	if err != nil {
		return nil, err
	}

	return video.NewVideo(
		item.Id,
		duration,
	), nil
}

func convertField[T any](a *Adapter, code string, item *youtube.Video) (T, error) {
	var zero T
	f, ok := a.ff[code]
	if !ok {
		return zero, fmt.Errorf("field not found: %s", code)
	}
	value, err := f.Value(item)
	if err != nil {
		return zero, fmt.Errorf("unable to parse %s: %v", code, err)
	}

	if typedValue, ok := value.(T); ok {
		return typedValue, nil
	}

	return zero, fmt.Errorf("type mismatch for field %s: expected %T, got %T", code, zero, value)
}
