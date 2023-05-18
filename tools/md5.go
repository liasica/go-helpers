// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-21
// Based on go-helpers by liasica, magicrolan@qq.com.

package tools

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5String hashes using md5 algorithm
func Md5String(text string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
