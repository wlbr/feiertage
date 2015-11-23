package feiertage

import (
	"fmt"
	//"sort"
	"time"
)

/* https://de.wikipedia.org/wiki/Gau%C3%9Fsche_Osterformel#Eine_erg.C3.A4nzte_Osterformel
Schritt	Bedeutung	Formel
1.	die Säkularzahl	K(X) = X div 100
2.	die säkulare Mondschaltung	M(K) = 15 + (3K + 3) div 4 − (8K + 13) div 25
3.	die säkulare Sonnenschaltung	S(K) = 2 − (3K + 3) div 4
4.	den Mondparameter	A(X) = X mod 19
5.	den Keim für den ersten Vollmond im Frühling	D(A,M) = (19A + M) mod 30
6.	die kalendarische Korrekturgröße	R(D,A) = (D + A div 11) div 29[13]
7.	die Ostergrenze	OG(D,R) = 21 + D − R
8.	den ersten Sonntag im März	SZ(X,S) = 7 − (X + X div 4 + S) mod 7
9.	die Entfernung des Ostersonntags von der Ostergrenze
(Osterentfernung in Tagen)	OE(OG,SZ) = 7 − (OG − SZ) mod 7
10.	das Datum des Ostersonntags als Märzdatum
(32. März = 1. April usw.)	OS = OG + OE
*/

/* https://de.wikipedia.org/wiki/Kalenderrechnung
w

in allen deutschen Bundesländern folgende unbewegliche Feiertage:
01.01. (Neujahr)
01.05. (Maifeiertag)
03.10. (Tag der Deutschen Einheit)
25.12. (1. Weihnachtsfeiertag)
26.12. (2. Weihnachtsfeiertag)

in einigen Bundesländern unbewegliche Feiertage:
06.01. (Heilige Drei Könige)
15.08. (Mariä Himmelfahrt)
31.10. (Reformationstag)
01.11. (Allerheiligen)

bewegliche Feiertage in allen Bundesländern:
2 Tage vor Ostern (Karfreitag)
39 Tage nach Ostern (Christi Himmelfahrt)
49 Tage nach Ostern (Pfingstsonntag)
60 Tage nach Ostern (Fronleichnam)

Keine Feiertage
50 Tage nach Ostern (Pfingstmontag)
46 Tage vor Ostern (Aschermittwoch)
Mittwoch vor dem 23. November (Buß- und Bettag)
24.12. (Heiligabend)
31.12. (Silvester)
*/

/* Thanksgiving
4. Donnerstag im November */

type Feiertag struct {
	time.Time
	Text string
}

func (f Feiertag) String() string {
	return fmt.Sprintf("%s %s", f.Format("02.01.2006"), f.Text)
}

// ----------------------------

type FeiertageByDate []Feiertag

func (a FeiertageByDate) Len() int {
	return len(a)
}

func (a FeiertageByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a FeiertageByDate) Less(i, j int) bool {
	return a[i].UnixNano() < a[j].UnixNano()
}

// ----------------------------

func Neujahr(x int) Feiertag {
	return Feiertag{time.Date(x, time.January, 1, 0, 0, 0, 0, time.UTC), "Neujahr"}
}

func Epiphanias(x int) Feiertag {
	return Feiertag{time.Date(x, time.January, 6, 0, 0, 0, 0, time.UTC), "Epiphanias"}
}

func HeiligeDreiKönige(x int) Feiertag {
	e := Epiphanias(x)
	e.Text = "Heilige drei Könige"
	return e
}

func Valentinstag(x int) Feiertag {
	return Feiertag{time.Date(x, time.February, 14, 0, 0, 0, 0, time.UTC), "Valentinstag"}
}

func Weiberfastnacht(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -52), "Weiberfastnacht"}
}

func Karnevalssonntag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -49), "Karnevalssonntag"}
}

func Rosenmontag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -48), "Rosenmontag"}
}

func Fastnacht(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -47), "Fastnacht"}
}

func Aschermittwoch(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -46), "Aschermittwoch"}
}

func Palmsonntag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -7), "Ostermontag"}
}

func Gründonnerstag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -3), "Gründonnerstag"}
}

func Karfreitag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -2), "Karfreitag"}
}

func Ostern(x int) Feiertag {
	k := x / 100
	m := 15 + (3*k+3)/4 - (8*k+13)/25
	s := 2 - (3*k+3)/4
	a := x % 19
	d := (19*a + m) % 30
	r := (d + a/11) / 29
	og := 21 + d - r
	sz := 7 - (x+x/4+s)%7
	oe := 7 - (og-sz)%7
	os := og + oe

	day := os % 31
	month := os/31 + 3
	// fmt.Println(x)
	return Feiertag{time.Date(x, time.Month(month), day, 0, 0, 0, 0, time.UTC), "Ostern"}
}

func BeginnSommerzeit(x int) Feiertag {
	o := time.Date(x, time.March, 30, 0, 0, 0, 0, time.UTC)
	d := (7 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Beginn Sommerzeit"}
}

func Ostermontag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, 1), "Ostermontag"}
}

func Walpurgisnacht(x int) Feiertag {
	return Feiertag{time.Date(x, time.April, 30, 0, 0, 0, 0, time.UTC), "Walpurgisnacht"}
}

