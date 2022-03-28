package main

import (
	"flag"
	"fmt"
	"strconv"
)

func callSignMaker(pierwszaImienia string, pierwszaNazwiska string, dzien int) string {
	var imie = map[string]string{"A": "Ananas", "B": "Borówka", "C": "Cytryna", "D": "Daktyl", "E": "Eukaliptus", "F": "Figa", "G": "Gruszka",
		"H": "Hibiskus", "I": "Imbir", "J": "Jabłko", "K": "Kiwi", "L": "Liczi", "M": "Mango", "N": "Nektarynka", "O": "Opuncja",
		"P": "Pomarańcza", "R": "Rambutan", "S": "Salak", "T": "Truskawka", "U": "Umkokolo", "W": "Winogrono", "Z": "Ziemniara"}

	var nazwisko = map[string]string{"A": "na antylopie", "B": "na biedronce", "C": "na czapli", "D": "na dziku", "E": "na emu",
		"F": "na foce", "G": "na gorylu", "H": "na hipopotamie", "I": "na iguanie", "J": "na jamniku", "K": "na kangurze",
		"L": "na lemurze", "M": "na małpie", "N": "na nietoperzu", "O": "na orce", "P": "na piesku", "R": "na rybie",
		"S": "na słoniu", "T": "na tygrysie", "U": "na uszatce", "W": "na wielbłądzie", "Z": "na zebrze"}

	var dzienMap = map[int]string{1: "w Sri Dźajawardanapura Kotte", 2: "w Pekinie", 3: "w Buenos Aires", 4: "w Warszawie", 5: "w Baku",
		6: "w Paryżu", 7: "w Malabo", 8: "w Hawanie", 9: "w Oslo", 10: "w Lizbonie", 11: "w Dili", 12: "w Kijowie", 13: "w Londynie",
		14: "w Tapeju", 15: "w Harare", 16: "w Rzymie", 17: "w Tunisie", 18: "w rękawiczkach", 19: "w bieliźnie", 20: "w sukience",
		21: "w spodniach", 22: "w koszulce", 23: "w sandałach", 24: "w spódnicy", 25: "w szaliku", 26: "w czapce", 27: "w tunice",
		28: "w rajstopach", 29: "w klapkach", 30: "w opasce", 31: "w okularach"}

	return imie[pierwszaImienia] + " " + nazwisko[pierwszaNazwiska] + " " + dzienMap[dzien]
}

func main() {
	var dzien int
	var namestring, surnamestring string

	flag.IntVar(&dzien, "day", 0, "dzień urodzenia")
	flag.StringVar(&namestring, "namechar", "", "Pierwsza imienia")
	flag.StringVar(&surnamestring, "surnamechar", "", "Pierwsza nazwiska")

	flag.Parse()

	if dzien == 0 {
		fmt.Print("Podaj dzień urodzenia: ")
		fmt.Scan(&dzien)
	}

	if namestring == "" {
		fmt.Print("Podaj imię: ")
		fmt.Scan(&namestring)
	}

	if surnamestring == "" {
		fmt.Print("Podaj nazwisko: ")
		fmt.Scan(&surnamestring)
	}

	fmt.Print("Dostaliśmy " + strconv.Itoa(dzien) + " " + namestring + " " + surnamestring + "\n")

	fmt.Print("Twoja ksywka to: ", callSignMaker(namestring[0:1], surnamestring[0:1], dzien))

}
