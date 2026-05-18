package feiertage

import (
	"encoding/json"
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
	time.Time
	Text    string `json:"name"`
	Regions []Region `json:"-"`
}

// MarshalJSON implements the json.Marshaler interface for Feiertag.
// func (f Feiertag) MarshalJSON() ([]byte, error) {
// 	fmt.Println("Called")
// 	type Alias Feiertag // Create an alias to avoid infinite recursion
// 	return json.Marshal(&struct {
// 		Alias
// 		Date string `json:"date"`
// 	}{
// 		Alias: (Alias)(f),
// 		Date:  f.Format("2006-01-02"), // Format the date as YYYY-MM-DD
// 	})
// }

func (f Feiertag) MarshalJSON() ([]byte, error) {
	ts := &struct {
		Date string `json:"date"`
		Name string `json:"name"`
	}{
		Date: f.Format(time.DateOnly),
		Name: f.Text,
	}
	return json.Marshal(ts)
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
	return Feiertag{time.Date(year, time.January, 1, 0, 0, 0, 0, getTimeZone()), "Neujahr", []Region{}}
}

// Epiphanias is Epiphany, a fixed date.
func Epiphanias(year int) Feiertag {
	return Feiertag{time.Date(year, time.January, 6, 0, 0, 0, 0, getTimeZone()), "Epiphanias", []Region{}}
}

// HeiligeDreiKönige is another name for Epiphany, a fixed date.
func HeiligeDreiKönige(year int) Feiertag {
	e := Epiphanias(year)
	e.Text = "Heilige drei Könige"
	return e
}

// InternationalerFrauentag is International Women's Day, a fixed date.
func InternationalerFrauentag(year int) Feiertag {
	return Feiertag{time.Date(year, time.March, 8, 0, 0, 0, 0, getTimeZone()), "Internationaler Frauentag", []Region{}}
}

// Josefitag is St Joseph's Day, a fixed date.
func Josefitag(year int) Feiertag {
	return Feiertag{time.Date(year, time.March, 19, 0, 0, 0, 0, getTimeZone()), "Josefitag", []Region{}}
}

// Karfreitag is Good Friday, the last Friday before Easter
func Karfreitag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -2), "Karfreitag", []Region{}}
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

	return Feiertag{time.Date(year, time.Month(month), day, 0, 0, 0, 0, getTimeZone()), "Ostern", []Region{}}
}

// Ostermontag is Easter Monday, the Monday after Easter.
func Ostermontag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, 1), "Ostermontag", []Region{}}
}

// TagDerArbeit is Labour Day, a fixed date.
func TagDerArbeit(year int) Feiertag {
	return Feiertag{time.Date(year, time.May, 1, 0, 0, 0, 0, getTimeZone()), "Tag der Arbeit", []Region{}}
}

// Staatsfeiertag is May 1st in Austria, a fixed date.
func Staatsfeiertag(year int) Feiertag {
	e := TagDerArbeit(year)
	e.Text = "Staatsfeiertag"
	return e
}

// Florianitag is St Florian's Day, a fixed date.
func Florianitag(year int) Feiertag {
	return Feiertag{time.Date(year, time.May, 4, 0, 0, 0, 0, getTimeZone()), "Florianitag", []Region{}}
}

// TagDerBefreiung is Victory in Europe Day, a fixed date.
func TagDerBefreiung(year int) Feiertag {
	return Feiertag{time.Date(year, time.May, 8, 0, 0, 0, 0, getTimeZone()), "Tag der Befreiung", []Region{}}
}

// ChristiHimmelfahrt is Ascension Day, 39 days after Easter, therefore always a Thursday.
func ChristiHimmelfahrt(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, 39), "Christi Himmelfahrt", []Region{}}
}

// Pfingsten is Pentecost, 49 days after Easter.
func Pfingsten(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, 49), "Pfingsten", []Region{}}
}

// Pfingstmontag is Whit Monday, the monday after Pentecost.
func Pfingstmontag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, 50), "Pfingstmontag", []Region{}}
}

