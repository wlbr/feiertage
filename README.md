# Feiertage
Feiertage is a Go/Golang library for calculating German and Austrian bank holidays. It includes the calculation of the date of Easter and, more importantly, offers ways to retrieve public holidays for a state of Germany or Austria (=Bundesland).

The library is probably useful only for people realizing use cases with special requirements inside of Austria or Germany, such as shift schedules or capacity calculation.

### Documentation
See https://godoc.org/github.com/wlbr/feiertage

### Usage:
There are two types of functions:

  * `<feiertag>(year)` and
  * `<region>(year optional:IncludingSundays:true)`

`<feiertag>` returns an extended `time` object (type `feiertag`). It carries the date of the holiday
in the requested year plus the name of the holiday. `<feiertag>` may be any of the following:

||||
|----|-----|----|
`Neujahr` | `Epiphanias` | `HeiligeDreiKönige`
`Valentinstag` | `Josefitag` | `Weiberfastnacht`
`Karnevalssonntag` | `Rosenmontag` | `Fastnacht`
`Aschermittwoch` | `Palmsonntag` | `Gründonnerstag`
`Karfreitag` | `Ostern` | `BeginnSommerzeit`
`Ostermontag` | `Walpurgisnacht` | `TagDerArbeit`
`TagDerBefreiung` | `Staatsfeiertag` | `Florianitag`
`Muttertag` | `ChristiHimmelfahrt` | `Vatertag`
`Pfingsten` | `Pfingstmontag` | `Dreifaltigkeitssonntag`
`Fronleichnam` | `MariäHimmelfahrt` | `Rupertitag`
`TagDerDeutschenEinheit` | `TagDerVolksabstimmung` | `Nationalfeiertag`
`Erntedankfest` | `Reformationstag` | `Halloween`
`BeginnWinterzeit` | `Allerheiligen` | `Allerseelen`
`Martinstag` | `Karnevalsbeginn` | `Leopolditag`
`BußUndBettag` | `Thanksgiving` | `Blackfriday`
`Volkstrauertag` | `Nikolaus` | `MariäUnbefleckteEmpfängnis`
`MariäEmpfängnis` | `Totensonntag` | `ErsterAdvent`
`ZweiterAdvent` | `DritterAdvent` | `VierterAdvent`
`Heiligabend` | `Weihnachten` | `Christtag`
`Stefanitag` | `ZweiterWeihnachtsfeiertag` | `Silvester`

`<region>` returns an object of type `region`. It offers a list of public holidays valid in the specified state as well as the name and the shortname of the state as attributes.
`<region>` may be any of:

||||
----|-----|----
`BadenWürttemberg` | `Bayern` | `Berlin`
`Brandenburg` | `Bremen` | `Hamburg`
`Hessen` | `MecklenburgVorpommern` | `Niedersachsen`
`NordrheinWestfalen` | `RheinlandPfalz` | `Saarland`
`Sachsen` | `SachsenAnhalt` | `SchleswigHolstein`
`Thüringen` | `Deutschland` | `Burgenland`
`Kärnten` | `Niederösterreich` | `Oberösterreich`
`Salzburg` | `Steiermark` | `Tirol`
`Vorarlberg` | `Wien` | `Österreich`
`All` | &nbsp; | &nbsp;

The optional region function argument `includingSundays` switches the behavior of the region function to include "gesetzliche Feiertage" that fall on Sundays in its output. This is important in Brandenburg, particularly for Easter and Pentecost Sunday. If you are calculating shift costs you will need to know even the holidays "hidden by Sunday".

The region functions return the public holidays ("gesetzliche Feiertage"). The function `all` returns all defined "special dates", such as Penance Day (Buß- und Bettag) or the begin/end of daylight saving time.

The regional functions for Austrian Bundesländer include saints' days which are state-level holidays, meaning
schools etc. are generally closed but workers don't get the day off by default. If you don't want to
include these days in your planning, it's okay to reference `Österreich` instead, as legal holidays are
(more or less) synchronised across all Austrian states (Bundesländer).

