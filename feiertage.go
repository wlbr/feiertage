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
50 Tage nach Ostern (Pfingstmontag)
60 Tage nach Ostern (Fronleichnam)

Keine Feiertage
46 Tage vor Ostern (Aschermittwoch)
Mittwoch vor dem 23. November (Buß- und Bettag)
24.12. (Heiligabend)
31.12. (Silvester)
*/

/* Thanksgiving
4. Donnerstag im November */

// Feiertag is an extented Time object. You may use it like any Time, but it offers an additional
// attribute carrying the name of the Feiertag.
type Feiertag struct {
	time.Time `json:"date"`
	Text      string `json:"name"`
	//TimeFormat string
}

var defaultTimeFormat = "02.01.2006"
var defaultTimeZone = time.UTC

func getTimeFormat() string {
	/*if f.TimeFormat != "" {
		return f.TimeFormat
	} */
	return defaultTimeFormat
}

// SetDefaultTimeFormat offers the possibility to change the default format for the ToString function.
// It defaults to "02.01.2006"
// The timezone is set at the creation time of the Feiertage object!
func SetDefaultTimeFormat(timeformat string) {
	defaultTimeFormat = timeformat
}

// SetDefaultTimeZone lets you set the timezone of the Feiertag functions. Default ist UTC.
func SetDefaultTimeZone(timezone *time.Location) {
	defaultTimeZone = timezone
}

func getTimeZone() *time.Location {
	return defaultTimeZone
}

// The String function of Feiertag will print its concrete Time (Date) plus the name of the Feiertag.
func (f Feiertag) String() string {
	return fmt.Sprintf("%s %s", f.Format(getTimeFormat()), f.Text)
}

// ----------------------------

// GetFeiertageForDateInRegion retrieves all bank holiday (gesetzliche Feiertage) for a given date in the specified region.
// Example:
//
//	regionfn := feiertage.Brandenburg
//	f := feiertage.GetFeiertageForDateInRegion(time.Date(2025, 6, 8, 1, 0, 0, 0, time.UTC), regionfn, true)
//	fmt.Printf("%s - %s: %v\n", f[0].Time.Format("2006-01-02"), regionfn(2025, true).Name, f)
//
// Output:
//
//	2025-06-08 - Brandenburg: [08.06.2025 Pfingsten]
func GetFeiertageForDateInRegion(t time.Time, regionFn func(y int, inklSonntage ...bool) Region, inklSonntage ...bool) []Feiertag {
	inclSundays := false
	if len(inklSonntage) > 0 && inklSonntage[0] {
		inclSundays = true
	}
	simplifiedT := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, getTimeZone())

	todaysFeiertage := []Feiertag{}
	for _, f := range regionFn(simplifiedT.Year(), inclSundays).Feiertage {
		if f.Time == simplifiedT {
			todaysFeiertage = append(todaysFeiertage, f)
		}
	}
	return todaysFeiertage
}

// GetFeiertageForDate returns all feiertage (defined dates in this library, even if not a 'gesetzlicher Feiertag') for a given date
// Example:
//
//	f := feiertage.GetFeiertageForDate(time.Date(2025, 6, 8, 1, 0, 0, 0, time.UTC))
//	fmt.Printf("%s - %s\n", f[0].Time.Format("2006-01-02"), f)
//
// Output:
//
//	2025-06-08 - [08.06.2025 Tag des Meeres 08.06.2025 Pfingsten]
func GetFeiertageForDate(t time.Time) []Feiertag {
	simplifiedT := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, getTimeZone())
	todaysFeiertage := []Feiertag{}
	for _, f := range All(simplifiedT.Year(), true).Feiertage {
		if f.Time == simplifiedT {
			todaysFeiertage = append(todaysFeiertage, f)
		}
	}
	return todaysFeiertage
}

// ----------------------------

// ByDate is the comparator object of Feiertag to be able to sort a list of Feiertage
type ByDate []Feiertag

// Len is the sort criteria for Feiertage
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
func Neujahr(year int) Feiertag {
	return Feiertag{time.Date(year, time.January, 1, 0, 0, 0, 0, getTimeZone()), "Neujahr"}
}

