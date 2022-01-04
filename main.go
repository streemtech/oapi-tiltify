package main

import (
	"fmt"
	"time"

	"github.com/antzucaro/matchr"
)

var testMessages = []string{
	"Love the cause, and all you guys are Working for. Here's to 200k. Money goes to Twilight Princess",
	"Hey y'all, thanks for supporting a great cause! This donation is for Battle Kid: Fortress of Peril.",
	"Hey guys. Love this Direct Reliefing that you've been doing. Let's smash that 500k ceiling!!! Also my donation's for Action Girlz Racing",
	"Thanks for the memories and supporting a great cause! My donation goes to Pokemon Heart Gold. Take care everyone",
	"Konbanwa! As the saying goes, every dollar to Direct Relief matters. Let's break another record this year!	Oh, and as for the game, we need some chaos here and there, so put my money on High-Risk Fortune Cookie!",
	"hey Jon. You know what it fucking is. Darksiders II Deathinitive Edition on PS4. I'll keep doing this until you beat the game on stream I swear to FU-",
}

func main() {
	for _, v := range testMessages {
		testMessage(v)
	}
}

func testMessage(message string) {
	var dlt time.Duration
	var levt time.Duration
	var osat time.Duration
	for _, v := range testGames {
		c := compareStrings(v, message)

		dlt += c.dlTime
		levt += c.levTime
		osat += c.osaTime
		if c.dlRes != c.levRes || c.levRes != c.osaRes {

			fmt.Printf("result mismatch:\n%s", c.String())
		}
	}

	fmt.Printf("dlt: %s\nlevt: %s\nosat: %s\n\n", dlt.String(), levt.String(), osat.String())
}

var testGames = []string{
	"The Legend of Zelda: Twilight Princess (GCN)",
	"The Legend of Zelda: Wind Waker (GCN)",
	"The Legend of Zelda: Minish Cap (GBA)",
	"Paper Mario: Color Splash (WiiU)",
	"Paper Mario: The Origami King (Switch)",
	"Paper Mario: The Thousand Year Door (GCN)",
	"Paper Mario (N64)",
	"Timespinner (NSDL)",
	"VVVVVV (Steam)",
	"Superman 64 (N64)",
	"Battle Kid: Fortress of Peril",
	"Action Girlz Racing",
	"Cookie, Fortune Cookie",
	"Darksiders II Deathinitive Edition",
	"Pokemon Heart Gold",
}

func compareStrings(s1, s2 string) res {
	r := res{
		s1: s1,
		s2: s2,
	}
	n := time.Now()
	r.dlRes = matchr.DamerauLevenshtein(s1, s2)
	r.dlTime = time.Since(n)
	n = time.Now()

	r.jaRes = matchr.Jaro(s1, s2)
	r.jaTime = time.Since(n)
	n = time.Now()

	r.levRes = matchr.Levenshtein(s1, s2)
	r.levTime = time.Since(n)
	n = time.Now()

	r.osaRes = matchr.OSA(s1, s2)
	r.osaTime = time.Since(n)
	return r
}

type res struct {
	s1, s2  string
	dlTime  time.Duration
	dlRes   int
	jaTime  time.Duration
	jaRes   float64
	levTime time.Duration
	levRes  int
	osaTime time.Duration
	osaRes  int
}

func (r res) String() string {
	return fmt.Sprintf("s1: %s\ns2: %s\nDL : %d, %s\nJAR: %f, %s\nLEV: %d, %s\nOSA: %d, %s\n\n",
		r.s1, r.s2,
		r.dlRes, r.dlTime.String(),
		r.jaRes, r.jaTime.String(),
		r.levRes, r.levTime.String(),
		r.osaRes, r.osaTime.String())
}
