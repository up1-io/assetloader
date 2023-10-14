# Asset Loader


![Version](https://img.shields.io/badge/Version-Prototype-red)
[![GoDoc](https://godoc.org/github.com/up1-io/ecs?status.svg)](https://godoc.org/github.com/up1-io/asset-loader)

> Note: This library is currently in prototype stage. It is not recommended to use this library in production.

A simple asset loader in Go. It can be used to load any type of assets, such as images, sounds, fonts, etc.

It's developed for the 2D Game Engine project [Raindrop](https://github.com/up1-io/raindrop).

## Getting Started

1. **Installation**: Get the package using `go get`:

```bash
   go get github.com/up1-io/asset-loader
```

2. **Usage**: Import the package in your code and start using it:

```go
   import "github.com/up1-io/asset-loader"
```

3. **Examples**:

```go
package main

func main() {
	loader := NewAssetLoader()
	asset, err := loader.LoadTexture("test", "test/test-image.png")
	if err != nil {
		panic(err)
	}
	
	// Use the asset
	println(asset.Name, asset.Path, asset.Type)
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

