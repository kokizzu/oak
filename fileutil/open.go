package fileutil

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/oakmound/oak/v3/dlog"
)

var (
	// BindataFn is a function to access binary data outside of os.Open
	BindataFn func(string) ([]byte, error)
	// BindataDir is a function to access directory representations alike to ioutil.ReadDir
	BindataDir func(string) ([]string, error)
	wd, _      = os.Getwd()
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}

// Open is a wrapper around os.Open that will also check BindataFn to access
// embedded data. The intended use is to use the an embedding library to create an
// Asset function that matches this signature.
func Open(file string) (io.ReadCloser, error) {
	var err error
	var rel string
	var data []byte
	// Check bindata
	if BindataFn != nil {
		// It looks like we need to clean this output sometimes--
		// we get capitalization where we don't want it occasionally?
		rel, err = filepath.Rel(wd, file)
		if err != nil {
			// Just try the relative path by itself if we can't form
			// an absolute path.
			rel = file
		}
		data, err = BindataFn(rel)
		if err == nil {
			dlog.Verb("Found file in binary,", rel)
			// convert data to io.Reader
			return nopCloser{bytes.NewReader(data)}, err
		}
		dlog.Info("File not found in binary", rel)
	}
	return os.Open(file)
}

// ReadFile replaces ioutil.ReadFile, trying to use the BindataFn if it exists.
func ReadFile(file string) ([]byte, error) {
	if BindataFn != nil {
		rel, err := filepath.Rel(wd, file)
		if err != nil {
			rel = file
		}
		return BindataFn(rel)
	}
	return ioutil.ReadFile(file)
}

// ReadDir replaces ioutil.ReadDir, trying to use the BinaryDir if it exists.
func ReadDir(file string) ([]os.FileInfo, error) {
	var fis []os.FileInfo
	var err error
	var rel string
	var strs []string
	if BindataDir != nil {
		dlog.Verb("Bindata not nil, reading directory", file)
		rel, err = filepath.Rel(wd, file)
		if err != nil {
			// Just try the relative path by itself if we can't form
			// an absolute path.
			rel = file
		}
		strs, err = BindataDir(rel)
		if err == nil {
			fis = make([]os.FileInfo, len(strs))
			for i, s := range strs {
				// If the data does not contain a period, we consider it
				// a directory
				// todo: match embed / fs packages to remove the above bad assumption
				fis[i] = dummyfileinfo{s, !strings.ContainsRune(s, '.')}
				dlog.Verb("Creating dummy file into for", s, fis[i])
			}
			return fis, nil
		}
		dlog.Error(err)
	}
	return ioutil.ReadDir(file)
}
