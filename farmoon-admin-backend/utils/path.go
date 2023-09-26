package utils

import "os"

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func CreatePath(path ...string) error {
	for _, p := range path {
		if !IsExist(p) {
			if err := os.MkdirAll(p, os.ModePerm); err != nil {
				return err
			}
		}
	}
	return nil
}
