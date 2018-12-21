package common

import (
	"os"
	"os/user"
	"path"
)

func GetAppWorkingDir() (string, error) {
	userHomeDir, err := GetUserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(userHomeDir, "gomblr"), nil
}

func GetUserHomeDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return currentUser.HomeDir, nil
}

func GetAndMakeInAppWorkingDir(sub ...string) (string, error) {
	dir, err := GetAppWorkingDir()
	if err != nil {
		return "", err
	}
	subPath := path.Join(dir, path.Join(sub...))
	err = os.MkdirAll(subPath, os.ModePerm)
	return subPath, err
}