// Epiphanias is Epiphany, a fixed date.
func Epiphanias(year int) Feiertag {
	return Feiertag{time.Date(year, time.January, 6, 0, 0, 0, 0, getTimeZone()), "Epiphanias"}
}

// HeiligeDreiKönige is another name for Epiphany, a fixed date.
func HeiligeDreiKönige(year int) Feiertag {
	e := Epiphanias(year)
	e.Text = "Heilige drei Könige"
	return e
}

// Weltknuddeltag is World Hug Day or National Hugging Day, a fixed date.
func Weltknuddeltag(year int) Feiertag {
	return Feiertag{time.Date(year, time.January, 21, 0, 0, 0, 0, getTimeZone()), "Weltknuddeltag"}
}

// Valentinstag is Valentine's Day, a fixed date.
func Valentinstag(year int) Feiertag {
	return Feiertag{time.Date(year, time.February, 14, 0, 0, 0, 0, getTimeZone()), "Valentinstag"}
}

// InternationalerTagDesGedenkensAnDieOpferDesHolocaust is (International Holocaust Remembrance Day, a fixed date.
func InternationalerTagDesGedenkensAnDieOpferDesHolocaust(year int) Feiertag {
	return Feiertag{time.Date(year, time.January, 27, 0, 0, 0, 0, getTimeZone()), "Internationaler Tag des Gedenkens an die Opfer des Holocaust"}
}

// InternationalerFrauentag is International Women's Day, a fixed date.
func InternationalerFrauentag(year int) Feiertag {
	return Feiertag{time.Date(year, time.March, 8, 0, 0, 0, 0, getTimeZone()), "Internationaler Frauentag"}
}

// Josefitag is St Joseph's Day, a fixed date.
func Josefitag(year int) Feiertag {
	return Feiertag{time.Date(year, time.March, 19, 0, 0, 0, 0, getTimeZone()), "Josefitag"}
}

// Weiberfastnacht is a part of carnival, 52 days before Easter.
func Weiberfastnacht(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -52), "Weiberfastnacht"}
}

// Karnevalssonntag is the sunday of carnival, 49 days before Easter.
func Karnevalssonntag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -49), "Karnevalssonntag"}
}

// Rosenmontag is the monday of carnival, 48 days before Easter.
func Rosenmontag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -48), "Rosenmontag"}
}

// Fastnacht is shrovetide, the Tuesday of carnival, 47 days before Easter.
func Fastnacht(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -47), "Fastnacht"}
}

// Aschermittwoch is Ash Wednesday, 46 days before Easter.
func Aschermittwoch(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -46), "Aschermittwoch"}
}

// Palmsonntag is Palm Sunday , the last Sunday before Easter
func Palmsonntag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -7), "Palmsonntag"}
}

// Gründonnerstag is Holy Thursday or Maundy Thursday, the last Thursday before Eastern
func Gründonnerstag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -3), "Gründonnerstag"}
}

// Karfreitag is Good Friday, the last Friday before Easter
func Karfreitag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -2), "Karfreitag"}
}

// Ostern is Easter. Calculated by an extended Gauss algorithm.
func Ostern(year int) Feiertag {
	k := year / 100
	m := 15 + (3*k+3)/4 - (8*k+13)/25
	s := 2 - (3*k+3)/4
	a := year % 19
	d := (19*a + m) % 30
	r := (d + a/11) / 29
	og := 21 + d - r
	sz := 7 - (year+year/4+s)%7
	oe := 7 - (og-sz)%7
	os := og + oe

	day := os % 31
	month := os/31 + 3

	return Feiertag{time.Date(year, time.Month(month), day, 0, 0, 0, 0, getTimeZone()), "Ostern"}
}

// BeginnSommerzeit is the start of daylight saving time. Last Sunday of March.
func BeginnSommerzeit(year int) Feiertag {
	o := time.Date(year, time.March, 31, 0, 0, 0, 0, getTimeZone())
	d := (7 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Beginn Sommerzeit"}
}

// Ostermontag is Easter Monday, the Monday after Easter.
func Ostermontag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, 1), "Ostermontag"}
}

