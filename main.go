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

func Wechselschaltung(s1 bool, s2 bool) bool {
	// Logic
	// Todo Logic besser zusammenfassen, schauen ob es besser geht...
	if s1 && s2 {
		return false
	}
	if s1 && !s2 {
		return true
	}
	if (!s1) && s2 {
		return true
	}
	return false
}

func main() {
	fmt.Println("a) Wechselschaltung...")

	strLicht := "aus"
	s1 := false
	s2 := false

	licht := Wechselschaltung(s1, s2)
	if licht {
		strLicht = "an"
	} else {
		strLicht = "aus"
	}
	fmt.Println("Licht ist", strLicht)
}
