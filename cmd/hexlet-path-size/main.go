package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	humanize "github.com/dustin/go-humanize"
)

var (
	flagRecursive     bool
	flagHumanReadable bool
	flagAll           bool
)

func init() {
	// Инициализация флагов и сообщения помощи
	flag.BoolVar(&flagRecursive, "r", false, "recursive size of directories")
	flag.BoolVar(&flagRecursive, "recursive", false, "recursive size of directories")
	flag.BoolVar(&flagHumanReadable, "H", false, "human-readable sizes (auto-select unit)")
	flag.BoolVar(&flagHumanReadable, "human", false, "human-readable sizes (auto-select unit)")
	flag.BoolVar(&flagAll, "a", false, "include hidden files and directories")
	flag.BoolVar(&flagAll, "all", false, "include hidden files and directories")

	flag.Usage = func() {
		executable, _ := os.Executable()
		executable = filepath.Base(executable)

		fmt.Println("NAME:")
		fmt.Println("\t", executable)
		fmt.Println()
		fmt.Println("USAGE:")
		fmt.Println("\t", executable, "[global options] path")
		fmt.Println()
		fmt.Println("GLOBAL OPTIONS:")

		flags := map[string][2]string{}
		flag.VisitAll(func(f *flag.Flag) {
			var name string
			if len(f.Name) > 1 {
				name = fmt.Sprintf("--%s", f.Name)
			} else {
				name = fmt.Sprintf("-%s", f.Name)
			}

			if val, ok := flags[f.Usage]; ok {
				name = fmt.Sprintf("%s, %s", val[0], name)
				flags[f.Usage] = [2]string{name, val[1]}
				return
			}

			usage := fmt.Sprintf("%s (default: %s)", f.Usage, f.DefValue)
			flags[f.Usage] = [2]string{name, usage}
		})

		flags[""] = [2]string{"-h, --help", "show help"}
		for _, f := range flags {
			fmt.Println("\t", f[0], "\t", f[1])
		}
	}
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Error: path is required")
		os.Exit(1)
	}

	path := flag.Arg(0)

	size, err := getSize(path, flagRecursive, flagAll)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(0)
	}

	var sizeStr string

	if flagHumanReadable {
		sizeStr = humanize.Bytes(uint64(size))
	} else {
		sizeStr = fmt.Sprint(size)
	}

	fmt.Println(sizeStr, path)
}

// Собираем размер файлов по пути
// Если путь - файл, возвращаем его размер
// Если путь - папка, возвращаем размер файлов находящихся в ней
//
//	если передан флаг rcrcv рекурсивно считаем размер файлов внутри папок в этой папке
//	если передан флаг all у скрытых файлов тоже будет считаться размер
func getSize(path string, rcrcv, all bool) (int64, error) {
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
				size, err := getSize(newPath, rcrcv, all)
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
