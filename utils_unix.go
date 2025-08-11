//go:build !windows

package main

func isHidden(filename string) (bool, error) {
	return filename[0] == '.', nil
}
