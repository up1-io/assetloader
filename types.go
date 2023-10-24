package assetloader

import "strings"

// AssetType is a type that defines the type of an asset.
type AssetType string

const (
	PngImageAssetType       AssetType = "image/png"
	JpegImageAssetType      AssetType = "image/jpeg"
	Mp3AudioAssetType       AssetType = "audio/mp3"
	WavAudioAssetType       AssetType = "audio/wav"
	Mp3AudioStreamAssetType AssetType = "audio/mp3-stream"
	WavAudioStreamAssetType AssetType = "audio/wav-stream"
	TFFAssetType            AssetType = "font/ttf"
)

type AssetData interface {
	FontAsset | AudioClipAsset | AudioStreamAsset | TextureAsset
}

// AssetResource is a type that defines an asset resource.
// An asset resource is a resource that is loaded by the asset loader. It holds the data of the asset.
type AssetResource[T AssetData] struct {
	Type AssetType
	Name string
	Path string
	Data T
}

// GetAssetTypeByPath gets the asset type by the path.
func GetAssetTypeByPath(path string) (AssetType, error) {
	rawType := strings.Split(path, ".")[1]

	var assetType AssetType
	switch rawType {
	case "png":
		assetType = PngImageAssetType
	case "jpeg":
		assetType = JpegImageAssetType
	case "jpg":
		assetType = JpegImageAssetType
	case "mp3":
		assetType = Mp3AudioAssetType
	case "wav":
		assetType = WavAudioAssetType
	default:
		return "", ErrInvalidFileFormat{rawType: rawType}
	}

	return assetType, nil
}
