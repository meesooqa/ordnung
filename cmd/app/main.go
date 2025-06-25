package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"

	"github.com/meesooqa/go-ytpl-custom-sort/internal/adapter"
	"github.com/meesooqa/go-ytpl-custom-sort/internal/fields"
	"github.com/meesooqa/go-ytpl-custom-sort/internal/yt"
)

func main() {
	sortBy := "duration"
	playlistID := "PLufCON52KijtaZCLAdC9hf7ykxMGjl6mT" // Test [Public]

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	err = run(playlistID, sortBy)
	if err != nil {
		log.Fatal(err)
	}
}

func run(playlistID, sortBy string) error {
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

	err = ytService.CopyAndSortPlaylist(playlistID, sortBy)
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
			youtube.YoutubeReadonlyScope,
		},
	}
	return newOAuthClient(ctx, config)
}
