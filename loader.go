package asset_loader

import (
	"github.com/gopxl/pixel"
	"image"
	_ "image/png"
	"os"
)

type AssetLoader struct {
	assets map[string]AssetResource
}

// IAssetLoader is an interface that defines the methods that an asset loader must implement.
type IAssetLoader interface {
	LoadTexture(name string, path string) (AssetResource, error)
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

// NewAssetLoader creates a new asset loader.
func NewAssetLoader() *AssetLoader {
	return &AssetLoader{
		assets: make(map[string]AssetResource),
	}
}

// LoadTexture loads a texture asset into the asset loader.
func (al AssetLoader) LoadTexture(name string, path string) (AssetResource, error) {
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

// Get returns an asset from the asset loader.
func (al *AssetLoader) Get(name string) *AssetResource {
	asset, ok := al.assets[name]

	if !ok {
		return nil
	}

	return &asset
}

// Remove removes an asset from the asset loader.
func (al *AssetLoader) Remove(name string) {
	delete(al.assets, name)
}

// Clear removes all assets from the asset loader.
func (al *AssetLoader) Clear() {
	al.assets = make(map[string]AssetResource)
}

// Count returns the number of assets in the asset loader.
func (al *AssetLoader) Count() int {
	return len(al.assets)
}

// List returns a list of asset names in the asset loader.
func (al *AssetLoader) List() []string {
	var list []string

	for name := range al.assets {
		list = append(list, name)
	}

	return list
}

// Has returns true if the asset loader has an asset with the given name.
func (al *AssetLoader) Has(name string) bool {
	_, ok := al.assets[name]
	return ok
}

// Replace replaces an asset in the asset loader.
func (al *AssetLoader) Replace(name string, asset AssetResource) {
	al.assets[name] = asset
}

// Rename renames an asset in the asset loader.
func (al *AssetLoader) Rename(oldName, newName string) {
	al.assets[newName] = al.assets[oldName]
	delete(al.assets, oldName)
}

// Each iterates over each asset in the asset loader.
func (al *AssetLoader) Each(fn func(name string, asset AssetResource)) {
	for name, asset := range al.assets {
		fn(name, asset)
	}
}

func (al AssetLoader) loadPicture(path string) (pixel.Picture, error) {
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

func (al *AssetLoader) load(name string, asset AssetResource) {
	al.assets[name] = asset
}
