package cache

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
)

func DiskSave(key string, data []byte, dir string) {
	os.MkdirAll(dir, 0755)
	path := filepath.Join(dir, hash(key)+".jpg")
	os.WriteFile(path, data, 0644)
}

func Key(src string, params map[string]string) string {
	return fmt.Sprintf("%s|%v", src, params)
}

func hash(s string) string {
	h := sha1.Sum([]byte(s))
	return fmt.Sprintf("%x", h)
}
