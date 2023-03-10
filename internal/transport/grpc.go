package transport

import (
	"context"

	"github.com/Anvyyy/playlist/internal/playlist"
	"github.com/Anvyyy/playlist/pkg"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Playlist struct {
	playlist playlist.Interface
}

func NewGrpcTransport() *Playlist {
	return &Playlist{
		playlist: playlist.New(),
	}
}

func (p *Playlist) PlaySong(context.Context, *pkg.Empty) (*pkg.Empty, error) {
	go func() {
		err := p.playlist.Play()
		if err != nil {
			log.Err(err).Msgf("error play songs: %w", err)
		}
	}()
	return &pkg.Empty{}, nil
}

func (p *Playlist) PauseSong(context.Context, *pkg.Empty) (*pkg.Empty, error) {
	p.playlist.Pause()
	return &pkg.Empty{}, nil
}

func (p *Playlist) AddSong(ctx context.Context, req *pkg.AddSongRequest) (*pkg.Empty, error) {
	p.playlist.Add(req.Duration, req.Name)
	return &pkg.Empty{}, nil
}

func (p *Playlist) NextSong(context.Context, *pkg.Empty) (*pkg.SongResponse, error) {
	song, err := p.playlist.Forward()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "error jump to next song: %w", err)
	}

	return &pkg.SongResponse{
		Name:     song.Title,
		Duration: song.Duration,
	}, nil
}

func (p *Playlist) PrevSong(context.Context, *pkg.Empty) (*pkg.SongResponse, error) {
	song, err := p.playlist.Backward()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "error jump to prev song: %w", err)
	}

	return &pkg.SongResponse{
		Name:     song.Title,
		Duration: song.Duration,
	}, nil
}

func (p *Playlist) GetSong(ctx context.Context, req *pkg.RequestSong) (*pkg.SongResponse, error) {
	song := p.playlist.Get()
	if song == nil {
		return nil, status.Errorf(codes.Internal, "player isn't playing")
	}

	return &pkg.SongResponse{}, nil
}

func (p *Playlist) UpdateSong(ctx context.Context, req *pkg.Update) (*pkg.Empty, error) {
	err := p.playlist.UpdateSong(req.OldName, req.NewName)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "song not found: %w", err)
	}
	return &pkg.Empty{}, nil
}

func (p *Playlist) DeleteSong(ctx context.Context, req *pkg.RequestSong) (*pkg.Empty, error) {
	err := p.playlist.DeleteSong(req.Name)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "song not found: %w", err)
	}
	return &pkg.Empty{}, nil
}
