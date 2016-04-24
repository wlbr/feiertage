package main

import (
	"feiertage"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func fmtTaskjuggler(reg feiertage.Region) string {
	s := fmt.Sprintf("// publich holidays for %s (%s)", reg.Name, reg.Shortname)
	for _, f := range reg.Feiertage {
		s = fmt.Sprintf("%s \n leaves holiday \"%s\" 2011-12-24 +3d", s, f)
	}
	return s
}

func getRegion(region string, year int, includingSundays bool) feiertage.Region {
	rs := strings.ToLower(region)
	rep := strings.NewReplacer("-", "", "ä", "ae", "ö", "oe", "ü", "ue", "ß", "ss")
	rs = rep.Replace(rs)
	fmt.Println(includingSundays)
	var r feiertage.Region
	switch rs {
	case "badenwuerttemberg":
		r = feiertage.BadenWürttemberg(year, includingSundays)
	case "bayern":
		r = feiertage.Bayern(year, includingSundays)
	case "berlin":
		r = feiertage.Berlin(year, includingSundays)
	case "brandenburg":
		r = feiertage.Brandenburg(year, includingSundays)
	case "bremen":
		r = feiertage.Bremen(year, includingSundays)
	case "hamburg":
		r = feiertage.Hamburg(year, includingSundays)
	case "hessen":
		r = feiertage.Hessen(year, includingSundays)
	case "mecklenburgvorpommern":
		r = feiertage.MecklenburgVorpommern(year, includingSundays)
	case "niedersachsen":
		r = feiertage.Niedersachsen(year, includingSundays)
	case "nordrheinwestfalen":
		r = feiertage.NordrheinWestfalen(year, includingSundays)
	case "rheinlandpfalz":
		r = feiertage.RheinlandPfalz(year, includingSundays)
	case "saarland":
		r = feiertage.Saarland(year, includingSundays)
	case "sachsen":
		r = feiertage.Sachsen(year, includingSundays)
	case "sachsenanhalt":
		r = feiertage.SachsenAnhalt(year, includingSundays)
	case "schleswigholstein":
		r = feiertage.SchleswigHolstein(year, includingSundays)
	case "thueringen":
		r = feiertage.Thüringen(year, includingSundays)
	case "deutschland":
		r = feiertage.Deutschland(year, includingSundays)
	default:
		r = feiertage.All(year, includingSundays)
	}
	return r
}

func main() {
	var region = flag.String("region", "all", "Feiertag für Region 'string'.\n"+
		"\tRegion kann sein:\n"+
		"\t\tBadenWürttemberg\n"+
		"\t\tBayern\n"+
		"\t\tBerlin\n"+
		"\t\tBrandenburg\n"+
		"\t\tBremen\n"+
		"\t\tHamburg\n"+
		"\t\tHessen\n"+
		"\t\tMecklenburgVorpommern\n"+
		"\t\tNiedersachsen\n"+
		"\t\tNordrheinWestfalen\n"+
		"\t\tRheinlandPfalz\n"+
		"\t\tSaarland\n"+
		"\t\tSachsen\n"+
		"\t\tSachsenAnhalt\n"+
		"\t\tSchleswigHolstein\n"+
		"\t\tThüringen\n"+
		"\t\tDeutschland\n"+
		"\t\tAll\n")
	var includingSundays = flag.Bool("inklusiveSonntage", false, "Sollen Feiertag an Sonntagen mit ausgegeben werden?")
	var asTaskjugglerCode = flag.Bool("asTaskjugglerCode", false, "Taskjuggler Code ausgeben.")
	flag.Parse()
	if len(flag.Args()) > 0 {
		year, err := strconv.Atoi(flag.Args()[0])
		if err != nil {
			fmt.Println("Jahr muss eine Zahl sein.")
		} else {
			reg := getRegion(*region, year, *includingSundays)
			if *asTaskjugglerCode {
				fmt.Println(fmtTaskjuggler(reg))
			} else {
				fmt.Println(reg)
			}
		}
	} else {
		fmt.Println("Kein Jahr angegeben.")
	}
}
