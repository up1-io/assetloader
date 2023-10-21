# Asset Loader

![Version](https://img.shields.io/badge/Version-Prototype-red)
[![GoDoc](https://godoc.org/github.com/up1-io/ecs?status.svg)](https://godoc.org/github.com/up1-io/asset-loader)
[![Go](https://github.com/up1-io/asset-loader/actions/workflows/go.yml/badge.svg)](https://github.com/up1-io/asset-loader/actions/workflows/go.yml)

> Note: This library is currently in prototype stage. It is not recommended to use this library in production.

Go Asset Loader is a straightforward asset loading library for the Go programming language. 
It is built on top of the [Pixel](https://github.com/gopxl/pixel) and [Beep](https://github.com/gopxl/beep) packages. 

This library has been developed as part of the 2D Game Engine project called [Raindrop](https://github.com/up1-io/raindrop).

## Features

Asset Loader supports the following features:

- **Textures**: Load textures from PNG and JPEG files.
- **Sounds**: Load sounds from MP3 files.
- **Fonts**: Load fonts from TTF files.

## Getting Started

1. **Installation**: Get the package using `go get`:

```bash
   go get github.com/up1-io/asset-loader
```

2. **Usage**: Import the package in your code and start using it:

```go
   import "github.com/up1-io/asset-loader"
```

3. **Simple Example**:

```go
package main

func main() {
	loader := asset_loader.NewLoader()

	asset, err := loader.LoadTexture("test", "test/test-image.png")
	if err != nil {
		panic(err)
	}

	println(asset.Name, asset.Path, asset.Type)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

