/*
Schaltungen abgebildet als Programm
- es gibt verschiedene Ein- und Ausgänge
- es gibt immer nur zwei Zustände am Ende: Ein oder Aus
*/

package main

import "fmt"

// a) Wechselschaltung
// 2 Schalter mit je 2 Zuständen: ein, aus / links, rechts ...
// jeder Schalter hat zwei Ausgänge
// jeder Schalter hat einen Eingang
// je nach Zustand wird z.B. Licht an / aus gehen
// die zwei Schalter brauchen zur Leuchte eine festgelegte verdratung

// Version a), mit 2 Schaltern Licht an/aus steuern
// Todo ... das gleiche mit (x) Schaltern programmieren
// Technisch würden zwischen den Wechselschaltern je ein oder mehrere Kreuzschalter hinzu gefügt...
// definition Kreuzschalter -> ein Kreuzschalter entspricht 2 Wechselschalter
func Wechselschaltung(ws1 bool, ws2 bool, ks bool) bool {
	// Logic Wechselschaltung
	// 1 && 0 = 1
	// 0 && 1 = 1
	// alles weitere = 0
	// entspricht der XOR Schalterfunktion!!!
	// ks3 ist der Kreuzschalter, der jeweils die anderen Schalterstellungen mit überprüft
	// ks = Kreuzschalter
	// ws1 = Wechselschalter 1
	// ws2 = Wechselschalter 2
	if !ks {
		if (ws1 && !ws2) || (!ws1 && ws2) {
			return true
		} else {
			return false
		}
	} else {
		if (ws1 && !ws2) || (!ws1 && ws2) {
			return false
		} else {
			return true
		}
	}
}

func main() {
	fmt.Println("a) Wechselschaltung...")

	// Hier die jeweiligen Schalter ein/aus schalten...
	Wechselschalter1 := false
	Wechselschalter2 := false
	Kreuzschalter := false
	licht := Wechselschaltung(Wechselschalter1, Wechselschalter2, Kreuzschalter)

	strLicht := "aus"
	if licht {
		strLicht = "an"
	} else {
		strLicht = "aus"
	}

	fmt.Println("Licht ist", strLicht)
}
