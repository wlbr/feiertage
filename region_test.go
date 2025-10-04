package feiertage

import (
	"fmt"
	"testing"
)

func TestRegion(t *testing.T) {
	fmt.Println(All(2016))
	fmt.Println(Deutschland(2016))
	fmt.Println(Brandenburg(2016))
	fmt.Println(Brandenburg(2016, false))
}

func TestRegionsInGermany(t *testing.T) {
	r := GetAllRegions(2020, false, "de")
	should := 16 + 1 //Bundesländer + 1 for "Alle"
	if should != len(r) {
		fmt.Printf("Count Regions in Germany is %d but should be %d\n", len(r), should)
		t.Fail()
	}
}

func TestRegionsInAustria(t *testing.T) {
	r := GetAllRegions(2020, false, "at")
	should := 9 + 1 //Bundesländer + 1 for "Alle"
	if should != len(r) {
		fmt.Printf("Count Regions in Germany is %d but should be %d\n", len(r), should)
		t.Fail()
	}
}

func TestRegionsAvailable(t *testing.T) {
	r := GetAllRegions(2020, false)
	should := 16 + 9 + 2 + 1 //German + austrian Bundesländer + 2 for the countries + 1 for "Alle"
	if should != len(r) {
		fmt.Printf("Count Regions in Germany is %d but should be %d\n", len(r), should)
		t.Fail()
	}
}

func checkAndFailRegionFeiertageZahl(t *testing.T, r Region, c int) {
	if len(r.Feiertage) != c {
		fmt.Printf("Count Feiertage in %s is %d but should be %d\n", r.Name, len(r.Feiertage), c)
		t.Fail()
	}
}

func TestFeiertageZahl(t *testing.T) {
	checkAndFailRegionFeiertageZahl(t, BadenWürttemberg(2016), 12)
	checkAndFailRegionFeiertageZahl(t, Bayern(2016), 12)
	checkAndFailRegionFeiertageZahl(t, Berlin(2016), 9)
	checkAndFailRegionFeiertageZahl(t, Berlin(2019), 10)
	checkAndFailRegionFeiertageZahl(t, Berlin(2020), 11)
	checkAndFailRegionFeiertageZahl(t, Berlin(2019), 10)
	checkAndFailRegionFeiertageZahl(t, Brandenburg(2020), 12)
	checkAndFailRegionFeiertageZahl(t, Brandenburg(2017), len(Brandenburg(2016).Feiertage))
	checkAndFailRegionFeiertageZahl(t, Brandenburg(2020, false), 10)
	checkAndFailRegionFeiertageZahl(t, Bremen(2016), 9)
	checkAndFailRegionFeiertageZahl(t, Bremen(2016, false), 9)
	checkAndFailRegionFeiertageZahl(t, Hamburg(2016), 9)
	checkAndFailRegionFeiertageZahl(t, Hessen(2016), 10)
	checkAndFailRegionFeiertageZahl(t, MecklenburgVorpommern(2016), 10)
	checkAndFailRegionFeiertageZahl(t, Niedersachsen(2016), 9)
	checkAndFailRegionFeiertageZahl(t, NordrheinWestfalen(2016), 11)
	checkAndFailRegionFeiertageZahl(t, RheinlandPfalz(2016), 11)
	checkAndFailRegionFeiertageZahl(t, Saarland(2016), 12)
	checkAndFailRegionFeiertageZahl(t, Sachsen(2016), 11)
	checkAndFailRegionFeiertageZahl(t, SachsenAnhalt(2016), 11)
	checkAndFailRegionFeiertageZahl(t, SchleswigHolstein(2016), 9)
	checkAndFailRegionFeiertageZahl(t, Thüringen(2016), 10)
	checkAndFailRegionFeiertageZahl(t, Deutschland(2016), 9)
	checkAndFailRegionFeiertageZahl(t, Deutschland(2017), len(Deutschland(2016).Feiertage)+1)
	checkAndFailRegionFeiertageZahl(t, Burgenland(2016), 14)
	checkAndFailRegionFeiertageZahl(t, Kärnten(2016), 15)
	checkAndFailRegionFeiertageZahl(t, Niederösterreich(2016), 14)
	checkAndFailRegionFeiertageZahl(t, Oberösterreich(2016), 14)
	checkAndFailRegionFeiertageZahl(t, Salzburg(2016), 14)
	checkAndFailRegionFeiertageZahl(t, Steiermark(2016), 14)
	checkAndFailRegionFeiertageZahl(t, Tirol(2016), 14)
	checkAndFailRegionFeiertageZahl(t, Vorarlberg(2016), 14)
	checkAndFailRegionFeiertageZahl(t, Wien(2016), 14)
	checkAndFailRegionFeiertageZahl(t, Österreich(2016), 13)
	checkAndFailRegionFeiertageZahl(t, All(2016), 79)
	checkAndFailRegionFeiertageZahl(t, All(2016, false), 67)
	//check Reformationstag 2017
	checkAndFailRegionFeiertageZahl(t, BadenWürttemberg(2017), 13)
	checkAndFailRegionFeiertageZahl(t, Bayern(2017), 13)
	checkAndFailRegionFeiertageZahl(t, Brandenburg(2017), 12)
	checkAndFailRegionFeiertageZahl(t, Deutschland(2017), 10)

	fmt.Println(Brandenburg(2017))

}