// Tag der Erde is Earth Day, a fixed date.
func TagDerErde(year int) Feiertag {
	return Feiertag{time.Date(year, time.April, 22, 0, 0, 0, 0, getTimeZone()), "Tag der Erde"}
}

// Walpurgisnacht is Walpurgis Night, a fixed date.
func Walpurgisnacht(year int) Feiertag {
	return Feiertag{time.Date(year, time.April, 30, 0, 0, 0, 0, getTimeZone()), "Walpurgisnacht"}
}

// TagDerArbeit is Labour Day, a fixed date.
func TagDerArbeit(year int) Feiertag {
	return Feiertag{time.Date(year, time.May, 1, 0, 0, 0, 0, getTimeZone()), "Tag der Arbeit"}
}

// Staatsfeiertag is May 1st in Austria, a fixed date.
func Staatsfeiertag(year int) Feiertag {
	e := TagDerArbeit(year)
	e.Text = "Staatsfeiertag"
	return e
}

// InternationalerTagDerPressefreiheit is World Press Freedom Day, a fixed date.
func InternationalerTagDerPressefreiheit(year int) Feiertag {
	return Feiertag{time.Date(year, time.May, 3, 0, 0, 0, 0, getTimeZone()), "Internationaler Tag der Pressefreiheit"}
}

// Florianitag is St Florian's Day, a fixed date.
func Florianitag(year int) Feiertag {
	return Feiertag{time.Date(year, time.May, 4, 0, 0, 0, 0, getTimeZone()), "Florianitag"}
}

// Star Wars Day is a fixed date.
func StarWarsDay(year int) Feiertag {
	return Feiertag{time.Date(year, time.May, 4, 0, 0, 0, 0, getTimeZone()), "Star Wars Day"}
}

// TagDerBefreiung is Victory in Europe Day, a fixed date.
func TagDerBefreiung(year int) Feiertag {
	return Feiertag{time.Date(year, time.May, 8, 0, 0, 0, 0, getTimeZone()), "Tag der Befreiung"}
}

// Muttertag is Mother's Day oder Mothering Sunday, the second Sunday in May.
func Muttertag(year int) Feiertag {
	o := time.Date(year, time.May, 1, 0, 0, 0, 0, getTimeZone())
	d := (7 - int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, d+7), "Muttertag"}
}

// ChristiHimmelfahrt is Ascension Day, 39 days after Easter, therefore always a Thursday.
func ChristiHimmelfahrt(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, 39), "Christi Himmelfahrt"}
}

// Vatertag is Father's Day, same day a Ascension Day, 39 days after Easter, therefore always a Thursday.
func Vatertag(year int) Feiertag {
	e := ChristiHimmelfahrt(year)
	e.Text = "Vatertag"
	return e
}

// Handtuchtag is Towel Day, May 25. It is celebrated as a tribute to the author Douglas Adams by his fans.
func Handtuchtag(year int) Feiertag {
	return Feiertag{time.Date(year, time.May, 25, 0, 0, 0, 0, getTimeZone()), "Handtuchtag"}
}

// TowelDay is, May 25. It is celebrated as a tribute to the author Douglas Adams by his fans.
func TowelDay(year int) Feiertag {
	e := Handtuchtag(year)
	e.Text = "Towel Day"
	return e
}

// Pfingsten is Pentecost, 49 days after Easter.
func Pfingsten(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, 49), "Pfingsten"}
}

// Pfingstmontag is Whit Monday, the monday after Pentecost.
func Pfingstmontag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, 50), "Pfingstmontag"}
}

// Dreifaltigkeitssonntag is Trinity Sunday, the Sunday after Pentecost
func Dreifaltigkeitssonntag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, 56), "Dreifaltigkeitssonntag"}
}

// Fronleichnam is Corpus Christi, 60 days after Eastern, therefore always a Thursday.
func Fronleichnam(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, 60), "Fronleichnam"}
}

// InternationalerKindertag is special to Germany and Austrian and
// isnot the same as Weltkindertag (World Children's Day), a fixed date.
func InternationalerKindertag(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 1, 0, 0, 0, 0, getTimeZone()), "Internationaler Kindertag"}
}

// Weltumwelttag is World Environment Day, a fixed date.
func Weltumwelttag(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 5, 0, 0, 0, 0, getTimeZone()), "Weltumwelttag"}
}

