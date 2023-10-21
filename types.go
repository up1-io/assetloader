package assetloader

// AssetType is a type that defines the type of an asset.
type AssetType string

const (
	PngImageAssetType       AssetType = "image/png"
	JpegImageAssetType      AssetType = "image/jpeg"
	Mp3AudioAssetType       AssetType = "audio/mp3"
	Mp3AudioStreamAssetType AssetType = "audio/mp3-stream"
	TFFAssetType            AssetType = "font/ttf"
)

type AssetData interface {
	FontAsset | AudioAsset | AudioStreamAsset | TextureAsset
}

// AssetResource is a type that defines an asset resource.
// An asset resource is a resource that is loaded by the asset loader. It holds the data of the asset.
type AssetResource[T AssetData] struct {
	Type AssetType
	Name string
	Path string
	Data T

	// IsDirty is a flag that indicates if the asset resource has been modified.
	IsDirty bool
}
