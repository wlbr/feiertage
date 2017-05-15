package feiertage

import (
	"fmt"
	"sort"
)

// Region represents a Federal State of Germany or Austria (Bundesland). Some
// public holidays are common throughout the whole country, while others depend on the Bundesland.
// Short names of Austrian states are as suggested in ÖNORM A 1080.
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

func createFeiertagsList(y int, l string, ffun []func(int) Feiertag) []Feiertag {
	var feiern []func(int) Feiertag
	if l == "AT" {
		feiern = []func(int) Feiertag{Neujahr, HeiligeDreiKönige, Ostermontag,
			Staatsfeiertag, ChristiHimmelfahrt, PfingstMontag, Fronleichnam,
			MariäHimmelfahrt, Nationalfeiertag, Allerheiligen, MariäEmpfängnis,
			Christtag, Stefanitag}
	} else {
		feiern = []func(int) Feiertag{Neujahr, Karfreitag, Ostermontag,
			TagDerArbeit, ChristiHimmelfahrt, PfingstMontag,
			TagDerDeutschenEinheit, Weihnachten, ZweiterWeihnachtsfeiertag}
		// in 2017 the Reformationstag is a one time Feiertag in all states
		if y == 2017 {
			feiern = append(feiern, Reformationstag)
		}
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
	return Region{"Baden-Württemberg", "BW", createFeiertagsList(y, "DE", ffun)}
}

// Bayern returns a Region object holding all public holidays in the state Bayern
func Bayern(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Epiphanias, Fronleichnam, Allerheiligen}
	return Region{"Bayern", "BY", createFeiertagsList(y, "DE", ffun)}
}

// Berlin returns a Region object holding all public holidays in the state Berlin
func Berlin(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Berlin", "BE", createFeiertagsList(y, "DE", ffun)}
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
	return Region{"Brandenburg", "BB", createFeiertagsList(y, "DE", ffun)}
}

// Bremen returns a Region object holding all public holidays in the state Bremen
func Bremen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Bremen", "HB", createFeiertagsList(y, "DE", ffun)}
}

// Hamburg returns a Region object holding all public holidays in the state Hamburg
func Hamburg(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Hamburg", "HH", createFeiertagsList(y, "DE", ffun)}
}

// Hessen returns a Region object holding all public holidays in the state Hessen
func Hessen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Fronleichnam}
	return Region{"Hessen", "HE", createFeiertagsList(y, "DE", ffun)}
}

// MecklenburgVorpommern returns a Region object holding all public holidays in
// the state Mecklenburg-Vorpommern
func MecklenburgVorpommern(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Reformationstag}
	return Region{"Mecklenburg-Vorpommern", "MV", createFeiertagsList(y, "DE", ffun)}
}

// Niedersachsen returns a Region object holding all public holidays in the
// state Niedersachsen
func Niedersachsen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Niedersachsen", "NI", createFeiertagsList(y, "DE", ffun)}
}

// NordrheinWestfalen returns a Region object holding all public holidays in the
// state Nordrhein-Westfalen
func NordrheinWestfalen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Fronleichnam, Allerheiligen}
	return Region{"Nordrhein-Westfalen", "NW", createFeiertagsList(y, "DE", ffun)}
}

// RheinlandPfalz returns a Region object holding all public holidays in the
// state Rheinland-Pfalz
func RheinlandPfalz(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Fronleichnam, Allerheiligen}
	return Region{"Rheinland-Pfalz", "RP", createFeiertagsList(y, "DE", ffun)}
}

// Saarland returns a Region object holding all public holidays in the state Saarland
func Saarland(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Fronleichnam, MariäHimmelfahrt, Allerheiligen}
	return Region{"Saarland", "SL", createFeiertagsList(y, "DE", ffun)}
}

// Sachsen returns a Region object holding all public holidays in the state Sachsen
func Sachsen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Reformationstag, BußUndBettag}
	return Region{"Sachsen", "SN", createFeiertagsList(y, "DE", ffun)}
}

// SachsenAnhalt returns a Region object holding all public holidays in the state SachsenAnhalt
func SachsenAnhalt(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Epiphanias, Reformationstag}
	return Region{"Sachsen-Anhalt", "ST", createFeiertagsList(y, "DE", ffun)}
}

// SchleswigHolstein returns a Region object holding all public holidays in the state SchleswigHolstein
func SchleswigHolstein(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Schleswig-Holstein", "SH", createFeiertagsList(y, "DE", ffun)}
}

