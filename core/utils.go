package core

import (
	"crypto/md5"
	"encoding/hex"
	"os/user"
	"path/filepath"
)

func expand(path string) string {
	if len(path) == 0 || path[0] != '~' {
		return path
	}
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(usr.HomeDir, path[1:])
}

// md5HashBytes func
func md5HashBytes(data []byte) string {
	hasher := md5.New()
	_, _ = hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

func procErr(err error) {
	if err != nil {
		panic(err)
	}
}
