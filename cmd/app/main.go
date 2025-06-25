package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"

	"github.com/meesooqa/ordnung/internal/adapter"
	"github.com/meesooqa/ordnung/internal/fields"
	"github.com/meesooqa/ordnung/internal/yt"
)

var (
	sort   = flag.String("sort", fields.DURATION, "Sort By")
	pls    = flag.String("pls", "", "Playlist ID from URL")
	remove = flag.Bool("remove", false, "Remove items from playlist")
)

func main() {
	flag.Parse()
	if *pls == "" {
		fmt.Println("[ERROR] `-pls`")
		flag.Usage()
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Printf("sortBy %s, playlistID %s, remove %t", *sort, *pls, *remove)
	err = run(*pls, *sort, *remove)
	if err != nil {
		log.Fatal(err)
	}
}

func run(playlistID, sortBy string, remove bool) error {
	ctx := context.Background()
	client := getClient(ctx)
	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("unable to create YouTube service: %v", err)
	}

	ff := map[string]fields.Field{
		fields.DURATION: fields.NewDuration(),
	}
	ytService := yt.NewService(service, ff, adapter.NewAdapter(ff))

	err = ytService.CopyAndSortPlaylist(playlistID, sortBy, remove)
	if err != nil {
		return fmt.Errorf("unable to copy playlist videos: %v", err)
	}
	return nil
}

func getClient(ctx context.Context) *http.Client {
	clientID := os.Getenv("CLIENT_ID")
	secret := os.Getenv("SECRET")
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: secret,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			youtube.YoutubeScope,
		},
	}
	return yt.NewOAuthClient(ctx, config)
}
