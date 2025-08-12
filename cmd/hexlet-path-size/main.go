package main

import (
	"context"
	"log"
	"os"

	cli "github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Usage: "print size of a file or directory",
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// flag.Parse()

	// if flag.NArg() < 1 {
	// 	fmt.Fprintln(os.Stderr, "Error: path is required")
	// 	os.Exit(1)
	// }

	// path := flag.Arg(0)

	// size, err := getSize(path, flagRecursive, flagAll)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error:", err)
	// 	os.Exit(0)
	// }

	// var sizeStr string

	// if flagHumanReadable {
	// 	sizeStr = humanize.Bytes(uint64(size))
	// } else {
	// 	sizeStr = fmt.Sprint(size)
	// }

	// fmt.Println(sizeStr, path)
}

// Собираем размер файлов по пути
// Если путь - файл, возвращаем его размер
// Если путь - папка, возвращаем размер файлов находящихся в ней
//
//	если передан флаг rcrcv рекурсивно считаем размер файлов внутри папок в этой папке
//	если передан флаг all у скрытых файлов тоже будет считаться размер
// func getSize(path string, rcrcv, all bool) (int64, error) {
// 	info, err := os.Lstat(path)
// 	if err != nil {
// 		return 0, err
// 	}

// 	if !info.IsDir() {
// 		return info.Size(), nil
// 	}

// 	var totalSize int64 = 0

// 	entries, err := os.ReadDir(path)
// 	if err != nil {
// 		return 0, err
// 	}

// 	for _, e := range entries {
// 		if !all {
// 			hdn, err := isHidden(e.Name())
// 			if err != nil {
// 				return 0, err
// 			}

// 			if hdn {
// 				continue
// 			}
// 		}

// 		if e.IsDir() {
// 			if rcrcv {
// 				newPath := filepath.Join(path, e.Name())
// 				size, err := getSize(newPath, rcrcv, all)
// 				if err != nil {
// 					return 0, err
// 				}

// 				totalSize += size
// 			} else {
// 				continue
// 			}
// 		} else {
// 			fi, err := e.Info()
// 			if err != nil {
// 				return 0, err
// 			}

// 			totalSize += fi.Size()
// 		}
// 	}

// 	return totalSize, nil
// }
