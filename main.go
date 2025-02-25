package main

import "fmt"

func main() {
	fmt.Println("Tervetuloa laskimeeni!")
	fmt.Println("Valitse toiminto:")
	fmt.Println("1. Summa")
	fmt.Println("2. Erotus")
	fmt.Println("3. Kerto")
	fmt.Println("4. Osamäärä")
	fmt.Println("5. Lopetus")
	var toiminto int

	fmt.Scan(&toiminto)

	var luku1 float64
	var luku2 float64

	fmt.Print("Syötä ensimmäinen luku: ")
	fmt.Scan(&luku1)

	fmt.Print("Syötä toinen luku: ")
	fmt.Scan(&luku2)

	if toiminto == 1 {
		fmt.Println("Summa: ", luku1+luku2)
	} else if toiminto == 2 {
		fmt.Println("Erotus: ", luku1-luku2)
	} else if toiminto == 3 {
		fmt.Println("Kerto: ", luku1*luku2)
	} else if toiminto == 4 {
		fmt.Println("Osamäärä: ", luku1/luku2)
	}

}
