// +build !linux,!darwin,!freebsd,!netbsd

package homedir

// Copyright 2013-2018 Docker, Inc.
// NOTE: this package has originally been copied from github.com/docker/docker.

import (
	"errors"
)

// GetRuntimeDir is unsupported on non-linux system.
func GetRuntimeDir() (string, error) {
	return "", errors.New("homedir.GetRuntimeDir() is not supported on this system")
}

// StickRuntimeDirContents is unsupported on non-linux system.
func StickRuntimeDirContents(files []string) ([]string, error) {
	return nil, errors.New("homedir.StickRuntimeDirContents() is not supported on this system")
}

// GetDataHome is unsupported on non-linux system.
func GetDataHome() (string, error) {
	return "", errors.New("homedir.GetDataHome() is not supported on this system")
}

// GetConfigHome is unsupported on non-linux system.
func GetConfigHome() (string, error) {
	return "", errors.New("homedir.GetConfigHome() is not supported on this system")
}

// GetCacheHome is unsupported on non-linux system.
func GetCacheHome() (string, error) {
	return "", errors.New("homedir.GetCacheHome() is not supported on this system")
}
