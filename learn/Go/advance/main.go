package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	FirecrackerVersionsDir = "/fc-versions"
	FirecrackerBinaryName  = "firecracker"
)

func main() {
	version := "v1.10.1_1fcdaec"

	path := filepath.Join(FirecrackerVersionsDir, version, FirecrackerBinaryName)

	_, err := os.Stat(path)
	if err != nil {
		fmt.Printf("stat error: %v\n", err)
	}
}
