package yt

import "google.golang.org/api/youtube/v3"

// Yt is an interface for interacting with YouTube data
type Yt interface {
	CopyAndSortPlaylist(id, sortBy string) error
}

type YtPl interface {
	FindByID(id string) (*youtube.Playlist, error)

	ItemsId(id string) ([]string, error)

	FindByTitle(title string) (*youtube.Playlist, error)

	Create(title string) (*youtube.Playlist, error)

	AddItem(id, videoID string, position int64) error

	RemoveItem(id, videoID string) error
}
