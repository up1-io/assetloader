package assetloader

import (
	"github.com/golang/freetype/truetype"
	"github.com/gopxl/pixel/text"
	"golang.org/x/image/font"
	"io"
	"os"
)

// FontAsset is a type that defines a font asset. It holds the font face and atlas.
type FontAsset struct {
	Face  font.Face
	Atlas *text.Atlas
}

type FontLoader struct {
	assets map[string]AssetResource[FontAsset]
}

func NewFontLoader() FontLoader {
	return FontLoader{
		assets: make(map[string]AssetResource[FontAsset]),
	}
}

// LoadFont loads a font asset into the asset loader.
func (fl *FontLoader) LoadFont(name, path string, options *truetype.Options, runeSets ...[]rune) (AssetResource[FontAsset], error) {
	asset, err := fl.loadFont(path, options, runeSets...)
	if err != nil {
		return AssetResource[FontAsset]{}, err
	}

	res := AssetResource[FontAsset]{
		Type: TFFAssetType,
		Name: name,
		Path: path,
		Data: asset,
	}

	fl.assets[name] = res

	return res, nil
}

// GetFont gets a font asset from the asset loader.
func (fl *FontLoader) GetFont(name string) (AssetResource[FontAsset], bool) {
	asset, ok := fl.assets[name]
	return asset, ok
}

// GetFonts gets all font assets from the asset loader.
func (fl *FontLoader) GetFonts() map[string]AssetResource[FontAsset] {
	return fl.assets
}

// RemoveFont removes a font asset from the asset loader.
func (fl *FontLoader) RemoveFont(name string) {
	delete(fl.assets, name)
}

// EachFont iterates over all font assets in the asset loader.
func (fl *FontLoader) EachFont(fn func(name string, asset AssetResource[FontAsset])) {
	for name, asset := range fl.assets {
		fn(name, asset)
	}
}

func (fl *FontLoader) loadFont(path string, options *truetype.Options, runeSets ...[]rune) (FontAsset, error) {
	face, err := fl.loadFontFace(path, options)
	if err != nil {
		return FontAsset{}, err
	}

	// Default to ASCII if no runes are provided.
	if len(runeSets) == 0 {
		runeSets = append(runeSets, text.ASCII)
	}

	atlas := text.NewAtlas(face, runeSets...)

	return FontAsset{
		Face:  face,
		Atlas: atlas,
	}, nil
}

func (fl *FontLoader) loadFontFace(path string, options *truetype.Options) (font.Face, error) {
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

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	f, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(f, options)

	return face, nil
}
