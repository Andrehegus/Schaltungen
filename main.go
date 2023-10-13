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

const (
	Schalter1 bool = false
	Schalter2
)

func Wechselschaltung() bool {
	array := [2]bool{Schalter1, Schalter2}

	// Schalter 1
	array[0] = false // Schalter1
	array[1] = false // Schalter2

	// Logic
	// Todo Logic besser zusammenfassen, schauen ob es besser geht...
	if array[0] && array[1] {
		return false
	}
	if array[0] && !array[1] {
		return true
	}
	if (!array[0]) && array[1] {
		return true
	}
	return false
}

func main() {
	fmt.Println("a) Wechselschaltung...")

	strLicht := "aus"

	licht := Wechselschaltung()
	if licht {
		strLicht = "an"
	} else {
		strLicht = "aus"
	}
	fmt.Println("Licht ist", strLicht)
}
