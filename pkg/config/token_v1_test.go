// Copyright 2024 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/tsuru/tsuru/fs/fstest"
	"gopkg.in/check.v1"
)

func (s *S) TestWriteToken(c *check.C) {
	rfs := &fstest.RecordingFs{}
	SetFileSystem(rfs)
	defer func() {
		ResetFileSystem()
	}()
	err := WriteTokenV1("abc")
	c.Assert(err, check.IsNil)
	tokenPath := JoinWithUserDir(".tsuru", "token")
	c.Assert(rfs.HasAction("mkdirall "+JoinWithUserDir(".tsuru")+" with mode 0700"), check.Equals, true)
	c.Assert(rfs.HasAction("openfile "+tokenPath+" with mode 0600"), check.Equals, true)
	fil, _ := Filesystem().Open(tokenPath)
	b, _ := io.ReadAll(fil)
	c.Assert(string(b), check.Equals, "abc")
}

func (s *S) TestWriteTokenWithTarget(c *check.C) {
	rfs := &fstest.RecordingFs{}
	SetFileSystem(rfs)
	initTestTarget()

	defer func() {
		ResetFileSystem()
	}()
	err := WriteTokenV1("abc")
	c.Assert(err, check.IsNil)
	tokenPath1 := JoinWithUserDir(".tsuru", "token")
	c.Assert(rfs.HasAction("openfile "+tokenPath1+" with mode 0600"), check.Equals, true)
	tokenPath2 := JoinWithUserDir(".tsuru", "token.d", "test")
	c.Assert(rfs.HasAction("openfile "+tokenPath2+" with mode 0600"), check.Equals, true)
	fil, _ := Filesystem().Open(tokenPath1)
	b, _ := io.ReadAll(fil)
	c.Assert(string(b), check.Equals, "abc")
	fil, _ = Filesystem().Open(tokenPath2)
	b, _ = io.ReadAll(fil)
	c.Assert(string(b), check.Equals, "abc")
}

func (s *S) TestWriteTokenFileMode(c *check.C) {
	if runtime.GOOS == "windows" {
		c.Skip("unix file modes")
	}
	home := c.MkDir()
	defer os.Setenv("HOME", os.Getenv("HOME"))
	os.Setenv("HOME", home)
	tokenPath := filepath.Join(home, ".tsuru", "token")
	c.Assert(os.MkdirAll(filepath.Dir(tokenPath), 0755), check.IsNil)
	c.Assert(os.WriteFile(tokenPath, []byte("old"), 0644), check.IsNil)
	err := WriteTokenV1("abc")
	c.Assert(err, check.IsNil)
	info, err := os.Stat(tokenPath)
	c.Assert(err, check.IsNil)
	c.Assert(info.Mode().Perm(), check.Equals, os.FileMode(0600))
	b, err := os.ReadFile(tokenPath)
	c.Assert(err, check.IsNil)
	c.Assert(string(b), check.Equals, "abc")
}

func (s *S) TestWriteTokenCreatesTsuruDir(c *check.C) {
	if runtime.GOOS == "windows" {
		c.Skip("unix file modes")
	}
	home := c.MkDir()
	defer os.Setenv("HOME", os.Getenv("HOME"))
	os.Setenv("HOME", home)
	err := WriteTokenV1("abc")
	c.Assert(err, check.IsNil)
	info, err := os.Stat(filepath.Join(home, ".tsuru"))
	c.Assert(err, check.IsNil)
	c.Assert(info.Mode().Perm(), check.Equals, os.FileMode(0700))
	info, err = os.Stat(filepath.Join(home, ".tsuru", "token"))
	c.Assert(err, check.IsNil)
	c.Assert(info.Mode().Perm(), check.Equals, os.FileMode(0600))
}

func (s *S) TestReadToken(c *check.C) {
	os.Unsetenv("TSURU_TOKEN")
	rfs := &fstest.RecordingFs{}
	SetFileSystem(rfs)
	initTestTarget()
	f, err := Filesystem().Create(JoinWithUserDir(".tsuru", "token.d", "test"))
	c.Assert(err, check.IsNil)
	f.WriteString("mytoken")
	defer func() {
		ResetFileSystem()
	}()
	token, err := ReadTokenV1()
	c.Assert(err, check.IsNil)
	c.Assert(token, check.Equals, "mytoken")
	tokenPath := JoinWithUserDir(".tsuru", "token.d", "test")
	c.Assert(rfs.HasAction("open "+tokenPath), check.Equals, true)
	tokenPath = JoinWithUserDir(".tsuru", "token")
	c.Assert(rfs.HasAction("open "+tokenPath), check.Equals, false)
}

func (s *S) TestReadTokenFallback(c *check.C) {
	os.Unsetenv("TSURU_TOKEN")
	rfs := &fstest.RecordingFs{}
	SetFileSystem(rfs)
	defer func() {
		ResetFileSystem()
	}()

	initTestTarget()
	f, err := Filesystem().Create(JoinWithUserDir(".tsuru", "token"))
	c.Assert(err, check.IsNil)
	f.WriteString("mytoken")
	token, err := ReadTokenV1()
	c.Assert(err, check.IsNil)
	c.Assert(token, check.Equals, "mytoken")
	tokenPath := JoinWithUserDir(".tsuru", "token.d", "test")
	c.Assert(rfs.HasAction("open "+tokenPath), check.Equals, true)
	tokenPath = JoinWithUserDir(".tsuru", "token")
	c.Assert(rfs.HasAction("open "+tokenPath), check.Equals, true)
}

func (s *S) TestReadTokenFileNotFound(c *check.C) {
	os.Unsetenv("TSURU_TOKEN")
	errFs := &fstest.FileNotFoundFs{}
	SetFileSystem(errFs)
	defer func() {
		ResetFileSystem()
	}()
	token, err := ReadTokenV1()
	c.Assert(err, check.IsNil)
	tokenPath := JoinWithUserDir(".tsuru", "token")
	c.Assert(err, check.IsNil)
	c.Assert(errFs.HasAction("open "+tokenPath), check.Equals, true)
	c.Assert(token, check.Equals, "")
}

func (s *S) TestReadTokenEnvironmentVariable(c *check.C) {
	os.Setenv("TSURU_TOKEN", "ABCDEFGH")
	defer os.Setenv("TSURU_TOKEN", "")
	token, err := ReadTokenV1()
	c.Assert(err, check.IsNil)
	c.Assert(token, check.Equals, "ABCDEFGH")
}

func initTestTarget() {
	f, _ := Filesystem().Create(JoinWithUserDir(".tsuru", "target"))
	f.Write([]byte("http://localhost"))
	f.Close()
	WriteOnTargetList("test", "http://localhost")
}
