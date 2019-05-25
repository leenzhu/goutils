package utils

import (
	"os"
	"strings"
)

func GetProcessName() string {
	path, err := os.Readlink("/proc/self/exe")
	if err != nil {
		panic("Can't get process name")
	}
	idx := strings.LastIndex(path, "/")
	return path[idx+1:]
}


