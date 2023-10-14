package asset_loader

// AssetResource is a type that defines an asset resource.
// An asset resource is a resource that is loaded by the asset loader. It holds the data of the asset.
type AssetResource struct {
	Type AssetType
	Name string
	Path string
	Data interface{}

	// IsDirty is a flag that indicates if the asset resource has been modified.
	IsDirty bool
}
