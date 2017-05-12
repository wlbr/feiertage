package main

import (
	"flag"
	"fmt"
	"github.com/uffish/feiertage"
	"strconv"
	"strings"
)

func fmtTaskjuggler(reg feiertage.Region) string {
	s := fmt.Sprintf("# public holidays for %s (%s)", reg.Name, reg.Shortname)
	for _, f := range reg.Feiertage {
		s = fmt.Sprintf("%s,\n leaves holiday \"%s\" %s", s, f.Text, f.Format("2006-01-02"))
	}
	return s
}

func getRegion(region string, year int, includingSundays bool) feiertage.Region {
	rs := strings.ToLower(region)
	rep := strings.NewReplacer("-", "", "ä", "ae", "ö", "oe", "ü", "ue", "ß", "ss")
	rs = rep.Replace(rs)
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
	case "burgenland":
		r = feiertage.Burgenland(year, includingSundays)
	case "kaernten":
		r = feiertage.Kärnten(year, includingSundays)
	case "niederoesterreich":
		r = feiertage.Niederösterreich(year, includingSundays)
	case "oberoesterreich":
		r = feiertage.Oberösterreich(year, includingSundays)
	case "salzburg":
		r = feiertage.Salzburg(year, includingSundays)
	case "steiermark":
		r = feiertage.Steiermark(year, includingSundays)
	case "tirol":
		r = feiertage.Tirol(year, includingSundays)
	case "vorarlberg":
		r = feiertage.Vorarlberg(year, includingSundays)
	case "wien":
		r = feiertage.Wien(year, includingSundays)
	case "oesterreich":
		r = feiertage.Österreich(year, includingSundays)
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
		"\t\tBurgenland\n"+
		"\t\tKärnten\n"+
		"\t\tNiederösterreich\n"+
		"\t\tOberösterreich\n"+
		"\t\tSalzburg\n"+
		"\t\tSteiermark\n"+
		"\t\tTirol\n"+
		"\t\tVorarlberg\n"+
		"\t\tWien\n"+
		"\t\tÖsterreich\n"+
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
