package feiertage

import "time"

// Weltknuddeltag is World Hug Day or National Hugging Day, a fixed date.
func Weltknuddeltag(year int) Feiertag {
	return Feiertag{time.Date(year, time.January, 21, 0, 0, 0, 0, getTimeZone()), "Weltknuddeltag", []Region{}}
}

// StarWarsDay is a fixed date.
func StarWarsDay(year int) Feiertag {
	return Feiertag{time.Date(year, time.May, 4, 0, 0, 0, 0, getTimeZone()), "Star Wars Day", []Region{}}
}

// Handtuchtag is Towel Day, May 25. It is celebrated as a tribute to the author Douglas Adams by his fans.
func Handtuchtag(year int) Feiertag {
	return Feiertag{time.Date(year, time.May, 25, 0, 0, 0, 0, getTimeZone()), "Handtuchtag", []Region{}}
}

// TowelDay is, May 25. It is celebrated as a tribute to the author Douglas Adams by his fans.
func TowelDay(year int) Feiertag {
	e := Handtuchtag(year)
	e.Text = "Towel Day"
	return e
}

// Weltumwelttag is World Environment Day, a fixed date.
func Weltumwelttag(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 5, 0, 0, 0, 0, getTimeZone()), "Weltumwelttag", []Region{}}
}

// Weltspieltag is International Day of Play, a fixed date.
func Weltspieltag(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 11, 0, 0, 0, 0, getTimeZone()), "Weltspieltag", []Region{}}
}

// Weltblutspendetag is World Blood Donor Day, a fixed date.
func Weltblutspendetag(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 14, 0, 0, 0, 0, getTimeZone()), "Weltblutspendetag", []Region{}}
}

// FêteDeLaMusique is World Music Day, a fixed date.
func FêteDeLaMusique(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 21, 0, 0, 0, 0, getTimeZone()), "Fête de la Musique", []Region{}}
}

// InternationalerTagGegenDrogenmissbrauch is International Day Against Drug Abuse and Illicit Trafficking, a fixed date.
func InternationalerTagGegenDrogenmissbrauch(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 26, 0, 0, 0, 0, getTimeZone()), "Internationaler Tag gegen Drogenmissbrauch", []Region{}}
}

// SystemAdministratorAppreciationDay is the last Fridy in July
func SystemAdministratorAppreciationDay(year int) Feiertag {
	o := time.Date(year, time.July, 31, 0, 0, 0, 0, getTimeZone())
	d := (2 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "System Administrator Appreciation Day", []Region{}}
}

// Hobbit Day is a fixed date.
func HobbitDay(year int) Feiertag {
	return Feiertag{time.Date(year, time.September, 22, 0, 0, 0, 0, getTimeZone()), "Hobbit Day", []Region{}}
}

// InternationalerMännertag is International Day Men's Day, a fixed date
func InternationalerMännertag(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 19, 0, 0, 0, 0, getTimeZone()), "Internationaler Männertag", []Region{}}
}

// Karnevalsbeginn is the beginning of carnival, a fixed date.
func Karnevalsbeginn(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 11, 11, 11, 11, 11, getTimeZone()), "Karnevalsbeginn", []Region{}}
}

// Thanksgiving in the US, the fourth Thursday of November.
func Thanksgiving(year int) Feiertag {
	o := time.Date(year, time.November, 1, 0, 0, 0, 0, getTimeZone())
	d := ((11 - int(o.Weekday())) % 7)
	return Feiertag{o.AddDate(0, 0, 21+d), "Thanksgiving (US)", []Region{}}
}

// Blackfriday is the Friday after Thanksgiving.
func Blackfriday(year int) Feiertag {
	return Feiertag{Thanksgiving(year).AddDate(0, 0, 1), "Blackfriday", []Region{}}
}

// Valentinstag is Valentine's Day, a fixed date.
func Valentinstag(year int) Feiertag {
	return Feiertag{time.Date(year, time.February, 14, 0, 0, 0, 0, getTimeZone()), "Valentinstag", []Region{}}
}