// Fronleichnam is Corpus Christi, 60 days after Eastern, therefore always a Thursday.
func Fronleichnam(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, 60), "Fronleichnam", []Region{}}
}

// MariäHimmelfahrt is Assumption Day, a fixed date.
func MariäHimmelfahrt(year int) Feiertag {
	return Feiertag{time.Date(year, time.August, 15, 0, 0, 0, 0, getTimeZone()), "Mariä Himmelfahrt", []Region{}}
}

// Rupertitag is St Rupert's Day, a fixed date.
func Rupertitag(year int) Feiertag {
	return Feiertag{time.Date(year, time.September, 24, 0, 0, 0, 0, getTimeZone()), "Rupertitag", []Region{}}
}

// TagDerDeutschenEinheit is German Unity Day, a fixed date.
func TagDerDeutschenEinheit(year int) Feiertag {
	return Feiertag{time.Date(year, time.October, 3, 0, 0, 0, 0, getTimeZone()), "Tag der deutschen Einheit", []Region{}}
}

// TagDerVolksabstimmung is Referendum Day in Carinthia, a fixed date.
func TagDerVolksabstimmung(year int) Feiertag {
	return Feiertag{time.Date(year, time.October, 10, 0, 0, 0, 0, getTimeZone()), "Tag der Volksabstimmung", []Region{}}
}

// Nationalfeiertag is the Austrian national day, a fixed date.
func Nationalfeiertag(year int) Feiertag {
	return Feiertag{time.Date(year, time.October, 26, 0, 0, 0, 0, getTimeZone()), "Nationalfeiertag", []Region{}}
}

// Reformationstag is Reformation Day, a fixed date.
func Reformationstag(year int) Feiertag {
	return Feiertag{time.Date(year, time.October, 31, 0, 0, 0, 0, getTimeZone()), "Reformationstag", []Region{}}
}

// Allerheiligen is All Saints' Day or Allhallows, a fixed date
func Allerheiligen(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 1, 0, 0, 0, 0, getTimeZone()), "Allerheiligen", []Region{}}
}

// Martinstag or Skt. Martin is Martinmas, a fixed date
func Martinstag(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 11, 0, 0, 0, 0, getTimeZone()), "Martinstag", []Region{}}
}

// Leopolditag is St Leopold's Day, a fixed date.
func Leopolditag(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 15, 0, 0, 0, 0, getTimeZone()), "Leopolditag", []Region{}}
}

// Weltkindertag is World Children's Day, a fixed date.
func Weltkindertag(year int) Feiertag {
	return Feiertag{time.Date(year, time.September, 20, 0, 0, 0, 0, getTimeZone()), "Weltkindertag", []Region{}}
}

// BußUndBettag is Penance Day, 11 days before the first Sunday in Advent
func BußUndBettag(year int) Feiertag {
	o := time.Date(year, time.November, 22, 0, 0, 0, 0, getTimeZone())
	d := (4 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Buß- und Bettag", []Region{}}
}

// MariäEmpfängnis has a shorter name for MariäUnbefleckteEmpfängnis in Austria.
func MariäEmpfängnis(year int) Feiertag {
	e := MariäUnbefleckteEmpfängnis(year)
	e.Text = "Mariä Empfängnis"
	return e
}

// Weihnachten is Christmas, a fixed date
func Weihnachten(year int) Feiertag {
	return Feiertag{time.Date(year, time.December, 25, 0, 0, 0, 0, getTimeZone()), "Weihnachten", []Region{}}
}

// Christtag is Christmas is  in Austria.
func Christtag(year int) Feiertag {
	e := Weihnachten(year)
	e.Text = "Christtag"
	return e
}

// ZweiterWeihnachtsfeiertag is day after Christmas, a fixed date
func ZweiterWeihnachtsfeiertag(year int) Feiertag {
	return Feiertag{time.Date(year, time.December, 26, 0, 0, 0, 0, getTimeZone()), "Zweiter Weihnachtsfeiertag", []Region{}}
}

// Stefanitag is December 26th in Austria.
func Stefanitag(year int) Feiertag {
	e := ZweiterWeihnachtsfeiertag(year)
	e.Text = "Stefanitag"
	return e
}


