package assetloader

import (
	"fmt"
	"testing"
)

func TestAudioLoader_LoadAudioClip(t *testing.T) {
	loader := NewAudioLoader()

	_, err := loader.LoadAudioClip("test", "test/audio.mp3")
	if err != nil {
		t.Error(err)
	}
}

func TestAudioLoader_GetAudioClip(t *testing.T) {
	loader := NewAudioLoader()

	_, err := loader.LoadAudioClip("test", "test/audio.mp3")
	if err != nil {
		t.Error(err)
	}

	_, ok := loader.GetAudioClip("test")
	if !ok {
		t.Error("audio clip not found")
	}
}

func TestAudioLoader_GetAudioClip_NotFound(t *testing.T) {
	loader := NewAudioLoader()

	_, ok := loader.GetAudioClip("test")
	if ok {
		t.Error("audio clip found")
	}
}

func TestAudioLoader_GetAudioClips(t *testing.T) {
	loader := NewAudioLoader()

	for i := 0; i < 5; i++ {
		_, err := loader.LoadAudioClip(fmt.Sprintf("test%v", i), "test/audio.mp3")
		if err != nil {
			t.Error(err)
		}
	}

	clips := loader.GetAudioClips()
	if len(clips) != 5 {
		t.Error("invalid audio clip count")
	}
}

func TestAudioLoader_GetAudioClips_Empty(t *testing.T) {
	loader := NewAudioLoader()

	clips := loader.GetAudioClips()
	if len(clips) != 0 {
		t.Error("invalid audio clip count")
	}
}

func TestAudioLoader_RemoveAudioClip(t *testing.T) {
	loader := NewAudioLoader()

	_, err := loader.LoadAudioClip("test", "test/audio.mp3")
	if err != nil {
		t.Error(err)
	}

	loader.RemoveAudioClip("test")

	_, ok := loader.GetAudioClip("test")
	if ok {
		t.Error("audio clip found")
	}
}

func TestAudioLoader_LoadAudioStream(t *testing.T) {
	loader := NewAudioLoader()

	_, err := loader.LoadAudioStream("test", "test/audio.mp3")
	if err != nil {
		t.Error(err)
	}
}

func TestAudioLoader_GetAudioStream(t *testing.T) {
	loader := NewAudioLoader()

	_, err := loader.LoadAudioStream("test", "test/audio.mp3")
	if err != nil {
		t.Error(err)
	}

	_, ok := loader.GetAudioStream("test")
	if !ok {
		t.Error("audio stream not found")
	}
}

func TestAudioLoader_GetAudioStream_NotFound(t *testing.T) {
	loader := NewAudioLoader()

	_, ok := loader.GetAudioStream("test")
	if ok {
		t.Error("audio stream found")
	}
}

func TestAudioLoader_GetAudioStreams(t *testing.T) {
	loader := NewAudioLoader()

	for i := 0; i < 5; i++ {
		_, err := loader.LoadAudioStream(fmt.Sprintf("test%v", i), "test/audio.mp3")
		if err != nil {
			t.Error(err)
		}
	}

	streams := loader.GetAudioStreams()
	if len(streams) != 5 {
		t.Error("invalid audio stream count")
	}
}
