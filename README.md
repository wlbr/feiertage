# Feiertage
Feiertage is a Go/Golang library for calculating german bank holidays. It includes the calculation of the date of Easter and, more important, offers ways to retrieve the public holiday for a state of Germany (=Bundesland).

The library is probably only useful for people realizing usecases with special "german" requirements. This might be shift schedules or capacity calculation.

### Usage:
There are to types of functions: 

  * `<feiertag>(year)` and 
  * `<region>(year optional:IncludingSundays:true)`

`<feiertag>` returns an extended `time` object (type `feiertag`). It carrys the concrete date plus the name of the holiday.
`<feiertag>` may be any of 

|||
----|-----|----
| `Neujahr` | `Epiphanias` | `HeiligeDreiKönige` |
| `Valentinstag` | `Weiberfastnacht` | `Karnevalssonntag` |
| `Rosenmontag` | `Fastnacht` | `Aschermittwoch` |
| `Palmsonntag` | `Gründonnerstag` | `Karfreitag` |
| `Ostern` | `BeginnSommerzeit` | `Ostermontag` |
| `Walpurgisnacht` | `TagDerArbeit` | `TagDerBefreiung` |
| `Muttertag` | `ChristiHimmelfahrt` | `Vatertag` |
| `Pfingsten` | `PfingstMontag` | `Dreifaltigkeitssonntag` |
| `Fronleichnam` | `MariäHimmelfahrt` | `TagDerDeutschenEinheit` |
| `Erntedankfest` | `Reformationstag` | `Halloween` |
| `BeginnWinterzeit` | `Allerheiligen` | `Allerseelen` |
| `Martinstag` | `Karnevalsbeginn` | `BußUndBettag` |
| `Volkstrauertag` | `Nikolaus` | `MariäUnbefleckteEmpfängnis` |
| `Totensonntag` | `ErsterAdvent` | `ZweiterAdvent` |
| `DritterAdvent` | `VierterAdvent` | `Heiligabend` |
| `Weihnachten` | `ZweiterWeihnachtsfeiertag` | `Silvester` |

`<region>` returns an object of type `region`. It offers a list of public holidays valid in the referred state` | `the name and the shortname of the state as attributes.
`<region>` may be any of 

|||
----|-----|----
`BadenWürttemberg` | `Bayern` | `Berlin` 
`Brandenburg` | `Bremen` | `Hamburg` 
`Hessen` | `MecklenburgVorpommern` | `Niedersachsen` 
`NordrheinWestfalen` | `RheinlandPfalz` | `Saarland` 
`Sachsen` | `SachsenAnhalt` | `SchleswigHolstein` 
`Thüringen` | `Deutschland` | `All`

The optional second argument `includingSundays` of the region functions switches the behavior, so that "gesetzliche Feiertage" on Sundays are incuded or not. This is important in Brandenburg and refers to Easter and Pentecost sunday. If you are calculation shift costs you will need to know even the holidays "hidden by sundays".

The region functions return the public holidays ("gesetzliche Feiertage"). The function `all`,  instead returns all the defined "special dates" as well. For example the Penance Day or the begin/end of daylight saving time.



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

    
    fmt.Println(Brandenburg(2016` | `false))
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

