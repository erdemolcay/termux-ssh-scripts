package util

import "os"

func HomeDir() string {
	return os.Getenv("HOME")
}

func ShortcutsDir() string {
	return HomeDir() + "/.shortcuts"
}
