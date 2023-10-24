package assetloader

import (
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/wav"
	"os"
)

// AudioClipAsset is a type that defines a audio clip asset. It holds the buffer and format.
// A audio clip asset is a audio asset that is loaded into memory.
type AudioClipAsset struct {
	Buffer *beep.Buffer
	Format beep.Format
}

// AudioStreamAsset is a type that defines a audio stream asset. It holds the stream and format.
// A audio stream asset is a audio asset that is streamed from the disk.
type AudioStreamAsset struct {
	Stream beep.StreamSeekCloser
	Format beep.Format
}

type AudioLoader struct {
	clips     map[string]AssetResource[AudioClipAsset]
	streamers map[string]AssetResource[AudioStreamAsset]
}

func NewAudioLoader() AudioLoader {
	return AudioLoader{
		clips:     make(map[string]AssetResource[AudioClipAsset]),
		streamers: make(map[string]AssetResource[AudioStreamAsset]),
	}
}

// LoadAudioClip loads a audio clip asset into the asset loader.
func (al *AudioLoader) LoadAudioClip(name, path string) (AssetResource[AudioClipAsset], error) {
	assetType, err := GetAssetTypeByPath(path)
	if err != nil {
		return AssetResource[AudioClipAsset]{}, err
	}

	var streamer beep.StreamSeekCloser
	var format beep.Format

	switch assetType {
	case Mp3AudioAssetType:
		streamer, format, err = al.loadMp3Audio(path)
		if err != nil {
			return AssetResource[AudioClipAsset]{}, err
		}
	case WavAudioAssetType:
		streamer, format, err = al.loadWavAudio(path)
		if err != nil {
			return AssetResource[AudioClipAsset]{}, err
		}
	default:
		return AssetResource[AudioClipAsset]{}, ErrUnsupportedAssetType{assetType: assetType}
	}

	buffer, err := al.createBuffer(streamer, format)
	if err != nil {
		return AssetResource[AudioClipAsset]{}, err
	}

	data := AudioClipAsset{
		Buffer: buffer,
		Format: format,
	}

	asset := AssetResource[AudioClipAsset]{
		Type: assetType,
		Name: name,
		Path: path,
		Data: data,
	}

	al.clips[name] = asset

	return asset, nil
}

// GetAudioClip gets a audio clip asset from the asset loader.
func (al *AudioLoader) GetAudioClip(name string) (AssetResource[AudioClipAsset], bool) {
	asset, ok := al.clips[name]
	return asset, ok
}

// GetAudioClips gets all audio clip assets from the asset loader.
func (al *AudioLoader) GetAudioClips() map[string]AssetResource[AudioClipAsset] {
	return al.clips
}

// RemoveAudioClip removes a audio clip asset from the asset loader.
func (al *AudioLoader) RemoveAudioClip(name string) {
	delete(al.clips, name)
}

// EachAudioClip iterates over all audio clip assets in the asset loader.
func (al *AudioLoader) EachAudioClip(fn func(name string, asset AssetResource[AudioClipAsset])) {
	for name, asset := range al.clips {
		fn(name, asset)
	}
}

// LoadAudioStream loads a audio stream asset into the asset loader.
func (al *AudioLoader) LoadAudioStream(name, path string) (AssetResource[AudioStreamAsset], error) {
	assetType, err := GetAssetTypeByPath(path)
	if err != nil {
		return AssetResource[AudioStreamAsset]{}, err
	}

	var streamer beep.StreamSeekCloser
	var format beep.Format

	switch assetType {
	case Mp3AudioAssetType:
		assetType = Mp3AudioStreamAssetType
		streamer, format, err = al.loadMp3Audio(path)
		if err != nil {
			return AssetResource[AudioStreamAsset]{}, err
		}
	case WavAudioAssetType:
		assetType = WavAudioStreamAssetType
		streamer, format, err = al.loadWavAudio(path)
		if err != nil {
			return AssetResource[AudioStreamAsset]{}, err
		}
	default:
		return AssetResource[AudioStreamAsset]{}, ErrUnsupportedAssetType{assetType: assetType}
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

func (al *AudioLoader) createBuffer(streamer beep.StreamSeekCloser, format beep.Format) (*beep.Buffer, error) {
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	err := streamer.Close()
	if err != nil {
		return nil, err
	}

	return buffer, nil
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
