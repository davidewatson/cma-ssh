// Package asset provides the assets to a virtual filesystem.
package homepage

import (
	_ "github.com/shurcooL/vfsgen"
)

//go:generate go run -tags=dev asset_generate.go
