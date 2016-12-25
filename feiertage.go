package feiertage

import (
	"fmt"
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

// Feiertag is an extented Time object. You may use it like any Time, but it offers an additonal
// attribute carrying the name of the Feiertag.
type Feiertag struct {
	time.Time
	Text string
}

//The String function of Firetag will print its concrete Time (Date) plus the name of the Feiertag.
func (f Feiertag) String() string {
	return fmt.Sprintf("%s %s", f.Format("02.01.2006"), f.Text)
}

// ----------------------------

// ByDate is the comparator object of Feiertag to be able to sort a list of Feiertage
type ByDate []Feiertag

// Len is sort criteria for Feiertage
func (a ByDate) Len() int {
	return len(a)
}

// Swap exchanges two Feiertage within an array.
func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less is a comparator for Feiertage sorting.
func (a ByDate) Less(i, j int) bool {
	return a[i].UnixNano() < a[j].UnixNano()
}

// ----------------------------

// Neujahr is NewYear, a fixed date.
func Neujahr(x int) Feiertag {
	return Feiertag{time.Date(x, time.January, 1, 0, 0, 0, 0, time.UTC), "Neujahr"}
}

// Epiphanias is Epiphany, a fixed date.
func Epiphanias(x int) Feiertag {
	return Feiertag{time.Date(x, time.January, 6, 0, 0, 0, 0, time.UTC), "Epiphanias"}
}

// HeiligeDreiKönige is another Name for Epiphany, a fixed date.
func HeiligeDreiKönige(x int) Feiertag {
	e := Epiphanias(x)
	e.Text = "Heilige drei Könige"
	return e
}

// Valentinstag is Valentine's Daym a fixed date.
func Valentinstag(x int) Feiertag {
	return Feiertag{time.Date(x, time.February, 14, 0, 0, 0, 0, time.UTC), "Valentinstag"}
}

// Weiberfastnacht is a part of carnival, 52 days before Easter.
func Weiberfastnacht(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -52), "Weiberfastnacht"}
}

// Karnevalssonntag is the sunday of carnival, 49 days before Easter.
func Karnevalssonntag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -49), "Karnevalssonntag"}
}

// Rosenmontag is the monday of carnival, 48 days before Easter.
func Rosenmontag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -48), "Rosenmontag"}
}

// Fastnacht is shrovetide, the Tuesday of carnival, 47 days before Easter.
func Fastnacht(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -47), "Fastnacht"}
}

//Aschermittwoch is Ash Wednesday, 46 days before Easter.
func Aschermittwoch(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -46), "Aschermittwoch"}
}

// Palmsonntag is Palm Sunday , the last Sunday before Eastern
func Palmsonntag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -7), "Palmsonntag"}
}

// Gründonnerstag is Holy Thursday or Maundy Thursday, the last Thursday before Eastern
func Gründonnerstag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -3), "Gründonnerstag"}
}

// Karfreitag is Good Friday, the last Friday before Eastern
func Karfreitag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, -2), "Karfreitag"}
}

// Ostern is Easter. Calculated by an extended Gauss algorithm.
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

	return Feiertag{time.Date(x, time.Month(month), day, 0, 0, 0, 0, time.UTC), "Ostern"}
}

// BeginnSommerzeit is the start of daylight saving time. Last Sunday of March.
func BeginnSommerzeit(x int) Feiertag {
	o := time.Date(x, time.March, 30, 0, 0, 0, 0, time.UTC)
	d := (7 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Beginn Sommerzeit"}
}

// Ostermontag is Easter Monday, the Monday after Easter.
func Ostermontag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, 1), "Ostermontag"}
}

// Walpurgisnacht is Walpurgis Night, a fixed date.
func Walpurgisnacht(x int) Feiertag {
	return Feiertag{time.Date(x, time.April, 30, 0, 0, 0, 0, time.UTC), "Walpurgisnacht"}
}