// Thüringen returns a Region object holding all public holidays in the state Thüringen
func Thüringen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{
		Reformationstag}
	return Region{"Thüringen", "TH", createFeiertagsList(y, "DE", ffun)}
}

// Deutschland returns a Region object holding all public holidays that are Common in Germany
func Deutschland(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}

	return Region{"Deutschland", "DE", createFeiertagsList(y, "DE", ffun)}
}

// Burgenland returns a Region object holding all public holidays in the state of Burgenland.
func Burgenland(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Martinstag}
	return Region{"Burgenland", "Bgld", createFeiertagsList(y, "AT", ffun)}
}

// Kärnten returns a Region object holding all public holidays in the state of Kärnten.
func Kärnten(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Josefitag, TagDerVolksabstimming}
	return Region{"Kärnten", "Ktn", createFeiertagsList(y, "AT", ffun)}
}

// Niederösterreich returns a Region object holding all public holidays in the state of Niederösterreich.
func Niederösterreich(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Leopolditag}
	return Region{"Niederösterreich", "NÖ", createFeiertagsList(y, "AT", ffun)}
}

// Oberösterreich returns a Region object holding all public holidays in the state of Oberösterreich.
func Oberösterreich(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Florianitag}
	return Region{"Oberösterreich", "OÖ", createFeiertagsList(y, "AT", ffun)}
}

// Salzburg returns a Region object holding all public holidays in the state of Salzburg.
func Salzburg(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Rupertitag}
	return Region{"Salzburg", "Sbg", createFeiertagsList(y, "AT", ffun)}
}

// Steiermark returns a Region object holding all public holidays in the state of Steiermark.
func Steiermark(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Josefitag}
	return Region{"Steiermark", "Stmk", createFeiertagsList(y, "AT", ffun)}
}

// Tirol returns a Region object holding all public holidays in the state of Tirol.
func Tirol(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Josefitag}
	return Region{"Tirol", "T", createFeiertagsList(y, "AT", ffun)}
}

// Vorarlberg returns a Region object holding all public holidays in the state of Vorarlberg.
func Vorarlberg(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Josefitag}
	return Region{"Vorarlberg", "Vbg", createFeiertagsList(y, "AT", ffun)}
}

// Wien returns a Region object holding all public holidays in the city and state of Vienna.
func Wien(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Leopolditag}
	return Region{"Wien", "W", createFeiertagsList(y, "AT", ffun)}
}

// Österreich returns a Region object holding all public holidays that are common in Austria.
func Österreich(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Österreich", "AT", createFeiertagsList(y, "AT", ffun)}
}

// All returns a Region object holding all public holidays/feast days known to this program.
// Not all of htem are public holidays (basically 'free').
func All(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Neujahr, Epiphanias, HeiligeDreiKönige, Valentinstag,
		Josefitag, Weiberfastnacht, Rosenmontag, Fastnacht, Aschermittwoch, Gründonnerstag,
		Karfreitag, BeginnSommerzeit, Ostermontag, Walpurgisnacht, TagDerArbeit, Staatsfeiertag,
		Florianitag, TagDerBefreiung, Muttertag, ChristiHimmelfahrt, Vatertag, PfingstMontag,
		Fronleichnam, MariäHimmelfahrt, Rupertitag, TagDerDeutschenEinheit,
		TagDerVolksabstimming, Nationalfeiertag, Reformationstag, Halloween, BeginnWinterzeit,
		Allerheiligen, Allerseelen, Martinstag, Karnevalsbeginn, Leopolditag, BußUndBettag,
		Thanksgiving, Blackfriday, Nikolaus, MariäUnbefleckteEmpfängnis, MariäEmpfängnis,
		Heiligabend, Weihnachten, Christtag, ZweiterWeihnachtsfeiertag, Stefanitag, Silvester}

	if len(inklSonntage) == 0 || inklSonntage[0] == true {
		ffun = append(ffun, Karnevalssonntag, Palmsonntag, Ostern, Pfingsten,
			Dreifaltigkeitssonntag, Erntedankfest, Volkstrauertag, Totensonntag,
			ErsterAdvent, ZweiterAdvent, DritterAdvent, VierterAdvent)
	}

	return Region{"Alle", "All", createFeiertagsList(y, "DE", ffun)}
}