func TagDerArbeit(x int) Feiertag {
	return Feiertag{time.Date(x, time.May, 1, 0, 0, 0, 0, time.UTC), "Tag der Arbeit"}
}

func TagDerBefreiung(x int) Feiertag {
	return Feiertag{time.Date(x, time.May, 8, 0, 0, 0, 0, time.UTC), "Tag der Befreiung"}
}

func Muttertag(x int) Feiertag {
	o := time.Date(x, time.May, 1, 0, 0, 0, 0, time.UTC)
	d := (7 - int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, d+7), "Muttertag"}
}

func ChristiHimmelfahrt(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, 39), "Christi Himmelfahrt"}
}

func Vatertag(x int) Feiertag {
	e := ChristiHimmelfahrt(x)
	e.Text = "Vatertag"
	return e
}

func Pfingsten(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, 49), "Pfingsten"}
}

func PfingstMontag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, 50), "Pfingstmontag"}
}

func Dreifaltigkeitssonntag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, 56), "Dreifaltigkeitssonntag"}
}

func Fronleichnam(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, 60), "Fronleichnam"}
}

func MariäHimmelfahrt(x int) Feiertag {
	return Feiertag{time.Date(x, time.August, 15, 0, 0, 0, 0, time.UTC), "Mariä Himmelfahrt"}
}

func TagDerDeutschenEinheit(x int) Feiertag {
	return Feiertag{time.Date(x, time.October, 3, 0, 0, 0, 0, time.UTC), "Tag der deutschen Einheit"}
}

func Erntedankfest(x int) Feiertag {
	o := time.Date(x, time.October, 1, 0, 0, 0, 0, time.UTC)
	d := (7 - int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, d), "Erntedankfest"}
}

func Reformationstag(x int) Feiertag {
	return Feiertag{time.Date(x, time.October, 31, 0, 0, 0, 0, time.UTC), "Reformationstag"}
}

func Halloween(x int) Feiertag {
	return Feiertag{time.Date(x, time.October, 31, 0, 0, 0, 0, time.UTC), "Halloween"}
}

func BeginnWinterzeit(x int) Feiertag {
	o := time.Date(x, time.October, 31, 0, 0, 0, 0, time.UTC)
	d := (7 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Beginn Winterzeit"}
}

func Allerheiligen(x int) Feiertag {
	return Feiertag{time.Date(x, time.November, 1, 0, 0, 0, 0, time.UTC), "Allerheiligen"}
}

func Allerseelen(x int) Feiertag {
	return Feiertag{time.Date(x, time.November, 2, 0, 0, 0, 0, time.UTC), "Allerseelen"}
}

func Martinstag(x int) Feiertag {
	return Feiertag{time.Date(x, time.November, 11, 0, 0, 0, 0, time.UTC), "Martinstag"}
}

func Karnevalsbeginn(x int) Feiertag {
	return Feiertag{time.Date(x, time.November, 11, 11, 11, 11, 11, time.UTC), "Karnevalsbeginn"}
}

func BußUndBettag(x int) Feiertag {
	o := time.Date(x, time.November, 22, 0, 0, 0, 0, time.UTC)
	d := (4 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Buß- und Bettag"}
}

func Thanksgiving(x int) Feiertag {
	o := time.Date(x, time.November, 1, 0, 0, 0, 0, time.UTC)
	d := ((11 - int(o.Weekday())) % 7)
	return Feiertag{o.AddDate(0, 0, 21+d), "Thanksgiving (US)"}
}

func Blackfriday(x int) Feiertag {
	return Feiertag{Thanksgiving(x).AddDate(0, 0, 1), "Blackfriday"}
}

func Volkstrauertag(x int) Feiertag {
	o := ErsterAdvent(x)
	return Feiertag{o.AddDate(0, 0, -14), "Volkstrauertag"}
}

func Nikolaus(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 6, 0, 0, 0, 0, time.UTC), "Nikolaus"}
}

func MariäUnbefleckteEmpfängnis(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 8, 0, 0, 0, 0, time.UTC), "Mariä unbefleckte Empfängnis"}
}

func Totensonntag(x int) Feiertag {
	o := VierterAdvent(x)
	return Feiertag{o.AddDate(0, 0, -28), "Totensonntag"}
}

func ErsterAdvent(x int) Feiertag {
	o := VierterAdvent(x)
	return Feiertag{o.AddDate(0, 0, -21), "Erster Advent"}
}

func ZweiterAdvent(x int) Feiertag {
	o := VierterAdvent(x)
	return Feiertag{o.AddDate(0, 0, -14), "Zweiter Advent"}
}

func DritterAdvent(x int) Feiertag {
	o := VierterAdvent(x)
	return Feiertag{o.AddDate(0, 0, -7), "Dritter Advent"}
}

func VierterAdvent(x int) Feiertag {
	o := time.Date(x, time.December, 24, 0, 0, 0, 0, time.UTC)
	d := (7 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Vierter Advent"}
}

func Heiligabend(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 24, 0, 0, 0, 0, time.UTC), "Heiligabend"}
}

func Weihnachten(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 25, 0, 0, 0, 0, time.UTC), "Weihnachten"}
}

func ZweiterWeihnachtsfeiertag(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 26, 0, 0, 0, 0, time.UTC), "Zweiter Weihnachtsfeiertag"}
}

func Silvester(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 31, 0, 0, 0, 0, time.UTC), "Silvester"}
}
