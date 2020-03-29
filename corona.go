package main

import (
	"fmt"
)

func main() {
	fmt.Println("\nDate: 2020/03/29")
	fmt.Println("Time: 15:45")
	fmt.Println("\nPlease select an option:")
	fmt.Println("1. Print Covid19 cases in Pakistan")
	fmt.Println("2. Print Covid19 cases in South Korea")
	fmt.Println("3. Print Covid19 cases in France")
	fmt.Println("4. Print my personalized message to CoronaVirus")
	var checkPK, checkSK, checkFR, check4 bool = false, false, false, false
	coronaPK, coronaSK, coronaFR := 1526, 9583, 37575
	var choice int
	for {
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			checkPK = true
			fmt.Print("Corona cases in Pakistan: ", coronaPK, "\n\n")
		case 2:
			checkSK = true
			fmt.Print("Corona cases in South Korea: ", coronaSK, "\n\n")
		case 3:
			checkFR = true
			fmt.Print("Corona cases in France: ", coronaFR, "\n\n")
		case 4:
			check4 = true
			fmt.Print("oof\n\n")
		default:
			fmt.Print("Error in input. Try again\n\n")
		}
		if check4 && checkFR && checkPK && checkSK {
			break
		}
	}

}
