package md5

import "crypto/md5"

func DoMD5(m string) string {
	h := md5.New()
	h.Write([]byte(m))
	return string(h.Sum(nil))
}
