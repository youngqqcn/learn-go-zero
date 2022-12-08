package tool

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MdByString(s string) string {
	m := md5.New()
	_, err := io.WriteString(m, s)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", m.Sum(nil))
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
