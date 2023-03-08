package playlist


// Interface 
type Interface interface {
	Add(val int32, val2 string)
	Forward() (*Song, error)
	Backward() (*Song, error) 
	Play() error
	Pause()
	Get() *Song
	UpdateSong(old, new string)
	DeleteSong(val string) error
}