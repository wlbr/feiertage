package feiertage

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

//-------------------------

//-------------------------

func compareAndFail(t *testing.T, f Feiertag, d string) {
	if f.Format("02.01.2006") != d {
		fmt.Printf("%s but should be %s\n", f, d)
		t.Fail()
	}
}

func TestOstern(t *testing.T) {
	compareAndFail(t, Ostern(2015), "05.04.2015")
	compareAndFail(t, Ostern(2016), "27.03.2016")
}

func TestOsternAusnahmejahre(t *testing.T) {
	compareAndFail(t, Ostern(1954), "18.04.1954")
	compareAndFail(t, Ostern(1981), "19.04.1981")
}

func TestSommerWinterZeit(t *testing.T) {
	compareAndFail(t, BeginnSommerzeit(2015), "29.03.2015")
	compareAndFail(t, BeginnWinterzeit(2016), "30.10.2016")
}

func TestBußUndBetTag(t *testing.T) {
	compareAndFail(t, BußUndBettag(2015), "18.11.2015")
	compareAndFail(t, BußUndBettag(2016), "16.11.2016")
}

func TestVorwärtssucher(t *testing.T) {
	compareAndFail(t, Erntedankfest(2015), "04.10.2015")
	compareAndFail(t, Erntedankfest(2016), "02.10.2016")
	compareAndFail(t, Muttertag(2015), "10.05.2015")
	compareAndFail(t, Muttertag(2016), "08.05.2016")
}

func TestAdvent(t *testing.T) {
	// VierterAdvent=Rückwärtssucher
	compareAndFail(t, VierterAdvent(2016), "18.12.2016")
	compareAndFail(t, DritterAdvent(2016), "11.12.2016")
	compareAndFail(t, ZweiterAdvent(2016), "04.12.2016")
	compareAndFail(t, ErsterAdvent(2016), "27.11.2016")
	compareAndFail(t, VierterAdvent(2006), "24.12.2006")
	compareAndFail(t, VierterAdvent(2006), VierterAdvent(2006).Format("02.01.2006"))
}

//-------------------------

func TestFeiertage(t *testing.T) {

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
		feiern := []Feiertag{}
		for _, f := range fun {
			feiern = append(feiern, f(y))
		}
		sort.Sort(ByDate(feiern))
		for _, f := range feiern {
			fmt.Println(f)
		}
	}
}

func TestThanksgiving(t *testing.T) {
	//Vorwärtssucher Donnerstag
	compareAndFail(t, Thanksgiving(2010), "25.11.2010")
	compareAndFail(t, Thanksgiving(2014), "27.11.2014")
	compareAndFail(t, Thanksgiving(2015), "26.11.2015")
	compareAndFail(t, Thanksgiving(2016), "24.11.2016")
	compareAndFail(t, Thanksgiving(2017), "23.11.2017")
	compareAndFail(t, Thanksgiving(2018), "22.11.2018")
	compareAndFail(t, Thanksgiving(2019), "28.11.2019")
	compareAndFail(t, Thanksgiving(2025), "27.11.2025")
	compareAndFail(t, Thanksgiving(2028), "23.11.2028")
	compareAndFail(t, Thanksgiving(2029), "22.11.2029")
}

func TestDifferentTimeFormat(t *testing.T) {
	tf := defaultTimeFormat
	SetDefaultTimeFormat("2006.01.02")
	d := "2010.11.25 Thanksgiving (US)"
	f := Thanksgiving(2010)
	if fmt.Sprint(f) != d {
		fmt.Printf("%s but should be %s\n", f, d)
		t.Fail()
	}
	SetDefaultTimeFormat(tf) //back to old default
}

func TestDStandardTimeZone(t *testing.T) {
	tf := defaultTimeFormat
	SetDefaultTimeFormat(time.UnixDate)

	d := "Sun May 12 00:00:00 UTC 2019 Muttertag"
	f := Muttertag(2019)
	if fmt.Sprint(f) != d {
		fmt.Printf("%s but should be %s\n", f, d)
		t.Fail()

	}
	SetDefaultTimeFormat(tf) //back to old default

}

func TestDifferentTimeZone(t *testing.T) {
	tf := defaultTimeFormat
	SetDefaultTimeFormat(time.UnixDate)
	tz := defaultTimeZone

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	cet := time.FixedZone("CET", secondsEastOfUTC)
	// it would simpler to look for a timezone, but that works only on systems
	// having a timezone database present and that woud make the test flaky.
	// cet, err := time.LoadLocation("Europe/Berlin")
	SetDefaultTimeZone(cet)

	d := "Sun May 12 00:00:00 CET 2019 Muttertag"
	f := Muttertag(2019)
	if fmt.Sprint(f) != d {
		fmt.Printf("%s but should be %s\n", f, d)
		t.Fail()

	}
	SetDefaultTimeZone(tz)   //back to old default
	SetDefaultTimeFormat(tf) //back to old default

}
