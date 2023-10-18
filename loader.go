package asset_loader

import (
	"github.com/golang/freetype/truetype"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/pixel"
	"golang.org/x/image/font"
	"image"
	_ "image/png"
	"io"
	"os"
)

type Loader struct {
	assets map[string]AssetResource
}

// AssetLoader is an interface that defines the methods that an asset loader must implement.
type AssetLoader interface {
	LoadTexture(name string, path string) (AssetResource, error)
	LoadAudio(name string, path string) (AssetResource, error)
	LoadAudioStream(name string, path string) (AssetResource, error)
	LoadTFF(name string, path string, size float64) (AssetResource, error)
	Get(name string) interface{}
	Remove(name string)
	Clear()
	Count() int
	List() []string
	Has(name string) bool
	Replace(name string, asset AssetResource)
	Rename(oldName, newName string)
	Each(fn func(name string, asset interface{}))
}

// NewLoader creates a new asset loader.
func NewLoader() *Loader {
	return &Loader{
		assets: make(map[string]AssetResource),
	}
}

// LoadTexture loads a texture asset into the asset loader.
func (al Loader) LoadTexture(name string, path string) (AssetResource, error) {
	pic, err := al.loadPicture(path)
	if err != nil {
		return AssetResource{}, err
	}

	asset := AssetResource{
		Type: PngImageAssetType,
		Name: name,
		Path: path,
		Data: pic,
	}

	al.load(name, asset)

	return asset, nil
}

// LoadAudio loads an audio asset into the asset loader.
func (al Loader) LoadAudio(name string, path string) (AssetResource, error) {
	buffer, format, err := al.loadAudioBuffer(path)
	if err != nil {
		return AssetResource{}, err
	}

	audio := AudioAsset{
		Buffer: buffer,
		Format: format,
	}

	asset := AssetResource{
		Type: Mp3AudioAssetType,
		Name: name,
		Path: path,
		Data: audio,
	}

	al.load(name, asset)

	return asset, nil
}

// LoadAudioStream loads an audio stream asset into the asset loader.
func (al Loader) LoadAudioStream(name, path string) (AssetResource, error) {
	streamer, format, err := al.loadAudioStream(path)
	if err != nil {
		return AssetResource{}, err
	}

	audio := AudioStreamAsset{
		Stream: streamer,
		Format: format,
	}

	asset := AssetResource{
		Type:    Mp3AudioStreamAssetType,
		Name:    name,
		Path:    path,
		Data:    audio,
		IsDirty: false,
	}

	al.load(name, asset)

	return asset, nil
}

// LoadTTF loads a TTF asset into the asset loader.
func (al Loader) LoadTTF(name, path string, size float64) (AssetResource, error) {
	f, err := al.loadFontFace(path, size)
	if err != nil {
		return AssetResource{}, err
	}

	asset := AssetResource{
		Type: TFFAssetType,
		Name: name,
		Path: path,
		Data: f,
	}

	al.load(name, asset)

	return asset, nil
}

// Get returns an asset from the asset loader.
func (al *Loader) Get(name string) *AssetResource {
	asset, ok := al.assets[name]

	if !ok {
		return nil
	}

	return &asset
}

// Remove removes an asset from the asset loader.
func (al *Loader) Remove(name string) {
	delete(al.assets, name)
}

// Clear removes all assets from the asset loader.
func (al *Loader) Clear() {
	al.assets = make(map[string]AssetResource)
}

// Count returns the number of assets in the asset loader.
func (al *Loader) Count() int {
	return len(al.assets)
}

// List returns a list of asset names in the asset loader.
func (al *Loader) List() []string {
	var list []string

	for name := range al.assets {
		list = append(list, name)
	}

	return list
}

// Has returns true if the asset loader has an asset with the given name.
func (al *Loader) Has(name string) bool {
	_, ok := al.assets[name]
	return ok
}

// Replace replaces an asset in the asset loader.
func (al *Loader) Replace(name string, asset AssetResource) {
	al.assets[name] = asset
}

// Rename renames an asset in the asset loader.
func (al *Loader) Rename(oldName, newName string) {
	al.assets[newName] = al.assets[oldName]
	delete(al.assets, oldName)
}

// Each iterates over each asset in the asset loader.
func (al *Loader) Each(fn func(name string, asset AssetResource)) {
	for name, asset := range al.assets {
		fn(name, asset)
	}
}

func (al Loader) loadAudioBuffer(path string) (*beep.Buffer, beep.Format, error) {
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

func (al Loader) loadAudioStream(path string) (beep.StreamSeekCloser, beep.Format, error) {
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

func (al Loader) loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func (al *Loader) load(name string, asset AssetResource) {
	al.assets[name] = asset
}

func (al Loader) loadFontFace(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	f, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(f, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	})

	return face, nil
}
