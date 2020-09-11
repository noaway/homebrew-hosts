package core

import (
	"crypto/md5"
	"encoding/hex"
	"os/user"
	"path/filepath"
	"strconv"
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

type StrTo string

func (f StrTo) Exist() bool {
	return string(f) != string(0x1E)
}

func (f StrTo) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}

func (f StrTo) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 0)
	return int(v), err
}

func (f StrTo) MustInt() int {
	v, _ := f.Int()
	return v
}
