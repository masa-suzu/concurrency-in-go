package chapter5

import "os"

type LowLevelError struct {
	error
}

func isGlobalExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelError{wrap(err, err.Error())}
	}
	return info.Mode().Perm()&0100 == 0100, nil
}
