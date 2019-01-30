package git

import (
	"os"
	"path/filepath"
)

func IsRepo(root string) bool {
	fi, err := os.Stat(filepath.Join(root, ".git"))
	if err != nil {
		return false
	}
	return fi.IsDir()
}
