package feiertage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	checkAndFailRegionFeiertageZahl(t, All(2016), 69)
	checkAndFailRegionFeiertageZahl(t, All(2016, false), 57)
	//check Reformationstag 2017
	checkAndFailRegionFeiertageZahl(t, BadenWürttemberg(2017), 13)
	checkAndFailRegionFeiertageZahl(t, Bayern(2017), 13)
	checkAndFailRegionFeiertageZahl(t, Brandenburg(2017), 12)
	checkAndFailRegionFeiertageZahl(t, Deutschland(2017), 10)

	fmt.Println(Brandenburg(2017))

}

func TestGetRegionFromString(t *testing.T) {
	_, err := GetRegionFromString("", 2022)
	assert.Error(t, err)
	_, err = GetRegionFromString("Sachen-Anhalt", 2022)
	assert.Error(t, err)

	region, err := GetRegionFromString("Sachsen-Anhalt", 2022)
	assert.Equal(t, region, SachsenAnhalt(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Sachsen Anhalt", 2022)
	assert.Equal(t, region, SachsenAnhalt(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("SachsenAnhalt", 2022)
	assert.Equal(t, region, SachsenAnhalt(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Baden Württemberg", 2022)
	assert.Equal(t, region, BadenWürttemberg(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Bayern", 2022)
	assert.Equal(t, region, Bayern(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("bayern", 2022)
	assert.Equal(t, region, Bayern(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Berlin", 2022)
	assert.Equal(t, region, Berlin(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Brandenburg", 2022)
	assert.Equal(t, region, Brandenburg(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Bremen", 2022)
	assert.Equal(t, region, Bremen(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Hamburg", 2022)
	assert.Equal(t, region, Hamburg(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Niedersachsen", 2022)
	assert.Equal(t, region, Niedersachsen(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("NordrheinWestfalen", 2022)
	assert.Equal(t, region, NordrheinWestfalen(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("RheinlandPfalz", 2022)
	assert.Equal(t, region, RheinlandPfalz(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Saarland", 2022)
	assert.Equal(t, region, Saarland(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Sachsen", 2022)
	assert.Equal(t, region, Sachsen(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("SachsenAnhalt", 2022)
	assert.Equal(t, region, SachsenAnhalt(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("SchleswigHolstein", 2022)
	assert.Equal(t, region, SchleswigHolstein(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Thüringen", 2022)
	assert.Equal(t, region, Thüringen(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Thüringen", 2022)
	assert.Equal(t, region, Thüringen(2022))

	region, err = GetRegionFromString("Burgenland", 2022)
	assert.Equal(t, region, Burgenland(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Kärnten", 2022)
	assert.Equal(t, region, Kärnten(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Niederösterreich", 2022)
	assert.Equal(t, region, Niederösterreich(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Oberösterreich", 2022)
	assert.Equal(t, region, Oberösterreich(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Salzburg", 2022)
	assert.Equal(t, region, Salzburg(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Steiermark", 2022)
	assert.Equal(t, region, Steiermark(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Tirol", 2022)
	assert.Equal(t, region, Tirol(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Vorarlberg", 2022)
	assert.Equal(t, region, Vorarlberg(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Wien", 2022)
	assert.Equal(t, region, Wien(2022))
	assert.NoError(t, err)
	region, err = GetRegionFromString("Österreich", 2022)
	assert.Equal(t, region, Österreich(2022))
	assert.NoError(t, err)

}
