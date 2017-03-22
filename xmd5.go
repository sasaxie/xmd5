/**
 * MD5 tools
 */

package xmd5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

/*
 * Get string MD5 value
 */
func GetStringMD5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	src := hash.Sum(nil)

	return hex.EncodeToString(src)
}

/*
 * Get file MD5 value
 */
func GetFileMD5(filename string) string {
	resultString := ""

	file, err := os.Open(filename)

	if err == nil {
		hash := md5.New()
		io.Copy(hash, file)
		resultString = hex.EncodeToString(hash.Sum(nil))
	}

	return resultString
}
