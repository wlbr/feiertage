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
