package main

import (
	"os/user"
	"errors"
)

func GetCurrentUserHomeDir() (homeDir string, err error){
	usr, err := user.Current()
	if err != nil {
			return "", errors.New("failed to get user homedir.") 
	}

	homeDir = usr.HomeDir 
	return homeDir, nil
}