// TagDerArbeit is Labour Day, a fixed date.
func TagDerArbeit(x int) Feiertag {
	return Feiertag{time.Date(x, time.May, 1, 0, 0, 0, 0, time.UTC), "Tag der Arbeit"}
}

// TagDerBefreiung is Victory in Europe Day, a fixed date.
func TagDerBefreiung(x int) Feiertag {
	return Feiertag{time.Date(x, time.May, 8, 0, 0, 0, 0, time.UTC), "Tag der Befreiung"}
}

// Muttertag is Mother's Day oder Mothering Sunday, the second Sunday in May.
func Muttertag(x int) Feiertag {
	o := time.Date(x, time.May, 1, 0, 0, 0, 0, time.UTC)
	d := (7 - int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, d+7), "Muttertag"}
}

// ChristiHimmelfahrt is Ascension Day, 39 days after Easter, therefore always a Thursday.
func ChristiHimmelfahrt(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, 39), "Christi Himmelfahrt"}
}

// Vatertag is Father's Day, same day a Ascension Day, 39 days after Easter, therefore always a Thursday.
func Vatertag(x int) Feiertag {
	e := ChristiHimmelfahrt(x)
	e.Text = "Vatertag"
	return e
}

//Pfingsten is Pentecost, 49 days after Easter.
func Pfingsten(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, 49), "Pfingsten"}
}

//PfingstMontag is Whit Monday, the monday after Pentecost.
func PfingstMontag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, 50), "Pfingstmontag"}
}

// Dreifaltigkeitssonntag is Trinity Sunday, the Sunday after Pentecost
func Dreifaltigkeitssonntag(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, 56), "Dreifaltigkeitssonntag"}
}

// Fronleichnam is Corpus Christi, 60 days after Eastern, therefore always a Thursday.
func Fronleichnam(x int) Feiertag {
	o := Ostern(x)
	return Feiertag{o.AddDate(0, 0, 60), "Fronleichnam"}
}

// MariäHimmelfahrt is Assumption Day, a fixed date.
func MariäHimmelfahrt(x int) Feiertag {
	return Feiertag{time.Date(x, time.August, 15, 0, 0, 0, 0, time.UTC), "Mariä Himmelfahrt"}
}

// TagDerDeutschenEinheit is German Unity Day, a fixed date.
func TagDerDeutschenEinheit(x int) Feiertag {
	return Feiertag{time.Date(x, time.October, 3, 0, 0, 0, 0, time.UTC), "Tag der deutschen Einheit"}
}

// Erntedankfest is Thanksgiving or Harvest Festival, the first Sunday of October.
// The german Erntedankfest is not the same than the US Thanksgiving.
func Erntedankfest(x int) Feiertag {
	o := time.Date(x, time.October, 1, 0, 0, 0, 0, time.UTC)
	d := (7 - int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, d), "Erntedankfest"}
}

// Reformationstag is Reformation Day, a fixed date.
func Reformationstag(x int) Feiertag {
	return Feiertag{time.Date(x, time.October, 31, 0, 0, 0, 0, time.UTC), "Reformationstag"}
}

// Halloween is a fixed date.
func Halloween(x int) Feiertag {
	return Feiertag{time.Date(x, time.October, 31, 0, 0, 0, 0, time.UTC), "Halloween"}
}

