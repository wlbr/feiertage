// Copyright (c) 2026 Michael Wolber
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/wlbr/feiertage"
)

func fmtTaskjuggler(reg feiertage.Region) string {
	s := fmt.Sprintf("# public holidays for %s (%s)", reg.Name, reg.Shortname)
	for _, f := range reg.Feiertage {
		s = fmt.Sprintf("%s,\n leaves holiday \"%s\" %s", s, f.Text, f.Format("2006-01-02"))
	}
	return s
}

func canonicalize(in string) (out string) {
	low := strings.ToLower(in)
	rep := strings.NewReplacer("-", "", "ä", "ae", "ö", "oe", "ü", "ue", "ß", "ss")
	return rep.Replace(low)

}

func getRegion(region string, year int, includingSundays bool) (feiertage.Region, error) {
	rs := canonicalize(region)
	var r feiertage.Region

	allRegions := feiertage.GetAllRegions(year, includingSundays)
	for _, r := range allRegions {
		if canonicalize(r.Name) == rs || canonicalize(r.Shortname) == rs {
			return r, nil
		}
	}
	return r, fmt.Errorf("region '%s' unbekannt", region)
}

func main() {

	var regions string
	for _, r := range feiertage.GetAllRegions(time.Now().Year(), false, "de") {
		regions = regions + "\t" + r.Name + "\n"
	}
	for _, r := range feiertage.GetAllRegions(time.Now().Year(), false, "at") {
		regions = regions + "\t" + r.Name + "\n"
	}
	regions = regions + "\tAlle\n"
	var region = flag.String("region", "Alle", "Feiertag für Region 'string'.\n"+
		"Region kann sein:\n"+regions)

	var includingSundays = flag.Bool("inklusiveSonntage", false, "Sollen Feiertag an Sonntagen mit ausgegeben werden?")
	var asTaskjugglerCode = flag.Bool("asTaskjugglerCode", false, "Taskjuggler Code ausgeben.")
	var asJSON = flag.Bool("asJSON", false, "JSON Code ausgeben.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n"+
			"  %s [options] year \n\n"+
			"%s zeigt alle Feiertage eins übergebenen Jahres an.\n\n"+
			"Options:\n", os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if len(flag.Args()) > 0 {
		year, err := strconv.Atoi(flag.Args()[0])
		if err != nil {
			log.Println("Jahr muss eine Zahl sein.")
		} else {
			reg, e := getRegion(*region, year, *includingSundays)
			if e != nil {
				fmt.Println(e)
			} else if *asTaskjugglerCode {
				fmt.Println(fmtTaskjuggler(reg))
			} else if *asJSON {
				jsonR, err := json.Marshal(reg)
				if err != nil {
					log.Println(err)
				}
				fmt.Printf("%s", jsonR)
			} else {
				fmt.Println(reg)
			}
		}
	} else {
		fmt.Println("Kein Jahr angegeben.")
	}
}
