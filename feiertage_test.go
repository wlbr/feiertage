// Copyright (c) 2026 Michael Wolber
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package feiertage

import (
	"testing"
	"time"
)

func TestOstern(t *testing.T) {
	tests := []struct {
		name     string
		year     int
		expected string
	}{
		{"2015", 2015, "05.04.2015"},
		{"2016", 2016, "27.03.2016"},
		{"1954", 1954, "18.04.1954"},
		{"1981", 1981, "19.04.1981"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Ostern(tt.year)
			if result.Format("02.01.2006") != tt.expected {
				t.Errorf("Ostern(%d) = %v, want %s", tt.year, result.Format("02.01.2006"), tt.expected)
			}
		})
	}
}

func TestSommerWinterZeit(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(int) Feiertag
		year     int
		expected string
	}{
		{"BeginnSommerzeit 2015", BeginnSommerzeit, 2015, "29.03.2015"},
		{"BeginnWinterzeit 2016", BeginnWinterzeit, 2016, "30.10.2016"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn(tt.year)
			if result.Format("02.01.2006") != tt.expected {
				t.Errorf("%s(%d) = %v, want %s", tt.name, tt.year, result.Format("02.01.2006"), tt.expected)
			}
		})
	}
}

func TestBußUndBettag(t *testing.T) {
	tests := []struct {
		year     int
		expected string
	}{
		{2015, "18.11.2015"},
		{2016, "16.11.2016"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := BußUndBettag(tt.year)
			if result.Format("02.01.2006") != tt.expected {
				t.Errorf("BußUndBettag(%d) = %v, want %s", tt.year, result.Format("02.01.2006"), tt.expected)
			}
		})
	}
}

func TestVorwärtssucher(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(int) Feiertag
		year     int
		expected string
	}{
		{"Erntedankfest 2015", Erntedankfest, 2015, "04.10.2015"},
		{"Erntedankfest 2016", Erntedankfest, 2016, "02.10.2016"},
		{"Muttertag 2015", Muttertag, 2015, "10.05.2015"},
		{"Muttertag 2016", Muttertag, 2016, "08.05.2016"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn(tt.year)
			if result.Format("02.01.2006") != tt.expected {
				t.Errorf("%s(%d) = %v, want %s", tt.name, tt.year, result.Format("02.01.2006"), tt.expected)
			}
		})
	}
}

func TestAdvent(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(int) Feiertag
		year     int
		expected string
	}{
		{"VierterAdvent 2016", VierterAdvent, 2016, "18.12.2016"},
		{"DritterAdvent 2016", DritterAdvent, 2016, "11.12.2016"},
		{"ZweiterAdvent 2016", ZweiterAdvent, 2016, "04.12.2016"},
		{"ErsterAdvent 2016", ErsterAdvent, 2016, "27.11.2016"},
		{"VierterAdvent 2006", VierterAdvent, 2006, "24.12.2006"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn(tt.year)
			if result.Format("02.01.2006") != tt.expected {
				t.Errorf("%s(%d) = %v, want %s", tt.name, tt.year, result.Format("02.01.2006"), tt.expected)
			}
		})
	}
}

func TestThanksgiving(t *testing.T) {
	tests := []struct {
		year     int
		expected string
	}{
		{2010, "25.11.2010"},
		{2014, "27.11.2014"},
		{2015, "26.11.2015"},
		{2016, "24.11.2016"},
		{2017, "23.11.2017"},
		{2018, "22.11.2018"},
		{2019, "28.11.2019"},
		{2025, "27.11.2025"},
		{2028, "23.11.2028"},
		{2029, "22.11.2029"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := Thanksgiving(tt.year)
			if result.Format("02.01.2006") != tt.expected {
				t.Errorf("Thanksgiving(%d) = %v, want %s", tt.year, result.Format("02.01.2006"), tt.expected)
			}
		})
	}
}

func TestDifferentTimeFormat(t *testing.T) {
	originalFormat := defaultTimeFormat
	defer func() { SetDefaultTimeFormat(originalFormat) }()

	SetDefaultTimeFormat("2006.01.02")
	expected := "2010.11.25 Thanksgiving (US)"
	result := Thanksgiving(2010)
	if result.String() != expected {
		t.Errorf("Thanksgiving(2010) with custom format = %q, want %q", result.String(), expected)
	}
}

func TestDefaultTimeZone(t *testing.T) {
	originalFormat := defaultTimeFormat
	originalZone := defaultTimeZone
	defer func() {
		SetDefaultTimeFormat(originalFormat)
		SetDefaultTimeZone(originalZone)
	}()

	SetDefaultTimeFormat(time.UnixDate)
	expected := "Sun May 12 00:00:00 UTC 2019 Muttertag"
	result := Muttertag(2019)
	if result.String() != expected {
		t.Errorf("Muttertag(2019) with UnixDate format = %q, want %q", result.String(), expected)
	}
}

func TestDifferentTimeZone(t *testing.T) {
	originalFormat := defaultTimeFormat
	originalZone := defaultTimeZone
	defer func() {
		SetDefaultTimeFormat(originalFormat)
		SetDefaultTimeZone(originalZone)
	}()

	SetDefaultTimeFormat(time.UnixDate)
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	cet := time.FixedZone("CET", secondsEastOfUTC)
	SetDefaultTimeZone(cet)

	expected := "Sun May 12 00:00:00 CET 2019 Muttertag"
	result := Muttertag(2019)
	if result.String() != expected {
		t.Errorf("Muttertag(2019) with CET timezone = %q, want %q", result.String(), expected)
	}
}

func TestFeiertage(t *testing.T) {
	// Test that all holiday functions return valid dates
	fun := []func(int) Feiertag{Neujahr, Epiphanias, HeiligeDreiKönige, Valentinstag,
		InternationalerTagDesGedenkensAnDieOpferDesHolocaust, Josefitag, Weiberfastnacht,
		Karnevalssonntag, Rosenmontag, Fastnacht, Aschermittwoch, InternationalerFrauentag,
		Palmsonntag, Gründonnerstag, Karfreitag, Ostern, BeginnSommerzeit, Ostermontag,
		Walpurgisnacht, TagDerArbeit, TagDerBefreiung, Staatsfeiertag,
		InternationalerTagDerPressefreiheit, Florianitag, Muttertag, ChristiHimmelfahrt,
		Vatertag, Pfingsten, Pfingstmontag, Dreifaltigkeitssonntag, Fronleichnam, TagDesMeeres,
		MariäHimmelfahrt, Rupertitag, InternationalerKindertag, Weltflüchtlingstag,
		TagDerDeutschenEinheit, TagDerVolksabstimmung, Nationalfeiertag, Erntedankfest,
		Reformationstag, Halloween, BeginnWinterzeit, Allerheiligen, Allerseelen, Martinstag,
		Karnevalsbeginn, Leopolditag, Weltkindertag, BußUndBettag, Thanksgiving, Blackfriday,
		Volkstrauertag, Nikolaus, MariäUnbefleckteEmpfängnis, MariäEmpfängnis, Totensonntag,
		ErsterAdvent, ZweiterAdvent, DritterAdvent, VierterAdvent, Heiligabend, Weihnachten,
		Christtag, Stefanitag, ZweiterWeihnachtsfeiertag, Silvester}

	years := []int{2015, 2016}

	for _, y := range years {
		for _, f := range fun {
			result := f(y)
			// Just verify it doesn't panic and returns a valid date
			if result.Year() != y {
				t.Errorf("Holiday function returned year %d, want %d", result.Year(), y)
			}
		}
	}
}
