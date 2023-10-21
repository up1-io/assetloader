package assetloader

import (
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"os"
)

type AudioAsset struct {
	Buffer *beep.Buffer
	Format beep.Format
}

type AudioStreamAsset struct {
	Stream beep.StreamSeekCloser
	Format beep.Format
}

type AudioLoader struct {
	clips     map[string]AssetResource[AudioAsset]
	streamers map[string]AssetResource[AudioStreamAsset]
}

func NewAudioLoader() AudioLoader {
	return AudioLoader{
		clips:     make(map[string]AssetResource[AudioAsset]),
		streamers: make(map[string]AssetResource[AudioStreamAsset]),
	}
}

// LoadAudioClip loads a audio clip asset into the asset loader.
func (al *AudioLoader) LoadAudioClip(name, path string) (AssetResource[AudioAsset], error) {
	buffer, format, err := al.loadAudioBuffer(path)
	if err != nil {
		return AssetResource[AudioAsset]{}, err
	}

	data := AudioAsset{
		Buffer: buffer,
		Format: format,
	}

	asset := AssetResource[AudioAsset]{
		Type: Mp3AudioAssetType,
		Name: name,
		Path: path,
		Data: data,
	}

	al.clips[name] = asset

	return asset, nil
}

// GetAudioClip gets a audio clip asset from the asset loader.
func (al *AudioLoader) GetAudioClip(name string) (AssetResource[AudioAsset], bool) {
	asset, ok := al.clips[name]
	return asset, ok
}

// GetAudioClips gets all audio clip assets from the asset loader.
func (al *AudioLoader) GetAudioClips() map[string]AssetResource[AudioAsset] {
	return al.clips
}

// RemoveAudioClip removes a audio clip asset from the asset loader.
func (al *AudioLoader) RemoveAudioClip(name string) {
	delete(al.clips, name)
}

// EachAudioClip iterates over all audio clip assets in the asset loader.
func (al *AudioLoader) EachAudioClip(fn func(name string, asset AssetResource[AudioAsset])) {
	for name, asset := range al.clips {
		fn(name, asset)
	}
}

// LoadAudioStream loads a audio stream asset into the asset loader.
func (al *AudioLoader) LoadAudioStream(name, path string) (AssetResource[AudioStreamAsset], error) {
	streamer, format, err := al.loadAudioStream(path)
	if err != nil {
		return AssetResource[AudioStreamAsset]{}, err
	}

	data := AudioStreamAsset{
		Stream: streamer,
		Format: format,
	}

	asset := AssetResource[AudioStreamAsset]{
		Type: Mp3AudioStreamAssetType,
		Name: name,
		Path: path,
		Data: data,
	}

	al.streamers[name] = asset

	return asset, nil
}

// GetAudioStream gets a audio stream asset from the asset loader.
func (al *AudioLoader) GetAudioStream(name string) (AssetResource[AudioStreamAsset], bool) {
	asset, ok := al.streamers[name]
	return asset, ok
}

// GetAudioStreams gets all audio stream assets from the asset loader.
func (al *AudioLoader) GetAudioStreams() map[string]AssetResource[AudioStreamAsset] {
	return al.streamers
}

// RemoveAudioStream removes a audio stream asset from the asset loader.
func (al *AudioLoader) RemoveAudioStream(name string) {
	delete(al.streamers, name)
}

// EachAudioStream iterates over all audio stream assets in the asset loader.
func (al *AudioLoader) EachAudioStream(fn func(name string, asset AssetResource[AudioStreamAsset])) {
	for name, asset := range al.streamers {
		fn(name, asset)
	}
}

func (al *AudioLoader) loadAudioBuffer(path string) (*beep.Buffer, beep.Format, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, beep.Format{}, err
	}

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		return nil, beep.Format{}, err
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	err = streamer.Close()
	if err != nil {
		return nil, beep.Format{}, err
	}

	return buffer, format, nil
}

func (al *AudioLoader) loadAudioStream(path string) (beep.StreamSeekCloser, beep.Format, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, beep.Format{}, err
	}

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		return nil, beep.Format{}, err
	}

	return streamer, format, nil
}