// TagDesMeeres is World Oceans Day, a fixed date.
func TagDesMeeres(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 8, 0, 0, 0, 0, getTimeZone()), "Tag des Meeres"}
}

// Weltspieltag is International Day of Play, a fixed date.
func Weltspieltag(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 11, 0, 0, 0, 0, getTimeZone()), "Weltspieltag"}
}

// Weltblutspendetag is World Blood Donor Day, a fixed date.
func Weltblutspendetag(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 14, 0, 0, 0, 0, getTimeZone()), "Weltblutspendetag"}
}

// Weltflüchtlingstag is World Refugee Day, a fixed date.
func Weltflüchtlingstag(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 20, 0, 0, 0, 0, getTimeZone()), "Weltflüchtlingstag"}
}

// Fête de la Musique is World Music Day, a fixed date.
func FêteDeLaMusique(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 21, 0, 0, 0, 0, getTimeZone()), "Fête de la Musique"}
}

// InternationalerTagGegenDrogenmissbrauch is International Day Against Drug Abuse and Illicit Trafficking, a fixed date.
func InternationalerTagGegenDrogenmissbrauch(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 26, 0, 0, 0, 0, getTimeZone()), "Internationaler Tag gegen Drogenmissbrauch"}
}

// SystemAdministratorAppreciationDay is the last Fridy in July
func SystemAdministratorAppreciationDay(year int) Feiertag {
	o := time.Date(year, time.July, 31, 0, 0, 0, 0, getTimeZone())
	d := (2 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "System Administrator Appreciation Day"}
}

// MariäHimmelfahrt is Assumption Day, a fixed date.
func MariäHimmelfahrt(year int) Feiertag {
	return Feiertag{time.Date(year, time.August, 15, 0, 0, 0, 0, getTimeZone()), "Mariä Himmelfahrt"}
}

// Hobbit Day is a fixed date.
func HobbitDay(year int) Feiertag {
	return Feiertag{time.Date(year, time.September, 22, 0, 0, 0, 0, getTimeZone()), "Hobbit Day"}
}

// Rupertitag is St Rupert's Day, a fixed date.
func Rupertitag(year int) Feiertag {
	return Feiertag{time.Date(year, time.September, 24, 0, 0, 0, 0, getTimeZone()), "Rupertitag"}
}

// TagDerDeutschenEinheit is German Unity Day, a fixed date.
func TagDerDeutschenEinheit(year int) Feiertag {
	return Feiertag{time.Date(year, time.October, 3, 0, 0, 0, 0, getTimeZone()), "Tag der deutschen Einheit"}
}

// TagDerVolksabstimmung is Referendum Day in Carinthia, a fixed date.
func TagDerVolksabstimmung(year int) Feiertag {
	return Feiertag{time.Date(year, time.October, 10, 0, 0, 0, 0, getTimeZone()), "Tag der Volksabstimmung"}
}

// Erntedankfest is Thanksgiving or Harvest Festival, the first Sunday of October.
// The german Erntedankfest is not the same than the US Thanksgiving.
func Erntedankfest(year int) Feiertag {
	o := time.Date(year, time.October, 1, 0, 0, 0, 0, getTimeZone())
	d := (7 - int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, d), "Erntedankfest"}
}

// Nationalfeiertag is the Austrian national day, a fixed date.
func Nationalfeiertag(year int) Feiertag {
	return Feiertag{time.Date(year, time.October, 26, 0, 0, 0, 0, getTimeZone()), "Nationalfeiertag"}
}

// Reformationstag is Reformation Day, a fixed date.
func Reformationstag(year int) Feiertag {
	return Feiertag{time.Date(year, time.October, 31, 0, 0, 0, 0, getTimeZone()), "Reformationstag"}
}

// Halloween is a fixed date.
func Halloween(year int) Feiertag {
	return Feiertag{time.Date(year, time.October, 31, 0, 0, 0, 0, getTimeZone()), "Halloween"}
}

