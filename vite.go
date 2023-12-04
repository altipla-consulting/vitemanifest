package vitemanifest

import (
	"encoding/json"
	"fmt"
	"os"
)

// File stores the structure of the whole manifest file.
type File map[string]Chunk

// Chunk is every entry in the manifest file.
type Chunk struct {
	File           string
	CSS            []string
	Assets         []string
	IsDynamicEntry bool
	IsEntry        bool
	DynamicImports []string
	Imports        []string
}

// Read the manifest file from its standard location of `dist/.vite/manifest.json`.
func Read() (File, error) {
	return ReadPath("dist/.vite/manifest.json")
}

// ReadPath reads the manifest file from a given path.
func ReadPath(path string) (File, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return File{}, nil
		}
		return nil, fmt.Errorf("vitemanifest: cannot open manifest: %w", err)
	}
	defer f.Close()

	manifest := make(File)
	if err := json.NewDecoder(f).Decode(&manifest); err != nil {
		return nil, fmt.Errorf("vitemanifest: cannot decode manifest: %w", err)
	}

	return manifest, nil
}
