package asset_loader

import "github.com/gopxl/beep"

// AssetType is a type that defines the type of an asset.
type AssetType string

const (
	PngImageAssetType       AssetType = "image/png"
	Mp3AudioAssetType       AssetType = "audio/mp3"
	Mp3AudioStreamAssetType AssetType = "audio/mp3-stream"
	TFFAssetType            AssetType = "font/ttf"
)

// AudioAsset is a type that defines an audio asset. It holds the audio buffer and format.
type AudioAsset struct {
	Buffer *beep.Buffer
	Format beep.Format
}

type AudioStreamAsset struct {
	Stream beep.StreamSeekCloser
	Format beep.Format
}
