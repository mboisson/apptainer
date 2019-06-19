// Copyright (c) 2018-2019, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package cache

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sylabs/scs-library-client/client"
)

const (
	// LibraryDir is the directory inside the cache.Dir where library images are cached
	LibraryDir = "library"
)

/*
// Library returns the directory inside the cache.Dir() where library images are cached
func Library() string {
	return updateCacheSubdir(LibraryDir)
}
*/

// Library returns the directory inside the cache.Dir() where library
// images are cached
func getLibraryCachePath(c *ImgCache) (string, error) {
	// This function may act on an cache object that is not fully
	// initialized so it is not a method on a ImgCache but
	// rather an independent function.

	return updateCacheSubdir(c, LibraryDir)
}

// LibraryImage creates a directory inside cache.Dir() with the name of the SHA sum of the image
func (c *ImgCache) LibraryImage(sum, name string) string {
	_, err := updateCacheSubdir(c, filepath.Join(LibraryDir, sum))
	if err != nil {
		return ""
	}

	return filepath.Join(c.Library, sum, name)
}

// LibraryImageExists returns whether the image with the SHA sum exists in the LibraryImage cache
func (c *ImgCache) LibraryImageExists(sum, name string) (bool, error) {
	imagePath := c.LibraryImage(sum, name)
	_, err := os.Stat(imagePath)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	cacheSum, err := client.ImageHash(imagePath)
	if err != nil {
		return false, err
	}
	if cacheSum != sum {
		return false, fmt.Errorf("cached file sum (%s) and expected sum (%s) does not match", cacheSum, sum)
	}

	return true, nil
}