### Example:

    fmt.Println(Ostern(2016))
    --> 27.03.2016 Ostern


    fmt.Println(BußUndBettag(2016))
    --> 16.11.2016 Buß- und Bettag



    fmt.Println(Brandenburg(2016))
    --> Brandenburg (BB)
        01.01.2016 Neujahr
        25.03.2016 Karfreitag
        27.03.2016 Ostern
        28.03.2016 Ostermontag
        01.05.2016 Tag der Arbeit
        05.05.2016 Christi Himmelfahrt
        15.05.2016 Pfingsten
        16.05.2016 Pfingstmontag
        03.10.2016 Tag der deutschen Einheit
        31.10.2016 Reformationstag
        25.12.2016 Weihnachten
        26.12.2016 Zweiter Weihnachtsfeiertag


    fmt.Println(Brandenburg(2016, false))
    --> Brandenburg (BB)
        01.01.2016 Neujahr
        25.03.2016 Karfreitag
        28.03.2016 Ostermontag
        01.05.2016 Tag der Arbeit
        05.05.2016 Christi Himmelfahrt
        16.05.2016 Pfingstmontag
        03.10.2016 Tag der deutschen Einheit
        31.10.2016 Reformationstag
        25.12.2016 Weihnachten
        26.12.2016 Zweiter Weihnachtsfeiertag


## Command line tool

A little command line tool is included as well. It can be compiled using `make buildcmd` or `go build cmd/feiertage/feiertage.go` This will create an executable `feiertage`.

See https://github.com/wlbr/feiertage/releases/latest for downloads.

### Synopsis

`feiertage: [options] year`<br>
<dl>
<dt>-asTaskjugglerCode (default false)</dt>
<dd>Print the result as valid source code (`leave x y`) for the <a href="http://www.taskjuggler.org/">Taskjuggler</a> planning tool.
<dt>-inklusiveSonntage (default false)</dt>
<dd>Should public holidays on a Sunday be included?</dd>
<dt>-region &lt;regionstring&gt; (default "All")</dt>
<dd>Return public holidays for region `<regionstring>`.<br>
<dd>&lt;regionstring&gt; may be (case insensitive, plus some other tricks to make it more tolerant):<br>
&nbsp;BadenWürttemberg<br>
&nbsp;Bayern<br>
&nbsp;Berlin<br>
&nbsp;Brandenburg<br>
&nbsp;Bremen<br>
&nbsp;Hamburg<br>
&nbsp;Hessen<br>
&nbsp;MecklenburgVorpommern<br>
&nbsp;Niedersachsen<br>
&nbsp;NordrheinWestfalen<br>
&nbsp;RheinlandPfalz<br>
&nbsp;Saarland<br>
&nbsp;Sachsen<br>
&nbsp;SachsenAnhalt<br>
&nbsp;SchleswigHolstein<br>
&nbsp;Thüringen<br>
&nbsp;Deutschland<br>
&nbsp;Burgenland<br>
&nbsp;Kärnten<br>
&nbsp;Niederösterreich<br>
&nbsp;Oberösterreich<br>
&nbsp;Salzburg<br>
&nbsp;Steiermark<br>
&nbsp;Tirol<br>
&nbsp;Vorarlberg<br>
&nbsp;Wien<br>
&nbsp;Österreich<br>
&nbsp;All</dd>
</dl>


## Code
* Documentation: https://godoc.org/github.com/wlbr/feiertage
* Lint: http://go-lint.appspot.com/github.com/wlbr/feiertage
* Continuous Integration: [![Travis Status](https://api.travis-ci.com/wlbr/feiertage.svg?branch=master)](https://travis-ci.com/wlbr/feiertage)
* Test Coverage: [![Coverage Status](https://coveralls.io/repos/github/wlbr/feiertage/badge.svg?branch=master)](https://coveralls.io/github/wlbr/feiertage?branch=master)
* Metrics: [![GoReportCard](https://goreportcard.com/badge/github.com/wlbr/feiertage)](https://goreportcard.com/report/github.com/wlbr/feiertage)
