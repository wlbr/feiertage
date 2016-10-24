package main

import (
	"testing"
	"fmt"
	"github.com/wlbr/feiertage"
	"regexp"
)

func compareAndFail(t *testing.T, f feiertage.Feiertag, d string) {
	if f.Format("02.01.2006") != d {
		fmt.Printf("%s but should be %s\n", f, d)
		t.Fail()
	}
}


func TestGetRegion(t *testing.T) {
	bay := getRegion("Bayern", 2016, true)
	fron := bay.Feiertage[7]
	compareAndFail(t, fron, "26.05.2016")

	bra := getRegion("Brandenburg", 2016, true)
	ostern := bra.Feiertage[2]
	compareAndFail(t, ostern, "27.03.2016")

}


func TestFmtTaskjuggler(t *testing.T) {
	reg := getRegion("Brandenburg", 2016, true)
	jug := fmtTaskjuggler(reg)

	regex := regexp.MustCompile(`leaves holiday "Ostern" 2016-03-27`)
	if !regex.MatchString(jug) {
		t.Fail()
	}

}
