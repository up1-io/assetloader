package assetloader

import (
	_ "image/png"
)

type Loader struct {
	FontLoader
	TextureLoader
	AudioLoader
}

// NewLoader creates a new asset loader.
func NewLoader() *Loader {
	return &Loader{
		FontLoader:    NewFontLoader(),
		TextureLoader: NewTextureLoader(),
		AudioLoader:   NewAudioLoader(),
	}
}
