package core

import (
	"fmt"
	"math"
	"strconv"

	logs "github.com/mt1976/ebEstimates/logs"
)

// rtn Rounds a number to the nearest multiple of the RoundingFactor
func RoundToNearest(number float64, RoundingFactor float64) (float64, error) {
	// Round to the nearest multiple of the RoundingFactor
	if RoundingFactor == 0 {
		return number, nil
	}
	fmt.Printf("Number: %v\n", number)
	fmt.Printf("RoundingFactor: %v\n", RoundingFactor)
	rtnVal := math.Round(number/RoundingFactor) * RoundingFactor
	fmt.Printf("rtnVal: %v\n", rtnVal)
	return rtnVal, nil
}

func StringToFloat(in string) float64 {
	if in == "" {
		logs.Warning("Cannot convert empty string to float - assuming 0")
		return 0
	}
	out, strErr := strconv.ParseFloat(in, 64)
	if strErr != nil {
		logs.Warning("Cannot convert string to float - " + DQuote(strErr.Error()))
	}
	return out
}

func StringToInt(in string) int {
	if in == "" {
		logs.Warning("Cannot convert empty string to int - assuming 0")
		return 0
	}
	out, strErr := strconv.Atoi(in)
	if strErr != nil {
		logs.Warning("Cannot convert string to int - " + DQuote(strErr.Error()))
	}
	return out
}

func Numeric(in string) (float64, error) {
	out, strErr := strconv.ParseFloat(in, 64)
	if strErr != nil {
		logs.Warning("Cannot convert string to float" + DQuote(strErr.Error()))
		return 0, strErr
	}
	return out, nil
}

func FloatToString(in float64) string {
	out := strconv.FormatFloat(in, 'f', 2, 64)
	return out
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
