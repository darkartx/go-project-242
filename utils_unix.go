//go:build !windows

package code

func isHidden(filename string) (bool, error) {
	return filename[0] == '.', nil
}