// InternationalerTagDesGedenkensAnDieOpferDesHolocaust is (International Holocaust Remembrance Day, a fixed date.
func InternationalerTagDesGedenkensAnDieOpferDesHolocaust(year int) Feiertag {
	return Feiertag{time.Date(year, time.January, 27, 0, 0, 0, 0, getTimeZone()), "Internationaler Tag des Gedenkens an die Opfer des Holocaust", []Region{}}
}

// Weiberfastnacht is a part of carnival, 52 days before Easter.
func Weiberfastnacht(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -52), "Weiberfastnacht", []Region{}}
}

// Karnevalssonntag is the sunday of carnival, 49 days before Easter.
func Karnevalssonntag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -49), "Karnevalssonntag", []Region{}}
}

// Rosenmontag is the monday of carnival, 48 days before Easter.
func Rosenmontag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -48), "Rosenmontag", []Region{}}
}

// Fastnacht is shrovetide, the Tuesday of carnival, 47 days before Easter.
func Fastnacht(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -47), "Fastnacht", []Region{}}
}

// Aschermittwoch is Ash Wednesday, 46 days before Easter.
func Aschermittwoch(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -46), "Aschermittwoch", []Region{}}
}

// Palmsonntag is Palm Sunday , the last Sunday before Easter
func Palmsonntag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -7), "Palmsonntag", []Region{}}
}

// Gründonnerstag is Holy Thursday or Maundy Thursday, the last Thursday before Eastern
func Gründonnerstag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, -3), "Gründonnerstag", []Region{}}
}

// BeginnSommerzeit is the start of daylight saving time. Last Sunday of March.
func BeginnSommerzeit(year int) Feiertag {
	o := time.Date(year, time.March, 31, 0, 0, 0, 0, getTimeZone())
	d := (7 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Beginn Sommerzeit", []Region{}}
}

// TagDerErde is Earth Day, a fixed date.
func TagDerErde(year int) Feiertag {
	return Feiertag{time.Date(year, time.April, 22, 0, 0, 0, 0, getTimeZone()), "Tag der Erde", []Region{}}
}

// Walpurgisnacht is Walpurgis Night, a fixed date.
func Walpurgisnacht(year int) Feiertag {
	return Feiertag{time.Date(year, time.April, 30, 0, 0, 0, 0, getTimeZone()), "Walpurgisnacht", []Region{}}
}

// InternationalerTagDerPressefreiheit is World Press Freedom Day, a fixed date.
func InternationalerTagDerPressefreiheit(year int) Feiertag {
	return Feiertag{time.Date(year, time.May, 3, 0, 0, 0, 0, getTimeZone()), "Internationaler Tag der Pressefreiheit", []Region{}}
}

// Muttertag is Mother's Day oder Mothering Sunday, the second Sunday in May.
func Muttertag(year int) Feiertag {
	o := time.Date(year, time.May, 1, 0, 0, 0, 0, getTimeZone())
	d := (7 - int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, d+7), "Muttertag", []Region{}}
}

// Vatertag is Father's Day, same day a Ascension Day, 39 days after Easter, therefore always a Thursday.
func Vatertag(year int) Feiertag {
	e := ChristiHimmelfahrt(year)
	e.Text = "Vatertag"
	return e
}

// Dreifaltigkeitssonntag is Trinity Sunday, the Sunday after Pentecost
func Dreifaltigkeitssonntag(year int) Feiertag {
	o := Ostern(year)
	return Feiertag{o.AddDate(0, 0, 56), "Dreifaltigkeitssonntag", []Region{}}
}

// InternationalerKindertag is special to Germany and Austrian and
// isnot the same as Weltkindertag (World Children's Day), a fixed date.
func InternationalerKindertag(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 1, 0, 0, 0, 0, getTimeZone()), "Internationaler Kindertag", []Region{}}
}

// TagDesMeeres is World Oceans Day, a fixed date.
func TagDesMeeres(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 8, 0, 0, 0, 0, getTimeZone()), "Tag des Meeres", []Region{}}
}

// Weltflüchtlingstag is World Refugee Day, a fixed date.
func Weltflüchtlingstag(year int) Feiertag {
	return Feiertag{time.Date(year, time.June, 20, 0, 0, 0, 0, getTimeZone()), "Weltflüchtlingstag", []Region{}}
}

