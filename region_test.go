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
)

func TestRegion(t *testing.T) {
	// Test that region functions return non-empty results
	tests := []struct {
		name         string
		regionFn     func(int, ...bool) Region
		year         int
		includeSunday bool
	}{
		{"All 2016", All, 2016, true},
		{"Deutschland 2016", Deutschland, 2016, false},
		{"Brandenburg 2016", Brandenburg, 2016, true},
		{"Brandenburg 2016 no Sunday", Brandenburg, 2016, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r Region
			if tt.includeSunday {
				r = tt.regionFn(tt.year, true)
			} else {
				r = tt.regionFn(tt.year, false)
			}
			if len(r.Feiertage) == 0 {
				t.Errorf("%s returned empty Feiertage list", tt.name)
			}
			if r.Name == "" {
				t.Errorf("%s returned empty region name", tt.name)
			}
		})
	}
}

func TestRegionsInGermany(t *testing.T) {
	r := GetAllRegions(2020, false, "de")
	should := 16 + 1 //Bundesländer + 1 for "Alle"
	if should != len(r) {
		t.Errorf("Count Regions in Germany is %d but should be %d", len(r), should)
	}
}

func TestRegionsInAustria(t *testing.T) {
	r := GetAllRegions(2020, false, "at")
	should := 9 + 1 //Bundesländer + 1 for "Alle"
	if should != len(r) {
		t.Errorf("Count Regions in Austria is %d but should be %d", len(r), should)
	}
}

func TestRegionsAvailable(t *testing.T) {
	r := GetAllRegions(2020, false)
	should := 16 + 9 + 2 + 1 //German + austrian Bundesländer + 2 for the countries + 1 for "Alle"
	if should != len(r) {
		t.Errorf("Count Regions available is %d but should be %d", len(r), should)
	}
}

func TestFeiertageZahl(t *testing.T) {
	tests := []struct {
		name         string
		regionFn     func(int, ...bool) Region
		year         int
		includeSunday bool
		expectedCount int
	}{
		{"BadenWürttemberg 2016", BadenWürttemberg, 2016, false, 12},
		{"Bayern 2016", Bayern, 2016, false, 12},
		{"Berlin 2016", Berlin, 2016, false, 9},
		{"Berlin 2019", Berlin, 2019, false, 10},
		{"Berlin 2020", Berlin, 2020, false, 11},
		{"Brandenburg 2020", Brandenburg, 2020, true, 12},
		{"Brandenburg 2017", Brandenburg, 2017, true, 12},
		{"Brandenburg 2020 no Sunday", Brandenburg, 2020, false, 10},
		{"Bremen 2016", Bremen, 2016, false, 9},
		{"Bremen 2016 no Sunday", Bremen, 2016, false, 9},
		{"Hamburg 2016", Hamburg, 2016, false, 9},
		{"Hessen 2016", Hessen, 2016, false, 10},
		{"MecklenburgVorpommern 2016", MecklenburgVorpommern, 2016, false, 10},
		{"Niedersachsen 2016", Niedersachsen, 2016, false, 9},
		{"NordrheinWestfalen 2016", NordrheinWestfalen, 2016, false, 11},
		{"RheinlandPfalz 2016", RheinlandPfalz, 2016, false, 11},
		{"Saarland 2016", Saarland, 2016, false, 12},
		{"Sachsen 2016", Sachsen, 2016, false, 11},
		{"SachsenAnhalt 2016", SachsenAnhalt, 2016, false, 11},
		{"SchleswigHolstein 2016", SchleswigHolstein, 2016, false, 9},
		{"Thüringen 2016", Thüringen, 2016, false, 10},
		{"Deutschland 2016", Deutschland, 2016, false, 9},
		{"Deutschland 2017", Deutschland, 2017, false, 10},
		{"Burgenland 2016", Burgenland, 2016, false, 14},
		{"Kärnten 2016", Kärnten, 2016, false, 15},
		{"Niederösterreich 2016", Niederösterreich, 2016, false, 14},
		{"Oberösterreich 2016", Oberösterreich, 2016, false, 14},
		{"Salzburg 2016", Salzburg, 2016, false, 14},
		{"Steiermark 2016", Steiermark, 2016, false, 14},
		{"Tirol 2016", Tirol, 2016, false, 14},
		{"Vorarlberg 2016", Vorarlberg, 2016, false, 14},
		{"Wien 2016", Wien, 2016, false, 14},
		{"Österreich 2016", Österreich, 2016, false, 13},
		{"All 2016", All, 2016, true, 81},
		{"All 2016 no Sunday", All, 2016, false, 69},
		{"BadenWürttemberg 2017", BadenWürttemberg, 2017, false, 13},
		{"Bayern 2017", Bayern, 2017, false, 13},
		{"Deutschland 2017", Deutschland, 2017, false, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r Region
			if tt.includeSunday {
				r = tt.regionFn(tt.year, true)
			} else {
				r = tt.regionFn(tt.year, false)
			}
			if len(r.Feiertage) != tt.expectedCount {
				t.Errorf("Count Feiertage in %s is %d but should be %d", tt.name, len(r.Feiertage), tt.expectedCount)
			}
		})
	}
}

func TestBrandenburg2017(t *testing.T) {
	// Specific test for Brandenburg 2017 to verify Reformationstag is included
	r := Brandenburg(2017, true)
	if len(r.Feiertage) != 12 {
		t.Errorf("Brandenburg(2017, true) should have 12 Feiertage, got %d", len(r.Feiertage))
	}
	// t.Log(r.String()) // Use t.Log instead of fmt.Println for debugging
}
