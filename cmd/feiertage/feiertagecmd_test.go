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

package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/wlbr/feiertage"
)

func compareAndFail(t *testing.T, f feiertage.Feiertag, d string) {
	if f.Format("02.01.2006") != d {
		fmt.Printf("%s but should be %s\n", f, d)
		t.Fail()
	}
}

func TestGetRegion(t *testing.T) {
	bay, e := getRegion("Bayern", 2016, true)
	if e != nil {
		fmt.Printf("Could not find region 'Bayern'")
		t.Fail()
	}
	fron := bay.Feiertage[7]
	compareAndFail(t, fron, "26.05.2016")

	bra, e := getRegion("Brandenburg", 2016, true)
	if e != nil {
		fmt.Printf("Could not find region 'Brandenburg'")
		t.Fail()
	}
	ostern := bra.Feiertage[2]
	compareAndFail(t, ostern, "27.03.2016")

}

func TestFmtTaskjuggler(t *testing.T) {
	reg, e := getRegion("Brandenburg", 2016, true)
	if e != nil {
		fmt.Printf("Could not find region 'Brandenburg'")
		t.Fail()
	}
	jug := fmtTaskjuggler(reg)

	regex := regexp.MustCompile(`leaves holiday "Ostern" 2016-03-27`)
	if !regex.MatchString(jug) {
		t.Fail()
	}

}
