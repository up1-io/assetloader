package asset_loader

import (
	"testing"
)

func TestAssetLoader_LoadTexture(t *testing.T) {
	loader := NewAssetLoader()

	asset, err := loader.LoadTexture("test", "test/test-image.png")

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if asset.Type != PngImageAssetType {
		t.Errorf("Expected %s, got %s", PngImageAssetType, asset.Type)
	}

	if asset.Name != "test" {
		t.Errorf("Expected %s, got %s", "test", asset.Name)
	}

	if asset.Path != "test/test-image.png" {
		t.Errorf("Expected %s, got %s", "test/test-image.png", asset.Path)
	}

	if asset.Data == nil {
		t.Errorf("Expected not nil, got nil")
	}
}

func TestAssetLoader_Get(t *testing.T) {
	loader := NewAssetLoader()

	_, err := loader.LoadTexture("test", "test/test-image.png")

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	asset := loader.Get("test")

	if asset == nil {
		t.Errorf("Expected asset, got nil")
	}
}

func TestAssetLoader_Remove(t *testing.T) {
	loader := NewAssetLoader()

	_, err := loader.LoadTexture("test", "test/test-image.png")

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	loader.Remove("test")

	asset := loader.Get("test")

	print(asset)

	if asset != nil {
		t.Errorf("Expected nil, got %v", asset)
	}
}

func TestAssetLoader_Clear(t *testing.T) {
	loader := NewAssetLoader()

	_, err := loader.LoadTexture("test1", "test/test-image.png")

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	_, err = loader.LoadTexture("test2", "test/test-image.png")
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	loader.Clear()

	if loader.Count() != 0 {
		t.Errorf("Expected 0, got %d", loader.Count())
	}
}

func TestAssetLoader_Count(t *testing.T) {
	loader := NewAssetLoader()

	_, err := loader.LoadTexture("test1", "test/test-image.png")

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	_, err = loader.LoadTexture("test2", "test/test-image.png")
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if loader.Count() != 2 {
		t.Errorf("Expected 2, got %d", loader.Count())
	}
}

func TestAssetLoader_List(t *testing.T) {
	loader := NewAssetLoader()

	_, err := loader.LoadTexture("test1", "test/test-image.png")

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	_, err = loader.LoadTexture("test2", "test/test-image.png")
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	list := loader.List()

	if len(list) != 2 {
		t.Errorf("Expected 2, got %d", len(list))
	}
}

func TestAssetLoader_Each(t *testing.T) {
	loader := NewAssetLoader()

	_, err := loader.LoadTexture("test1", "test/test-image.png")

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	_, err = loader.LoadTexture("test2", "test/test-image.png")
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	loader.Each(func(name string, asset AssetResource) {
		if name != "test1" && name != "test2" {
			t.Errorf("Expected test1 or test2, got %s", name)
		}
	})
}

func TestAssetLoader_LoadTexture_Error(t *testing.T) {
	loader := NewAssetLoader()

	_, err := loader.LoadTexture("test", "test/notExist.png")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestAssetLoader_LoadAudio(t *testing.T) {
	loader := NewAssetLoader()

	asset, err := loader.LoadAudio("test", "test/audio.mp3")

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if asset.Type != Mp3AudioAssetType {
		t.Errorf("Expected %s, got %s", Mp3AudioAssetType, asset.Type)
	}

	if asset.Name != "test" {
		t.Errorf("Expected %s, got %s", "test", asset.Name)
	}

	if asset.Path != "test/audio.mp3" {
		t.Errorf("Expected %s, got %s", "test/test-audio.mp3", asset.Path)
	}

	if asset.Data == nil {
		t.Errorf("Expected not nil, got nil")
	}
}

func TestAssetLoader_LoadAudio_Error(t *testing.T) {
	loader := NewAssetLoader()

	_, err := loader.LoadAudio("test", "test/notExist.mp3")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestAssetLoader_LoadAudioStream(t *testing.T) {
	loader := NewAssetLoader()

	asset, err := loader.LoadAudioStream("test", "test/audio.mp3")

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if asset.Type != Mp3AudioStreamAssetType {
		t.Errorf("Expected %s, got %s", Mp3AudioStreamAssetType, asset.Type)
	}

	if asset.Name != "test" {
		t.Errorf("Expected %s, got %s", "test", asset.Name)
	}

	if asset.Path != "test/audio.mp3" {
		t.Errorf("Expected %s, got %s", "test/test-audio.mp3", asset.Path)
	}

	if asset.Data == nil {
		t.Errorf("Expected not nil, got nil")
	}
}
