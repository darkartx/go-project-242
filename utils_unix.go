//go:build !windows

package path_size

func isHidden(filename string) (bool, error) {
	return filename[0] == '.', nil
}
