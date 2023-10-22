package assetloader

import (
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/wav"
	"os"
	"strings"
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
	assetType, buffer, format, err := al.loadAudioBuffer(path)
	if err != nil {
		return AssetResource[AudioAsset]{}, err
	}

	data := AudioAsset{
		Buffer: buffer,
		Format: format,
	}

	asset := AssetResource[AudioAsset]{
		Type: assetType,
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
	assetType, streamer, format, err := al.loadAudio(path)
	if err != nil {
		return AssetResource[AudioStreamAsset]{}, err
	}

	data := AudioStreamAsset{
		Stream: streamer,
		Format: format,
	}

	asset := AssetResource[AudioStreamAsset]{
		Type: assetType,
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

func (al *AudioLoader) loadAudioBuffer(path string) (AssetType, *beep.Buffer, beep.Format, error) {
	assetType, streamer, format, err := al.loadAudio(path)

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	err = streamer.Close()
	if err != nil {
		return assetType, nil, beep.Format{}, err
	}

	return assetType, buffer, format, nil
}

func (al *AudioLoader) loadAudio(path string) (AssetType, beep.StreamSeekCloser, beep.Format, error) {
	// Get the file extension.
	rawType := strings.Split(path, ".")[1]

	var assetType AssetType
	var streamer beep.StreamSeekCloser
	var format beep.Format
	var err error
	switch rawType {
	case "mp3":
		assetType = Mp3AudioAssetType
		streamer, format, err = al.loadMp3Audio(path)
	case "wav":
		assetType = Mp3AudioAssetType
		streamer, format, err = al.loadWavAudio(path)
	default:
		return assetType, nil, beep.Format{}, ErrInvalidFileFormat{rawType: rawType}
	}

	return assetType, streamer, format, err
}

func (al *AudioLoader) loadMp3Audio(path string) (beep.StreamSeekCloser, beep.Format, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, beep.Format{}, err
	}

	return mp3.Decode(file)
}

func (al *AudioLoader) loadWavAudio(path string) (beep.StreamSeekCloser, beep.Format, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, beep.Format{}, err
	}

	return wav.Decode(file)
}