// BeginnWinterzeit is the end of daylight saving time. Last Sunday of October.
func BeginnWinterzeit(x int) Feiertag {
	o := time.Date(x, time.October, 31, 0, 0, 0, 0, time.UTC)
	d := (7 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Beginn Winterzeit"}
}

// Allerheiligen is All Saints' Day or Allhallows, a fixed date
func Allerheiligen(x int) Feiertag {
	return Feiertag{time.Date(x, time.November, 1, 0, 0, 0, 0, time.UTC), "Allerheiligen"}
}

// Allerseelen is All Souls' Day, the day after All Saints' Day,
func Allerseelen(x int) Feiertag {
	return Feiertag{time.Date(x, time.November, 2, 0, 0, 0, 0, time.UTC), "Allerseelen"}
}

// Martinstag or Skt. Martin is Martinmas, a fixed date
func Martinstag(x int) Feiertag {
	return Feiertag{time.Date(x, time.November, 11, 0, 0, 0, 0, time.UTC), "Martinstag"}
}

// Karnevalsbeginn is the beginning of carnival, a fixed date.
func Karnevalsbeginn(x int) Feiertag {
	return Feiertag{time.Date(x, time.November, 11, 11, 11, 11, 11, time.UTC), "Karnevalsbeginn"}
}

// BußUndBettag is Penance Day, 11 days before the first Sunday in Advent
func BußUndBettag(x int) Feiertag {
	o := time.Date(x, time.November, 22, 0, 0, 0, 0, time.UTC)
	d := (4 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Buß- und Bettag"}
}

// Thanksgiving in the US, the fourth Thursday of November.
func Thanksgiving(x int) Feiertag {
	o := time.Date(x, time.November, 1, 0, 0, 0, 0, time.UTC)
	d := ((11 - int(o.Weekday())) % 7)
	return Feiertag{o.AddDate(0, 0, 21+d), "Thanksgiving (US)"}
}

// Blackfriday is the Friday after Thanksgiving.
func Blackfriday(x int) Feiertag {
	return Feiertag{Thanksgiving(x).AddDate(0, 0, 1), "Blackfriday"}
}

// Volkstrauertag is Remembrance Sunday, the second sunday before the first Sunday in Advent
func Volkstrauertag(x int) Feiertag {
	o := ErsterAdvent(x)
	return Feiertag{o.AddDate(0, 0, -14), "Volkstrauertag"}
}

// Nikolaus is St Nicholas' Day, a fixed date
func Nikolaus(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 6, 0, 0, 0, 0, time.UTC), "Nikolaus"}
}

// MariäUnbefleckteEmpfängnis is Day of Immaculate Conception, a fixed date.
func MariäUnbefleckteEmpfängnis(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 8, 0, 0, 0, 0, time.UTC), "Mariä unbefleckte Empfängnis"}
}

// Totensonntag is Sunday in commemoration of the dead, the last Sunday before the fourth Sunday in Advent
func Totensonntag(x int) Feiertag {
	o := VierterAdvent(x)
	return Feiertag{o.AddDate(0, 0, -28), "Totensonntag"}
}

// ErsterAdvent is the first Sunday in Advent
func ErsterAdvent(x int) Feiertag {
	o := VierterAdvent(x)
	return Feiertag{o.AddDate(0, 0, -21), "Erster Advent"}
}

// ZweiterAdvent is the second Sunday in Advent
func ZweiterAdvent(x int) Feiertag {
	o := VierterAdvent(x)
	return Feiertag{o.AddDate(0, 0, -14), "Zweiter Advent"}
}

// DritterAdvent is the third Sunday in Advent
func DritterAdvent(x int) Feiertag {
	o := VierterAdvent(x)
	return Feiertag{o.AddDate(0, 0, -7), "Dritter Advent"}
}

// VierterAdvent is the fourth Sunday in Advent
func VierterAdvent(x int) Feiertag {
	o := time.Date(x, time.December, 24, 0, 0, 0, 0, time.UTC)
	d := (7 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Vierter Advent"}
}

// Heiligabend is Christmas Eve, the last day before Crhistmas.
func Heiligabend(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 24, 0, 0, 0, 0, time.UTC), "Heiligabend"}
}

// Weihnachten is Christmas, a fixed date
func Weihnachten(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 25, 0, 0, 0, 0, time.UTC), "Weihnachten"}
}

// ZweiterWeihnachtsfeiertag is day after Christmas, a fixed date
func ZweiterWeihnachtsfeiertag(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 26, 0, 0, 0, 0, time.UTC), "Zweiter Weihnachtsfeiertag"}
}

// Silvester is NewYearsEve, a fixed date.
func Silvester(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 31, 0, 0, 0, 0, time.UTC), "Silvester"}
}
