package clean

import (
	"strings"
)

var cleanableExtensions = map[string]struct{}{
	"log":      {},
	"DS_Store": {},
	"dmg":      {},
}

func IsExtensionCleanable(ext string) bool {
	ext, _ = strings.CutPrefix(ext, ".")
	_, isCleanable := cleanableExtensions[ext]

	return isCleanable
}
