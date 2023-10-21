package asset_loader

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/gopxl/pixel/text"
	"testing"
)

func TestFontLoader_LoadFont(t *testing.T) {
	loader := NewFontLoader()

	opts := truetype.Options{
		Size: 12,
	}

	_, err := loader.LoadFont("test", "test/unifont.ttf", &opts, text.ASCII)
	if err != nil {
		t.Error(err)
	}
}

func TestFontLoader_GetFont(t *testing.T) {
	loader := NewFontLoader()

	opts := truetype.Options{
		Size: 12,
	}

	_, err := loader.LoadFont("test", "test/unifont.ttf", &opts, text.ASCII)
	if err != nil {
		t.Error(err)
	}

	_, ok := loader.GetFont("test")
	if !ok {
		t.Error("font not found")
	}
}

func TestFontLoader_GetFonts(t *testing.T) {
	loader := NewFontLoader()

	opts := truetype.Options{
		Size: 12,
	}

	for i := 0; i < 5; i++ {
		_, err := loader.LoadFont(fmt.Sprintf("test%v", i), "test/unifont.ttf", &opts, text.ASCII)
		if err != nil {
			t.Error(err)
		}
	}

	fonts := loader.GetFonts()
	if len(fonts) != 5 {
		t.Error("invalid font count")
	}
}

func TestFontLoader_RemoveFont(t *testing.T) {
	loader := NewFontLoader()

	opts := truetype.Options{
		Size: 12,
	}

	_, err := loader.LoadFont("test", "test/unifont.ttf", &opts, text.ASCII)
	if err != nil {
		t.Error(err)
	}

	loader.RemoveFont("test")

	_, ok := loader.GetFont("test")
	if ok {
		t.Error("font found")
	}
}

func TestFontLoader_EachFont(t *testing.T) {
	loader := NewFontLoader()

	opts := truetype.Options{
		Size: 12,
	}

	for i := 0; i < 5; i++ {
		_, err := loader.LoadFont(fmt.Sprintf("test%v", i), "test/unifont.ttf", &opts, text.ASCII)
		if err != nil {
			t.Error(err)
		}
	}

	count := 0
	loader.EachFont(func(name string, asset AssetResource[FontAsset]) {
		count++
	})

	if count != 5 {
		t.Error("invalid font count")
	}
}
