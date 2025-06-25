package yt

import (
	"fmt"

	"google.golang.org/api/youtube/v3"

	"github.com/meesooqa/go-ytpl-custom-sort/internal/adapter"
	"github.com/meesooqa/go-ytpl-custom-sort/internal/video"
)

type Service struct {
	s       *youtube.Service
	adapter adapter.YtAdapter
}

func NewService(s *youtube.Service, adapter adapter.YtAdapter) Yt {
	return &Service{
		s:       s,
		adapter: adapter,
	}
}

func (yt *Service) PlaylistVideo(playlistID string) ([]video.YtVideo, error) {
	ids, err := yt.playlistVideoId(playlistID)
	if err != nil {
		return nil, err
	}

	r, err := yt.s.Videos.List([]string{"contentDetails", "snippet", "statistics", "status"}).
		Id(ids...).
		Do()
	if err != nil {
		return nil, fmt.Errorf("unable to get video by ids: %v", err)
	}

	var result []video.YtVideo
	for _, item := range r.Items {
		videoItem, err := yt.adapter.Convert(item)
		if err != nil {
			return nil, err
		}
		result = append(result, videoItem)
	}
	return result, nil
}

func (yt *Service) playlistVideoId(id string) ([]string, error) {
	r, err := yt.s.PlaylistItems.List([]string{"contentDetails"}).
		PlaylistId(id).
		Do()
	if err != nil {
		return nil, fmt.Errorf("unable to get playlist items: %v", err)
	}

	var result []string
	for _, item := range r.Items {
		result = append(result, item.ContentDetails.VideoId)
	}
	return result, nil
}
