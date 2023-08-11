package fileutil

import "path/filepath"

func GetVideoFileExt(s string) (string, bool) {
	ext := filepath.Ext(s)
	switch ext {
	case ".mp4", ".avi", ".wmv", ".flv", ".mpeg", ".mov":
		return ext[1:], true
	default:
		return "", false
	}
}
