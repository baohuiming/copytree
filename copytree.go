package copytree

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Copy file from srcPath to dstPath. If dstPath already exists, it will be overwritten.
func CopyFile(srcPath, dstPath string) error {
	inputFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("error open source file: %v", err)
	}
	defer inputFile.Close()

	fileInfo, err := inputFile.Stat()
	if err != nil {
		return fmt.Errorf("error get file info: %v", err)
	}

	dir := filepath.Dir(dstPath)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error create dst dir: %v", err)
	}

	outputFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("error create dst file: %v", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("error copy to dst from source: %v", err)
	}

	err = os.Chtimes(dstPath, fileInfo.ModTime(), fileInfo.ModTime())
	if err != nil {
		return fmt.Errorf("error set modification time: %v", err)
	}

	return nil
}

// Copy files from srcPath to dstPath. If dstPath already exists, it will be overwritten.
func CopyFiles(srcPath string, dstPath string) error {
	// check if srcPath is a directory
	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		return fmt.Errorf("error get file info: %v", err)
	}
	if !srcInfo.IsDir() {
		err := CopyFile(srcPath, dstPath)
		if err != nil {
			return fmt.Errorf("error copy file: %v", err)
		}
		return nil
	}

	err = os.MkdirAll(dstPath, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("error create dst directory: %w", err)
	}

	files, err := os.ReadDir(srcPath)
	if err != nil {
		return fmt.Errorf("error read source directory: %w", err)
	}

	for _, f := range files {
		_srcPath := filepath.Join(srcPath, f.Name())
		_dstPath := filepath.Join(dstPath, f.Name())

		if f.IsDir() {
			err = CopyFiles(_srcPath, _dstPath)
			if err != nil {
				return fmt.Errorf("error move directory: %w", err)
			}
		} else {
			err = CopyFile(_srcPath, _dstPath)
			if err != nil {
				return fmt.Errorf("error copy file: %w", err)
			}
		}
	}
	return nil
}

// Move files from srcPath to dstPath, even if dstPath is in a different volume.
func MoveFiles(srcPath string, dstPath string) error {
	err := CopyFiles(srcPath, dstPath)
	if err != nil {
		return fmt.Errorf("error copy files: %w", err)
	}
	err = os.RemoveAll(srcPath)
	if err != nil {
		return fmt.Errorf("error remove source directory: %w", err)
	}

	return nil
}
