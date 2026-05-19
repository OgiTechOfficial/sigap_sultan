package test

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"testing"
	"time"
)

func TestCoba(t *testing.T) {
	time, err := time.Parse("01/2006", "02/2020")
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Format("January 2006"))
}

func TestConvertDate(t *testing.T) {
	BulanIndonesiaFromEnglish := map[string]string{
		"january":   "Januari",
		"february":  "Februari",
		"march":     "Maret",
		"april":     "April",
		"may":       "Mei",
		"june":      "Juni",
		"july":      "Juli",
		"august":    "Agustus",
		"september": "September",
		"october":   "Oktober",
		"november":  "November",
		"december":  "Desember",
	}

	fmt.Println(BulanIndonesiaFromEnglish["january"])
}

func TestConvertCase(t *testing.T) {
	s := "KOTA MAKASSAR"

	fmt.Println(cases.Title(language.Indonesian).String(s))
}
