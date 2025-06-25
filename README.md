# Ordnung
## Sort Youtube Playlist Videos

The application sorts video in the YouTube playlist.
It creates a new playlist based on the existing
and sorts video by a certain criterion (by `duration` by default).
For work, authorization is required through Google API (YouTube Data API V3).

### Requirements

- Go 1.18+
- Make (optional, for building)
- Google API credentials for YouTube Data API v3

### Usage

1. Create a project in the [Google Cloud Console](https://console.developers.google.com/).
2. Enable YouTube Data API v3, scope: `youtube`.
3. Create OAuth 2.0 credentials.
4. Set environment variables `CLIENT_ID` and `SECRET` in `.env` file.

### Environment Variables

- `CLIENT_ID`: Your OAuth 2.0 Client ID.
- `SECRET`: Your OAuth 2.0 Client Secret.

### Installation

1. Clone the repository:

```shell
git clone https://github.com/meesooqa/ordnung
cd ordnung
```

2. Build:

```shell
make build
```

3. Run the application with the command line argument `-pls` followed by the playlist ID you want to sort.

```shell
./build/ordnung -pls=PLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

Or

```shell
./build/ordnung -remove -pls=PLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
to clear playlist pls.

### Sort

By default the app sorts by `duration`.

#### For Developers: How to Add New Sort Field

1. Add struct `fields.MyField` implements `fields.Field`.
2. Add field to `video.Video` and getter method to `video.YtVideo`.
3. Edit adapter `adapter.Adapter.convert()`.
4. Register field in `./cmd/app/main.go`:
```go
ff := map[string]fields.Field{
    ...
    fields.MY_FIELD_CODE: fields.NewMyField(),
}
```
5. Build the app
6. Run
```shell
./build/ordnung -sort=myfield -pls=PLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
