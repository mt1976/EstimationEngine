package core

import (
	"math"
	"strconv"

	logs "github.com/mt1976/ebEstimates/logs"
)

// rtn Rounds a number to the nearest multiple of the RoundingFactor
func RoundUpTo(number float64, roundto float64) (float64, error) {
	// Round to the nearest multiple of the RoundingFactor
	if roundto == 0 {
		return number, nil
	}
	//fmt.Printf("Number: %v\n", number)
	//fmt.Printf("RoundingFactor: %v\n", RoundingFactor)
	//rtnVal := math.Round(number/RoundingFactor) * RoundingFactor
	rtnVal := math.Floor(number) + math.Ceil((number-math.Floor(number))/roundto)*roundto
	//msg := fmt.Sprintf("Rounding %v to %v (%v)", number, rtnVal, RoundingFactor)
	//logs.Information("RoundToNearest", msg)

	//fmt.Printf("rtnVal: %v\n", rtnVal)
	return rtnVal, nil
}

// func RoundMe(num float64, roundto float64) float64 {
// 	fmt.Printf(">>> Round Up: %v to nearest %v\n", num, roundto)

// 	//num = num + .0000001
// 	//roundto = 1 / roundto
// 	//ceilIn := num * roundto
// 	//returnval := clng(ceilIn) / roundto

// 	returnval := flr(num) + clng((num-flr(num))/roundto)*roundto
// 	fmt.Printf(">>> Result : %v\n", returnval)

// 	return returnval
// }

// func clng(num float64) float64 {
// 	rtn := math.Ceil(num)
// 	fmt.Printf("Result clng: %v from %v\n", rtn, num)
// 	return rtn
// }

// func flr(num float64) float64 {
// 	rtn := math.Floor(num)
// 	fmt.Printf("Result flr: %v from %v\n", rtn, num)
// 	return rtn
// }

func StringToFloat(in string) float64 {
	if in == "" {
		//logs.Warning("Cannot convert empty string to float - assuming 0")
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
		//logs.Warning("Cannot convert empty string to int - assuming 0")
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
