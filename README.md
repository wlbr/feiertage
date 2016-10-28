# Feiertage
Feiertage is a Go/Golang library for calculating german bank holidays. It includes the calculation of the date of Easter and, more important, offers ways to retrieve the public holiday for a state of Germany (=Bundesland).

The library is probably useful only for people realizing use cases with special "german" requirements. This might be shift schedules or capacity calculation.

### Documentation
See https://godoc.org/github.com/wlbr/feiertage

### Usage:
There are to types of functions: 

  * `<feiertag>(year)` and 
  * `<region>(year optional:IncludingSundays:true)`

`<feiertag>` returns an extended `time` object (type `feiertag`). It carrys the concrete date plus the name of the holiday.
`<feiertag>` may be any of 

||||
|----|-----|----|
`Neujahr` | `Epiphanias` | `HeiligeDreiKönige` 
`Valentinstag` | `Weiberfastnacht` | `Karnevalssonntag` 
`Rosenmontag` | `Fastnacht` | `Aschermittwoch` 
`Palmsonntag` | `Gründonnerstag` | `Karfreitag` 
`Ostern` | `BeginnSommerzeit` | `Ostermontag` 
`Walpurgisnacht` | `TagDerArbeit` | `TagDerBefreiung` 
`Muttertag` | `ChristiHimmelfahrt` | `Vatertag` 
`Pfingsten` | `PfingstMontag` | `Dreifaltigkeitssonntag` 
`Fronleichnam` | `MariäHimmelfahrt` | `TagDerDeutschenEinheit` 
`Erntedankfest` | `Reformationstag` | `Halloween` 
`BeginnWinterzeit` | `Allerheiligen` | `Allerseelen` 
`Martinstag` | `Karnevalsbeginn` | `BußUndBettag`
`Thanksgiving` | `Blackfriday` | `Volkstrauertag` 
`Nikolaus` | `MariäUnbefleckteEmpfängnis` | `Totensonntag` 
`ErsterAdvent` | `ZweiterAdvent` | `DritterAdvent` 
`VierterAdvent` | `Heiligabend` | `Weihnachten` 
`ZweiterWeihnachten` | `Silvester` | &nbsp; 

`<region>` returns an object of type `region`. It offers a list of public holidays valid in the referred state` | `the name and the shortname of the state as attributes.
`<region>` may be any of 

||||
----|-----|----
`BadenWürttemberg` | `Bayern` | `Berlin` 
`Brandenburg` | `Bremen` | `Hamburg` 
`Hessen` | `MecklenburgVorpommern` | `Niedersachsen` 
`NordrheinWestfalen` | `RheinlandPfalz` | `Saarland` 
`Sachsen` | `SachsenAnhalt` | `SchleswigHolstein` 
`Thüringen` | `Deutschland` | `All`

The optional region functions second argument `includingSundays` switches the behavior o the region function, so that "gesetzliche Feiertage" on Sundays are included or not. This is important in Brandenburg and refers to Easter and Pentecost sunday. If you are calculating shift costs you will need to know even the holidays "hidden by sundays".

The region functions return the public holidays ("gesetzliche Feiertage"). The function `all` instead returns all the defined "special dates" as well. For example the Penance Day (Buß- und Bettag) or the begin/end of daylight saving time.



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

A little command line tool is included as well. It can be compiled using `make buildcmd` or `go build cmd/feiertage/feiertage.go` This will create an executable `feiertage`. See https://github.com/wlbr/feiertage/releases/latest for latest releases.

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
&nbsp;All</dd>
</dl>


## Code
* Documentation: https://godoc.org/github.com/wlbr/feiertage
* Lint: http://go-lint.appspot.com/github.com/wlbr/feiertage
* Continous Integration: [![Travis Status](https://api.travis-ci.org/wlbr/feiertage.svg?branch=master)](https://travis-ci.org/wlbr/feiertage)
* Test Coverage: [![Coverage Status](https://coveralls.io/repos/github/wlbr/feiertage/badge.svg?branch=master)](https://coveralls.io/github/wlbr/feiertage?branch=master)
* Metrics: [![GoReportCard](https://goreportcard.com/badge/github.com/wlbr/feiertage)](https://goreportcard.com/report/github.com/wlbr/feiertage)

