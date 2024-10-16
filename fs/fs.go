// Package fs has useful functions for file and directories management that
// I personally miss in Go standard library.
package fs

import (
	"fmt"
	"io"
	"os"
	"path"
)

// Check wether the given file or directory exists. Since file checking may
// return other errors apart from os.IsNotExist, this function propagates those
// errors. It may seem pointless to write a function that checks the os.IsNotExist
// error for you but leaves the rest of possible errors for you to handle, but if
// in your particular case, any other errors (usually lack of privileges or hardware
// drive errors) cause the program to panic, this function can be combined with the
// [yagul.g.Unwrap] function like this:
//
//	import (
//		"github.com/n-nou/yagul/g"
//		"github.com/n-mou/yagul/fs"
//	)
//
//	fileExists := g.Unwrap(fs.Exists("file"))
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CopyFile copies source file to the dest path. It relies in [io.Copy] in the
// background so it works with files of any size (it doesn't load all of the
// file bytes at once in memory so it won't crash with files that weight several
// GBs). If source is a directory, CopyFile returns an [os.ErrInvalid] error and if
// dest exists, it returns an [os.ErrExist] error instead of overwriting it. Any
// other errors returned by the functions used inside will also be propagated.
func CopyFile(source, dest string) (int64, error) {
	srcInfo, err := os.Stat(source)
	if err != nil {
		return 0, err
	}
	if srcInfo.IsDir() {
		return 0, fmt.Errorf("%s is a directory: %w", source, os.ErrInvalid)
	}

	srcFile, err := os.Open(source)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	dstExists, err := Exists(dest)
	if err != nil {
		return 0, err
	}
	if dstExists {
		return 0, fmt.Errorf("%s exists and will not be replaced: %w", dest, os.ErrExist)
	}

	dstFile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	n, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return 0, err
	}

	return n, nil
}

// CopyDir copies source directory (and all it's contents) to dest. Currently
// it doesn't support content merging (if dest exists CopyDir will return an
// [os.IsExist] error instead of trying to merge it's contents). If source is
// a file, it returns an [os.ErrInvalid] error. Any other errors returned by
// the functions used inside are also propagated.
func CopyDir(source, dest string) error {
	srcStat, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !srcStat.IsDir() {
		return fmt.Errorf("%v is not a directory: %w", source, os.ErrInvalid)
	}

	_, err = os.Stat(dest)
	if err == nil {
		return fmt.Errorf("%v exists and will not be replaced: %w", source, os.ErrExist)
	} else {
		if !os.IsNotExist(err) {
			return err
		}
	}

	entries, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, i := range entries {
		srcFile := path.Join(source, i.Name())
		dstFile := path.Join(dest, i.Name())

		if i.IsDir() {
			err := CopyDir(srcFile, dstFile)
			if err != nil {
				return err
			}
		} else {
			_, err := CopyFile(srcFile, dstFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
