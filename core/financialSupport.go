package core

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/leekchan/accounting"
)

func Financial_AbbreviationsToInt(str string) int {
	str = strings.ToUpper(str)
	number := ""
	number = strings.ReplaceAll(str, "$", "")
	number = strings.ReplaceAll(number, "€", "")
	number = strings.ReplaceAll(number, "£", "")
	fact := strings.ToUpper(number[len(number)-1:])
	number = strings.ReplaceAll(fact, "M", "")
	number = strings.ReplaceAll(number, "K", "")
	number = strings.ReplaceAll(number, "T", "")
	number = strings.ReplaceAll(number, "B", "")

	intNum, err := strconv.Atoi(number)
	if err != nil {
		log.Println(err.Error())
	}

	var retNum int
	switch fact {
	case "T":
		retNum = intNum * 1000
	case "K":
		retNum = intNum * 1000
	case "M":
		retNum = intNum * 1000000
	case "B":
		retNum = intNum * 1000000000
	default:
		retNum = intNum
	}

	return retNum
}

// Financial_GetSpotDate(inTime invalid type)
func Financial_GetSpotDate(inTime time.Time) time.Time {
	spot := inTime.AddDate(0, 0, 2)
	return wibbleDate(spot)
}

func wibbleDate(inDate time.Time) time.Time {
	if inDate.Weekday() == time.Saturday {
		inDate = inDate.AddDate(0, 0, 2)
	}
	if inDate.Weekday() == time.Sunday {
		inDate = inDate.AddDate(0, 0, 1)
	}
	return inDate
}

// CalculateSpotDate(inTime invalid type)
func Financial_GetTenorDate(inTime time.Time, inMonth string) time.Time {
	month, _ := strconv.Atoi(inMonth)
	spot := inTime.AddDate(0, month, 0)
	return wibbleDate(spot)
}

func Financial_GetFirstDayOfYear(inTime time.Time) time.Time {
	// Assuking 1st Jan is a holiday therefore first day is 2, then wibble the date.
	tempDate := time.Date(inTime.Year(), 1, 2, 0, 0, 0, inTime.Nanosecond(), inTime.Location())
	return wibbleDate(tempDate)
}

// Financial_FormatAmount returns a formated string version of a CCY amount
func Financial_FormatAmount(inAmount string, inCCY string) string {
	ac := accounting.Accounting{Symbol: inCCY, Precision: 2, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 ;\u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}

// Financial_FormatAmountFullDPS returns a formated string version of a CCY amount to 7dps
func Financial_FormatAmountFullDPS(inAmount string, inCCY string) string {
	prec, _ := strconv.Atoi("7")
	ac := accounting.Accounting{Symbol: inCCY, Precision: prec, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 \u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}

// Financial_FormatAmountToDPS returns a formated string version of a CCY amount to a given DPS
func Financial_FormatAmountToDPS(inAmount string, inCCY string, inPrec string) string {
	prec, _ := strconv.Atoi(inPrec)
	ac := accounting.Accounting{Symbol: inCCY, Precision: prec, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 \u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}
