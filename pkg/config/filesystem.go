// Copyright 2022 tsuru-client authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"os"
	"path/filepath"
	"sync/atomic"

	"github.com/tsuru/tsuru/fs"
)

var (
	fsystem atomic.Pointer[fs.Fs]
)

func Filesystem() fs.Fs {
	f := fsystem.Load()
	if f == nil {
		return &fs.OsFs{}
	}
	return *f
}

func SetFileSystem(f fs.Fs) {
	fsystem.Store(&f)
}

func ResetFileSystem() {
	fsystem.Store(nil)
}

func getHome() string {
	envs := []string{"HOME", "HOMEPATH"}
	var home string
	for i := 0; i < len(envs) && home == ""; i++ {
		home = os.Getenv(envs[i])
	}
	return home
}

func JoinWithUserDir(p ...string) string {
	paths := []string{getHome()}
	paths = append(paths, p...)
	return filepath.Join(paths...)
}

// openPrivateFile opens path for writing, truncating it, and keeps it
// readable only by its owner, including files created with a wider mode.
func openPrivateFile(path string) (fs.File, error) {
	file, err := Filesystem().OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}
	if f, ok := file.(*os.File); ok {
		if err := f.Chmod(0600); err != nil {
			f.Close()
			return nil, err
		}
	}
	return file, nil
}
