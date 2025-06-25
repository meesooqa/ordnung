package yt

import (
	"fmt"
	"log"

	"google.golang.org/api/youtube/v3"

	"github.com/meesooqa/go-ytpl-custom-sort/internal/adapter"
	"github.com/meesooqa/go-ytpl-custom-sort/internal/fields"
	"github.com/meesooqa/go-ytpl-custom-sort/internal/video"
)

type Service struct {
	s       *youtube.Service
	ff      map[string]fields.Field
	adapter adapter.YtAdapter

	pl YtPl
}

func NewService(s *youtube.Service, ff map[string]fields.Field, adapter adapter.YtAdapter) Yt {
	return &Service{
		s:       s,
		ff:      ff,
		adapter: adapter,

		pl: NewPl(s),
	}
}

func (yt *Service) CopyAndSortPlaylist(id, sortBy string) error {
	// sort
	field, ok := yt.ff[sortBy]
	if !ok {
		return fmt.Errorf("unable to find sort field %s", sortBy)
	}
	videos, err := yt.playlistVideo(id)
	if err != nil {
		return fmt.Errorf("unable to get playlist videos: %v", err)
	}
	field.Sort(videos)

	// copy
	plFrom, err := yt.pl.FindByID(id)
	if err != nil {
		return fmt.Errorf("unable to get playlist info: %v", err)
	}
	newTitle := fmt.Sprintf("%s [sorted by %s]", plFrom.Snippet.Title, sortBy)
	plTo, err := yt.findOrCreatePlaylistByTitle(newTitle)
	if err != nil {
		return fmt.Errorf("unable to get find or create playlist %s: %v", newTitle, err)
	}

	for idx, v := range videos {
		if err = yt.pl.AddItem(plTo.Id, v.ID(), int64(idx)); err != nil {
			return fmt.Errorf("unable to add %s: %v", v.ID(), err)
		}
		log.Printf("[INFO] added %s to playlist %s", v.ID(), plTo.Id)
		_ = yt.pl.RemoveItem(plFrom.Id, v.ID())
	}

	return nil
}

func (yt *Service) playlistVideo(playlistID string) ([]video.YtVideo, error) {
	ids, err := yt.pl.ItemsId(playlistID)
	if err != nil {
		return nil, err
	}

	r, err := yt.s.Videos.List([]string{"contentDetails", "snippet", "statistics", "status"}).
		Id(ids...).
		Do()
	if err != nil {
		return nil, fmt.Errorf("unable to get video by ids: %v", err)
	}

	return yt.adapter.ConvertItems(r.Items)
}

func (yt *Service) findOrCreatePlaylistByTitle(title string) (*youtube.Playlist, error) {
	pl, err := yt.pl.FindByTitle(title)
	if err != nil {
		return nil, fmt.Errorf("unable to find playlist by title %s: %v", title, err)
	}
	if pl == nil {
		pl, err = yt.pl.Create(title)
	}
	if err != nil {
		return nil, fmt.Errorf("unable to create playlist %s: %v", title, err)
	}
	return pl, nil
}
