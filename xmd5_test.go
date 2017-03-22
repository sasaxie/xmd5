package xmd5

import (
	"log"
	"testing"
)

func TestGetStringMD5(t *testing.T) {
	log.Println(GetStringMD5("test md5"))
}

func TestGetFileMD5(t *testing.T) {
	log.Println(GetFileMD5("xmd5_test.go"))
}
