package assetloader

import "fmt"

// ErrInvalidFileFormat is an error that is returned when the file format is invalid.
type ErrInvalidFileFormat struct {
	rawType string
}

func (e ErrInvalidFileFormat) Error() string {
	return fmt.Sprintf("invalid file format: %s", e.rawType)
}

// ErrAssetAlreadyExists is an error that is returned when the asset already exists.
type ErrAssetAlreadyExists struct {
	name string
}

func (e ErrAssetAlreadyExists) Error() string {
	return fmt.Sprintf("asset already exists: %s", e.name)
}
