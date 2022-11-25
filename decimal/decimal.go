// Copyright (C) liasica. 2022-present.
//
// Created at 2022-11-25
// Based on aurservd by liasica, magicrolan@qq.com.

package decimal

import (
    "github.com/shopspring/decimal"
    "math"
)

// Sum returns f1 + f2
func Sum(f1, f2 float64) float64 {
    f, _ := decimal.NewFromFloat(f1).Add(decimal.NewFromFloat(f2)).Float64()
    return math.Round(f*100.00) / 100.0
}

// Sub returns f1 - f2
func Sub(f1, f2 float64) float64 {
    f, _ := decimal.NewFromFloat(f1).Sub(decimal.NewFromFloat(f2)).Float64()
    return math.Round(f*100.00) / 100.0
}

// Mul returns f1 × f2
func Mul(f1, f2 float64) float64 {
    f, _ := decimal.NewFromFloat(f1).Mul(decimal.NewFromFloat(f2)).Float64()
    return math.Round(f*100.0) / 100.0
}

// Div returns f1 ÷ f2
func Div(f1, f2 float64) float64 {
    f, _ := decimal.NewFromFloat(f1).DivRound(decimal.NewFromFloat(f2), 2).Float64()
    return f
}
