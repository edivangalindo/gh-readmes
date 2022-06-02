package main

import "os"

// Save a file to the filesystem
func SaveFile(path string, filename string, data []byte) error {

	// If the directory doesn't exist, create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)

		if err != nil {
			return err
		}
	}

	dest, err := os.Create(path + "/" + filename)

	if err != nil {
		return err
	}

	defer dest.Close()

	_, err = dest.Write(data)

	if err != nil {
		return err
	}

	return nil
}
