package path_size

import (
	"os"
	"path/filepath"
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
		if !all {
			hdn, err := isHidden(e.Name())
			if err != nil {
				return 0, err
			}

			if hdn {
				continue
			}
		}

		if e.IsDir() {
			if rcrcv {
				newPath := filepath.Join(path, e.Name())
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