// BeginnWinterzeit is the end of daylight saving time. Last Sunday of October.
func BeginnWinterzeit(year int) Feiertag {
	o := time.Date(year, time.October, 31, 0, 0, 0, 0, getTimeZone())
	d := (7 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Beginn Winterzeit"}
}

// Allerheiligen is All Saints' Day or Allhallows, a fixed date
func Allerheiligen(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 1, 0, 0, 0, 0, getTimeZone()), "Allerheiligen"}
}

// Allerseelen is All Souls' Day, the day after All Saints' Day,
func Allerseelen(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 2, 0, 0, 0, 0, getTimeZone()), "Allerseelen"}
}

// Weltmännertag is Men's World Day.
func Weltmännertag(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 3, 0, 0, 0, 0, getTimeZone()), "Weltmännertag"}
}

// Martinstag or Skt. Martin is Martinmas, a fixed date
func Martinstag(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 11, 0, 0, 0, 0, getTimeZone()), "Martinstag"}
}

// Internationaler Männertag is International Day Men's Day, a fixed date
func InternationalerMännertag(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 19, 0, 0, 0, 0, getTimeZone()), "Internationaler Männertag"}
}

// Karnevalsbeginn is the beginning of carnival, a fixed date.
func Karnevalsbeginn(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 11, 11, 11, 11, 11, getTimeZone()), "Karnevalsbeginn"}
}

// Leopolditag is St Leopold's Day, a fixed date.
func Leopolditag(x int) Feiertag {
	return Feiertag{time.Date(x, time.November, 15, 0, 0, 0, 0, getTimeZone()), "Leopolditag"}
}

// Weltkindertag is World Children's Day, a fixed date.
func Weltkindertag(x int) Feiertag {
	return Feiertag{time.Date(x, time.September, 20, 0, 0, 0, 0, getTimeZone()), "Weltkindertag"}
}

// BußUndBettag is Penance Day, 11 days before the first Sunday in Advent
func BußUndBettag(x int) Feiertag {
	o := time.Date(x, time.November, 22, 0, 0, 0, 0, getTimeZone())
	d := (4 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Buß- und Bettag"}
}

// Thanksgiving in the US, the fourth Thursday of November.
func Thanksgiving(x int) Feiertag {
	o := time.Date(x, time.November, 1, 0, 0, 0, 0, getTimeZone())
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
	return Feiertag{time.Date(x, time.December, 6, 0, 0, 0, 0, getTimeZone()), "Nikolaus"}
}

// MariäUnbefleckteEmpfängnis is Day of Immaculate Conception, a fixed date.
func MariäUnbefleckteEmpfängnis(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 8, 0, 0, 0, 0, getTimeZone()), "Mariä unbefleckte Empfängnis"}
}

// MariäEmpfängnis has a shorter name for MariäUnbefleckteEmpfängnis in Austria.
func MariäEmpfängnis(x int) Feiertag {
	e := MariäUnbefleckteEmpfängnis(x)
	e.Text = "Mariä Empfängnis"
	return e
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
	o := time.Date(x, time.December, 24, 0, 0, 0, 0, getTimeZone())
	d := (7 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Vierter Advent"}
}

// Heiligabend is Christmas Eve, the last day before Christmas.
func Heiligabend(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 24, 0, 0, 0, 0, getTimeZone()), "Heiligabend"}
}

// Weihnachten is Christmas, a fixed date
func Weihnachten(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 25, 0, 0, 0, 0, getTimeZone()), "Weihnachten"}
}

// Christtag is Christmas is  in Austria.
func Christtag(x int) Feiertag {
	e := Weihnachten(x)
	e.Text = "Christtag"
	return e
}

// ZweiterWeihnachtsfeiertag is day after Christmas, a fixed date
func ZweiterWeihnachtsfeiertag(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 26, 0, 0, 0, 0, getTimeZone()), "Zweiter Weihnachtsfeiertag"}
}

// Stefanitag is December 26th in Austria.
func Stefanitag(x int) Feiertag {
	e := ZweiterWeihnachtsfeiertag(x)
	e.Text = "Stefanitag"
	return e
}

// Silvester is NewYearsEve, a fixed date.
func Silvester(x int) Feiertag {
	return Feiertag{time.Date(x, time.December, 31, 0, 0, 0, 0, getTimeZone()), "Silvester"}
}
