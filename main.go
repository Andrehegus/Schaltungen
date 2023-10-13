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
	Input bool = false
	OutputLeft
	OutputRight
)

type Schalter_1 int
type Schalter_2 int

func Wechselschaltung() {
	// Schalter 1
	Input = true
	OutputLeft = true
	OutputRight = false
}

func main() {
	fmt.Println("a) Wechselschaltung...")
	fmt.Println("Licht ist ")
}
