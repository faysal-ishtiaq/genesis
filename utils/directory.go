package utils

import "os"

// MakeDir creates directory at an absolute path. Returns an error if failed
func MakeDir(absolutePath string) error {
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		if err := os.Mkdir(absolutePath, 0754); err != nil {
			return err
		}
	}

	return nil
}