// Antikriegstag is Anti-War Day, a fixed date.
func Antikriegstag(year int) Feiertag {
	return Feiertag{time.Date(year, time.September, 1, 0, 0, 0, 0, getTimeZone()), "Antikriegstag", []Region{}}
}

// Halloween is a fixed date.
func Halloween(year int) Feiertag {
	return Feiertag{time.Date(year, time.October, 31, 0, 0, 0, 0, getTimeZone()), "Halloween", []Region{}}
}

// BeginnWinterzeit is the end of daylight saving time. Last Sunday of October.
func BeginnWinterzeit(year int) Feiertag {
	o := time.Date(year, time.October, 31, 0, 0, 0, 0, getTimeZone())
	d := (7 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Beginn Winterzeit", []Region{}}
}

// Allerseelen is All Souls' Day, the day after All Saints' Day,
func Allerseelen(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 2, 0, 0, 0, 0, getTimeZone()), "Allerseelen", []Region{}}
}

// Weltmännertag is Men's World Day.
func Weltmännertag(year int) Feiertag {
	return Feiertag{time.Date(year, time.November, 3, 0, 0, 0, 0, getTimeZone()), "Weltmännertag", []Region{}}
}

// Erntedankfest is Thanksgiving or Harvest Festival, the first Sunday of October.
// The german Erntedankfest is not the same than the US Thanksgiving.
func Erntedankfest(year int) Feiertag {
	o := time.Date(year, time.October, 1, 0, 0, 0, 0, getTimeZone())
	d := (7 - int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, d), "Erntedankfest", []Region{}}
}

// Nikolaus is St Nicholas' Day, a fixed date
func Nikolaus(year int) Feiertag {
	return Feiertag{time.Date(year, time.December, 6, 0, 0, 0, 0, getTimeZone()), "Nikolaus", []Region{}}
}

// MariäUnbefleckteEmpfängnis is Day of Immaculate Conception, a fixed date.
func MariäUnbefleckteEmpfängnis(year int) Feiertag {
	return Feiertag{time.Date(year, time.December, 8, 0, 0, 0, 0, getTimeZone()), "Mariä unbefleckte Empfängnis", []Region{}}
}

// Volkstrauertag is Remembrance Sunday, the second sunday before the first Sunday in Advent
func Volkstrauertag(year int) Feiertag {
	o := ErsterAdvent(year)
	return Feiertag{o.AddDate(0, 0, -14), "Volkstrauertag", []Region{}}
}

// Totensontag is Sunday in commemoration of the dead, the last Sunday before the fourth Sunday in Advent
func Totensonntag(year int) Feiertag {
	o := VierterAdvent(year)
	return Feiertag{o.AddDate(0, 0, -28), "Totensonntag", []Region{}}
}

// ErsterAdvent is the first Sunday in Advent
func ErsterAdvent(year int) Feiertag {
	o := VierterAdvent(year)
	return Feiertag{o.AddDate(0, 0, -21), "Erster Advent", []Region{}}
}

// ZweiterAdvent is the second Sunday in Advent
func ZweiterAdvent(year int) Feiertag {
	o := VierterAdvent(year)
	return Feiertag{o.AddDate(0, 0, -14), "Zweiter Advent", []Region{}}
}

// DritterAdvent is the third Sunday in Advent
func DritterAdvent(year int) Feiertag {
	o := VierterAdvent(year)
	return Feiertag{o.AddDate(0, 0, -7), "Dritter Advent", []Region{}}
}

// VierterAdvent is the fourth Sunday in Advent
func VierterAdvent(year int) Feiertag {
	o := time.Date(year, time.December, 24, 0, 0, 0, 0, getTimeZone())
	d := (7 + int(o.Weekday())) % 7
	return Feiertag{o.AddDate(0, 0, -1*d), "Vierter Advent", []Region{}}
}

// Heiligabend is Christmas Eve, the last day before Christmas.
func Heiligabend(year int) Feiertag {
	return Feiertag{time.Date(year, time.December, 24, 0, 0, 0, 0, getTimeZone()), "Heiligabend", []Region{}}
}

// Silvester is NewYearsEve, a fixed date.
func Silvester(year int) Feiertag {
	return Feiertag{time.Date(year, time.December, 31, 0, 0, 0, 0, getTimeZone()), "Silvester", []Region{}}
}
