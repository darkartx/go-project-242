package code

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path, recursive, all)

	if err != nil {
		return "", err
	}

	return FormatSize(size, human), nil
}

func GetSize(path string, recursive, all bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	var totalSize int64 = 0

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, e := range entries {
		if !all {
			if isHidden(e.Name()) {
				continue
			}
		}

		if e.IsDir() {
			if recursive {
				newPath := filepath.Join(path, e.Name())
				size, err := GetSize(newPath, recursive, all)
				if err != nil {
					return 0, err
				}

				totalSize += size
			} else {
				continue
			}
		} else {
			fi, err := e.Info()
			if err != nil {
				return 0, err
			}

			totalSize += fi.Size()
		}
	}

	return totalSize, nil
}

func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}
	div, exp := uint64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f%cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func isHidden(filename string) bool {
	return filename[0] == '.'
}
