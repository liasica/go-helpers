// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-30
// Based on go-helpers by liasica, magicrolan@qq.com.

package tools

import "strconv"

func StrToFloat64(str string) (v float64) {
    v, _ = strconv.ParseFloat(str, 64)
    return
}

func StrToInt64(str string) (v int64) {
    v, _ = strconv.ParseInt(str, 10, 64)
    return
}
