package feiertage

import (
	"fmt"
	"sort"
)

type Region struct {
	Name      string
	Shortname string
	Feiertage []Feiertag
}

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

	for _, f := range ffun {
		feiern = append(feiern, f)
	}
	feiertermine := []Feiertag{}
	for _, f := range feiern {
		feiertermine = append(feiertermine, f(y))
	}
	sort.Sort(FeiertageByDate(feiertermine))
	return feiertermine
}

func BadenWürttemberg(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Epiphanias, Fronleichnam, Allerheiligen}
	return Region{"Baden-Württemberg", "BW", createFeiertagsList(y, ffun)}
}

func Bayern(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Epiphanias, Fronleichnam, Allerheiligen}
	return Region{"Bayern", "BY", createFeiertagsList(y, ffun)}
}

func Berlin(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Berlin", "BE", createFeiertagsList(y, ffun)}
}

func Brandenburg(y int, inklSonntage ...bool) Region {
	var ffun []func(int) Feiertag
	if len(inklSonntage) > 0 && inklSonntage[0] == false {
		ffun = []func(int) Feiertag{Reformationstag}
	} else {
		ffun = []func(int) Feiertag{Ostern, Pfingsten, Reformationstag}
	}
	return Region{"Brandenburg", "BB", createFeiertagsList(y, ffun)}
}

func Bremen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Bremen", "HB", createFeiertagsList(y, ffun)}
}

func Hamburg(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Hamburg", "HH", createFeiertagsList(y, ffun)}
}

func Hessen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Fronleichnam}
	return Region{"Hessen", "HE", createFeiertagsList(y, ffun)}
}

func MecklenburgVorpommern(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Reformationstag}
	return Region{"Mecklenburg-Vorpommern", "MV", createFeiertagsList(y, ffun)}
}

func Niedersachsen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Niedersachsen", "NI", createFeiertagsList(y, ffun)}
}

func NordrheinWestfalen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Fronleichnam, Allerheiligen}
	return Region{"Nordrhein-Westfalen", "NW", createFeiertagsList(y, ffun)}
}

func RheinlandPfalz(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Fronleichnam, Allerheiligen}
	return Region{"Rheinland-Pfalz", "RP", createFeiertagsList(y, ffun)}
}

func Saarland(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Fronleichnam, MariäHimmelfahrt, Allerheiligen}
	return Region{"Saarland", "SL", createFeiertagsList(y, ffun)}
}

func Sachsen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Reformationstag, BußUndBettag}
	return Region{"Sachsen", "SN", createFeiertagsList(y, ffun)}
}

func SachsenAnhalt(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{Epiphanias, Reformationstag}
	return Region{"Sachsen-Anhalt", "ST", createFeiertagsList(y, ffun)}
}

func SchleswigHolstein(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}
	return Region{"Schleswig-Holstein", "SH", createFeiertagsList(y, ffun)}
}

func Thüringen(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{

		Reformationstag}
	return Region{"Thüringen", "TH", createFeiertagsList(y, ffun)}
}

func Deutschland(y int, inklSonntage ...bool) Region {
	ffun := []func(int) Feiertag{}

	return Region{"Deutschland", "DE", createFeiertagsList(y, ffun)}
}

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
