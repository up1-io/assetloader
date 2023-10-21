package assetloader

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTextureLoader_GetTexture(t *testing.T) {
	loader := NewTextureLoader()

	_, err := loader.LoadTexture("test", "test/test-image.png")
	if err != nil {
		t.Error(err)
	}

	_, ok := loader.GetTexture("test")
	if !ok {
		t.Error("texture not found")
	}
}

func TestTextureLoader_GetTexture_NotFound(t *testing.T) {
	loader := NewTextureLoader()

	_, ok := loader.GetTexture("test")
	if ok {
		t.Error("texture found")
	}
}

func TestTextureLoader_LoadTexture(t *testing.T) {
	loader := NewTextureLoader()

	_, err := loader.LoadTexture("test", "test/test-image.png")
	if err != nil {
		t.Error(err)
	}
}

func TestTextureLoader_LoadTexture_InvalidFileFormat(t *testing.T) {
	loader := NewTextureLoader()

	_, err := loader.LoadTexture("test", "test/test-image.gif")
	if reflect.TypeOf(err) != reflect.TypeOf(ErrInvalidFileFormat{}) {
		t.Error("invalid error type")
	}
}

func TestTextureLoader_LoadTexture_FileNotFound(t *testing.T) {
	loader := NewTextureLoader()

	_, err := loader.LoadTexture("test", "test/test-image-2.png")
	if err == nil {
		t.Error("error is nil")
	}
}

func TestTextureLoader_LoadTexture_AlreadyExists(t *testing.T) {
	loader := NewTextureLoader()

	_, err := loader.LoadTexture("test", "test/test-image.png")
	if err != nil {
		t.Error(err)
	}

	_, err = loader.LoadTexture("test", "test/test-image.png")
	if err == nil {
		t.Error("error is nil")
	}
}

func TestTextureLoader_loadPicture(t *testing.T) {
	loader := NewTextureLoader()

	_, err := loader.loadPicture("test/test-image.png")
	if err != nil {
		t.Error(err)
	}
}

func TestTextureLoader_loadPicture_FileNotFound(t *testing.T) {
	loader := NewTextureLoader()

	_, err := loader.loadPicture("test/test-image-2.png")
	if err == nil {
		t.Error("error is nil")
	}
}

func TestTextureLoader_GetTextures(t *testing.T) {
	loader := NewTextureLoader()

	for i := 0; i < 5; i++ {
		_, err := loader.LoadTexture(fmt.Sprintf("test%v", i), "test/test-image.png")
		if err != nil {
			t.Error(err)
		}
	}

	textures := loader.GetTextures()
	if len(textures) != 5 {
		t.Error("invalid texture count")
	}
}

func TestTextureLoader_GetTextures_Empty(t *testing.T) {
	loader := NewTextureLoader()

	textures := loader.GetTextures()
	if len(textures) != 0 {
		t.Error("invalid texture count")
	}
}

func TestTextureLoader_RemoveTexture(t *testing.T) {
	loader := NewTextureLoader()

	_, err := loader.LoadTexture("test", "test/test-image.png")
	if err != nil {
		t.Error(err)
	}

	loader.RemoveTexture("test")

	_, ok := loader.GetTexture("test")
	if ok {
		t.Error("texture found")
	}
}

func TestTextureLoader_EachTexture(t *testing.T) {
	loader := NewTextureLoader()

	for i := 0; i < 5; i++ {
		_, err := loader.LoadTexture(fmt.Sprintf("test%v", i), "test/test-image.png")
		if err != nil {
			t.Error(err)
		}
	}

	var count int
	loader.EachTexture(func(name string, texture AssetResource[TextureAsset]) {
		count++
	})

	if count != 5 {
		t.Error("invalid texture count")
	}
}
