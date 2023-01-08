package core

import (
	"math"
	"strconv"

	logs "github.com/mt1976/ebEstimates/logs"
)

// rtn Rounds a number to the nearest multiple of the RoundingFactor
func RoundTo(number float64, RoundingFactor float64) (float64, error) {
	// Round to the nearest multiple of the RoundingFactor

	return math.Round(number/RoundingFactor) * RoundingFactor, nil
}

func StringToFloat(in string) float64 {
	out, strErr := strconv.ParseFloat(in, 64)
	if strErr != nil {
		logs.Warning("Cannot convert string to float" + DQuote(strErr.Error()))
	}
	return out
}

func FloatToString(in float64) string {
	out := strconv.FormatFloat(in, 'f', 2, 64)
	return out
}
