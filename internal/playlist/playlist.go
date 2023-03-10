package playlist

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

//song structure
type Song struct {
	Duration int32
	Title    string
	prev     *Song
	next     *Song
}

//playlist structure
type Playlist struct {
	head    *Song
	tail    *Song
	current *Song
	length  int
	sleep   bool
	done    chan struct{}
	pause   chan struct{}
	start   chan struct{}
}

func New() *Playlist {
	return &Playlist{
		done: make(chan struct{}),
	}
}

//add method
//add song in doubly linked list
func (l *Playlist) Add(val int32, val2 string) {
	newSong := &Song{Duration: val, Title: val2}
	if l.head == nil {
		l.head = newSong
		l.tail = newSong
	} else {
		l.tail.next = newSong
		newSong.prev = l.tail
		l.tail = newSong
	}
	l.length++
}

//next song
func (l *Playlist) Forward() (*Song, error) {
	l.done <- struct{}{}
	if l.current.next == nil {
		return nil, fmt.Errorf("end list")
	}
	return l.current.next, nil
}

//previous song
func (l *Playlist) Backward() (*Song, error) {
	if l.current.prev == nil {
		return nil, fmt.Errorf("end list")
	}
	return l.current.prev, nil
}

//pause song
func (l *Playlist) Pause() {
	l.pause <- struct{}{}
}

//start playlist
func (l *Playlist) Start() {
	l.start <- struct{}{}
}

//song delete method
func (l *Playlist) DeleteSong(val string) error {
	current := l.head

	for current != nil {
		if current.Title == val {
			if current.next == nil {
				l.tail = current.prev
				l.tail.next = nil
			} else if current.prev == nil {
				l.head = current.next
				l.head.prev = nil
			} else {
				current.prev.next = current.next
				current.next.prev = current.prev
			}
		}
		current = current.next
	}
	return fmt.Errorf("song not found")
}

// starts playback
func (l *Playlist) Play() error {
	var err error
	currentNode := l.head
	for currentNode != nil {
		l.current = currentNode
		ticker := time.NewTicker(time.Duration(currentNode.Duration))
	LOOP:
		for {
			select {
			case <-ticker.C:
				ticker.Stop()
			case <-l.done:
				break LOOP
			case <-l.pause:
				ticker.Stop()
			}
		}
		currentNode, err = l.Forward()
		if err != nil {
			log.Err(err).Msgf("err jump to next song: %w", err)
			return fmt.Errorf("err jump to next song: %w", err)
		}
	}
	return nil
}

// return current song
func (l *Playlist) Get() *Song {
	return l.current
}

// update title song
func (l *Playlist) UpdateSong(old, new string) error {
	current := l.head
	for current != nil {
		if current.Title == old {
			current.Title = new
			return nil
		}
	}
	return fmt.Errorf("song not found")
}
