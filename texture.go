package assetloader

import (
	"github.com/gopxl/pixel"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
)

// TextureAsset is a type that defines a texture asset. It holds the picture.
type TextureAsset struct {
	Picture pixel.Picture
}

type TextureLoader struct {
	assets map[string]AssetResource[TextureAsset]
}

func NewTextureLoader() TextureLoader {
	return TextureLoader{
		assets: make(map[string]AssetResource[TextureAsset]),
	}
}

// LoadTexture loads a texture asset into the asset loader.
func (tl *TextureLoader) LoadTexture(name string, path string) (AssetResource[TextureAsset], error) {
	if _, ok := tl.assets[name]; ok {
		return AssetResource[TextureAsset]{}, ErrAssetAlreadyExists{name: name}
	}

	pic, err := tl.loadPicture(path)
	if err != nil {
		return AssetResource[TextureAsset]{}, err
	}

	// Get the file extension.
	rawType := strings.Split(path, ".")[1]

	var assetType AssetType
	switch rawType {
	case "png":
		assetType = PngImageAssetType
	case "jpeg":
		assetType = JpegImageAssetType
	case "jpg":
		assetType = JpegImageAssetType
	default:
		return AssetResource[TextureAsset]{}, ErrInvalidFileFormat{rawType: rawType}
	}

	asset := AssetResource[TextureAsset]{
		Type: assetType,
		Name: name,
		Path: path,
		Data: TextureAsset{
			Picture: pic,
		},
	}

	tl.assets[name] = asset

	return asset, nil
}

// GetTexture gets a texture asset from the asset loader.
func (tl *TextureLoader) GetTexture(name string) (AssetResource[TextureAsset], bool) {
	asset, ok := tl.assets[name]
	return asset, ok
}

// GetTextures gets all texture assets from the asset loader.
func (tl *TextureLoader) GetTextures() map[string]AssetResource[TextureAsset] {
	return tl.assets
}

// RemoveTexture removes a texture asset from the asset loader.
func (tl *TextureLoader) RemoveTexture(name string) {
	delete(tl.assets, name)
}

// EachTexture iterates over all texture assets in the asset loader.
func (tl *TextureLoader) EachTexture(fn func(name string, asset AssetResource[TextureAsset])) {
	for name, asset := range tl.assets {
		fn(name, asset)
	}
}

func (tl *TextureLoader) loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}
