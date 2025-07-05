package clean

import (
	"io/fs"
	"path/filepath"
	"strings"
)

var cleanableExtensions = map[string]struct{}{
	"log":      {},
	"DS_Store": {},
	"dmg":      {},
	"cache":    {},
}

func IsExtensionCleanable(info fs.FileInfo) bool {
	ext := filepath.Ext(info.Name())
	ext, _ = strings.CutPrefix(ext, ".")
	_, isCleanable := cleanableExtensions[ext]

	return isCleanable
}
