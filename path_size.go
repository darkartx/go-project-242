package code

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
)

// Собираем размер файлов по пути
// Если путь - файл, возвращаем его размер
// Если путь - папка, возвращаем размер файлов находящихся в ней
//
//	если передан флаг rcrcv рекурсивно считаем размер файлов внутри папок в этой папке
//	если передан флаг all у скрытых файлов тоже будет считаться размер
func GetSize(path string, rcrcv, all bool) (int64, error) {
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
		newPath := filepath.Join(path, e.Name())

		if !all {
			hdn, err := isHidden(newPath)
			if err != nil {
				return 0, err
			}

			if hdn {
				continue
			}
		}

		if e.IsDir() {
			if rcrcv {
				size, err := GetSize(newPath, rcrcv, all)
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
	if human {
		return humanize.Bytes(uint64(size))
	} else {
		return fmt.Sprint(size, " B")
	}
}
