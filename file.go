package utils
import (
	"os"
	"fmt"
	"io"
)

/*
Warning: this function will work correctly if there is priviledge to access 'path'
*/
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false

	}
	return true
}

/*
https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
*/
func FileExistsEx(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err != nil, err
}

/*
https://stackoverflow.com/questions/50740902/move-a-file-to-a-different-drive-with-go
1. os.Rename(source, destination) will not work for different Driver
2. use io.Copy() rather than ioutil.ReadFile(), since Using ioutil.ReadFile is generally fine for smaller files,
   but it does copy the entire file into a slice.
   This means means that "moving" a large file will take lots of memory if you use those functions.
   io.Copy can even make use of kernel-specific syscalls that avoid loading the file into userspace at all. You should use io.Copy.
   ioutil. ReadFile` reads the entire file into memory at one time. It's only good for small filles.
*/
func FileMove(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}
