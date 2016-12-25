package feiertage

import (
	"fmt"
	"sort"
)

// Region represents a Federal State of Germany (Bundesland). Some of the public holidays in
// Germany are common throughout the whole country, while other depend of the Bundesland.
type Region struct {
	Name      string
	Shortname string
	Feiertage []Feiertag
}

// String returns a String containing the name and Shortname of the region and the list of
// public holidays (objects of type Feiertage).
func (r Region) String() string {
	s := fmt.Sprintf("%s (%s)", r.Name, r.Shortname)
	for _, f := range r.Feiertage {
		s = fmt.Sprintf("%s\n  %s", s, f)
	}
	return s
}

func createFeiertagsList(y int, ffun []func(int) Feiertag) []Feiertag {
	feiern := []func(int) Feiertag{Neujahr, Karfreitag, Ostermontag,
		TagDerArbeit, ChristiHimmelfahrt, PfingstMontag,
		TagDerDeutschenEinheit, Weihnachten, ZweiterWeihnachtsfeiertag}
	if y == 2017 { // in 2017 the Reformationstag is a one time feiertag in all states
		feiern = append(feiern, Reformationstag)
	}

	for _, f := range ffun {
		if (y != 2017) || (f(y).Text != Reformationstag(y).Text) {
			feiern = append(feiern, f)
		}
	}
	feiertermine := []Feiertag{}
	for _, f := range feiern {
		feiertermine = append(feiertermine, f(y))
	}
	sort.Sort(ByDate(feiertermine))
	return feiertermine
}

// BadenWürttemberg returns a Region object holding all public holidays in the state
// Baden-Württemberg
func BadenWürttemberg(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Epiphanias, Fronleichnam, Allerheiligen}
	return Region{"Baden-Württemberg", "BW", createFeiertagsList(y, ffun)}
}

// Bayern returns a Region object holding all public holidays in the state Bayern
func Bayern(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Epiphanias, Fronleichnam, Allerheiligen}
	return Region{"Bayern", "BY", createFeiertagsList(y, ffun)}
}

// Berlin returns a Region object holding all public holidays in the state Berlin
func Berlin(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Berlin", "BE", createFeiertagsList(y, ffun)}
}

// Brandenburg returns a Region object holding all public holidays in the state
// Brandenburg
func Brandenburg(y int, inklSonntage ...bool) Region {
	var ffun []func(int) Feiertag
	if len(inklSonntage) > 0 && inklSonntage[0] == false {
		ffun = []func(int) Feiertag{Reformationstag}
	} else {
		ffun = []func(int) Feiertag{Ostern, Pfingsten, Reformationstag}
	}
	return Region{"Brandenburg", "BB", createFeiertagsList(y, ffun)}
}

// Bremen returns a Region object holding all public holidays in the state Bremen
func Bremen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Bremen", "HB", createFeiertagsList(y, ffun)}
}

// Hamburg returns a Region object holding all public holidays in the state Hamburg
func Hamburg(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Hamburg", "HH", createFeiertagsList(y, ffun)}
}

// Hessen returns a Region object holding all public holidays in the state Hessen
func Hessen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Fronleichnam}
	return Region{"Hessen", "HE", createFeiertagsList(y, ffun)}
}

// MecklenburgVorpommern returns a Region object holding all public holidays in
// the state Mecklenburg-Vorpommern
func MecklenburgVorpommern(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Reformationstag}
	return Region{"Mecklenburg-Vorpommern", "MV", createFeiertagsList(y, ffun)}
}

// Niedersachsen returns a Region object holding all public holidays in the
// state Niedersachsen
func Niedersachsen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Niedersachsen", "NI", createFeiertagsList(y, ffun)}
}

// NordrheinWestfalen returns a Region object holding all public holidays in the
// state Nordrhein-Westfalen
func NordrheinWestfalen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Fronleichnam, Allerheiligen}
	return Region{"Nordrhein-Westfalen", "NW", createFeiertagsList(y, ffun)}
}

// RheinlandPfalz returns a Region object holding all public holidays in the
// state Rheinland-Pfalz
func RheinlandPfalz(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Fronleichnam, Allerheiligen}
	return Region{"Rheinland-Pfalz", "RP", createFeiertagsList(y, ffun)}
}

// Saarland returns a Region object holding all public holidays in the state Saarland
func Saarland(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Fronleichnam, MariäHimmelfahrt, Allerheiligen}
	return Region{"Saarland", "SL", createFeiertagsList(y, ffun)}
}

// Sachsen returns a Region object holding all public holidays in the state Sachsen
func Sachsen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Reformationstag, BußUndBettag}
	return Region{"Sachsen", "SN", createFeiertagsList(y, ffun)}
}

// SachsenAnhalt returns a Region object holding all public holidays in the state SachsenAnhalt
func SachsenAnhalt(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Epiphanias, Reformationstag}
	return Region{"Sachsen-Anhalt", "ST", createFeiertagsList(y, ffun)}
}

// SchleswigHolstein returns a Region object holding all public holidays in the state SchleswigHolstein
func SchleswigHolstein(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Schleswig-Holstein", "SH", createFeiertagsList(y, ffun)}
}

// Thüringen returns a Region object holding all public holidays in the state Thüringen
func Thüringen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{
		Reformationstag}
	return Region{"Thüringen", "TH", createFeiertagsList(y, ffun)}
}

// Deutschland returns a Region object holding all public holidays that are Common in Germany
func Deutschland(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}

	return Region{"Deutschland", "DE", createFeiertagsList(y, ffun)}
}

// All returns a Region object holding all public holidays/feast days known to this program.
// Not all of htem are public holidays (basically 'free').
func All(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Epiphanias, HeiligeDreiKönige, Valentinstag,
		Weiberfastnacht, Rosenmontag, Fastnacht, Aschermittwoch, Gründonnerstag,
		BeginnSommerzeit, Walpurgisnacht, TagDerBefreiung, Muttertag, Vatertag, Fronleichnam,
		MariäHimmelfahrt, Reformationstag, Halloween, BeginnWinterzeit,
		Allerheiligen, Allerseelen, Martinstag, Karnevalsbeginn, BußUndBettag, Thanksgiving,
		Blackfriday, Volkstrauertag, Nikolaus, MariäUnbefleckteEmpfängnis, Heiligabend, Silvester}

	if len(inklSonntage) == 0 || inklSonntage[0] == true {
		ffun = append(ffun, Karnevalssonntag, Palmsonntag, Ostern, Pfingsten,
			Dreifaltigkeitssonntag, Erntedankfest, Totensonntag, ErsterAdvent, ZweiterAdvent,
			DritterAdvent, VierterAdvent)
	}

	return Region{"Alle", "All", createFeiertagsList(y, ffun)}
}
