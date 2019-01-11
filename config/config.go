package config

import (
	"os"
	"os/user"
	"path"
)

// ConfigClientBaseDir is the path in $HOME where configs are stored/
const (
	ConfigRunnerBaseDir         = ".taask/runner/config/"
	ConfigRunnerDefaultFilename = "default-auth.yaml"
)

// DefaultRunnerConfigDir returns ~/.taask/runner/config unless XDG_CONFIG_HOME is set
func DefaultRunnerConfigDir() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}

	root := u.HomeDir
	xdgConfig, useXDG := os.LookupEnv("XDG_CONFIG_HOME")
	if useXDG {
		root = xdgConfig
	}

	return path.Join(root, ConfigRunnerBaseDir)
}
