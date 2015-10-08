package main

import (
	"fmt"
)

func main(){
	var power int
	power = 9000
	fmt.Printf("muh REGULAR power lvl is over %d",power)

	inferedPower := 8000
	fmt.Printf("\nmuh INFERED power lvl is over %d", inferedPower)

	functionInferedPower := getPowerLvl()
	fmt.Printf("\nmuh FUNCTION INFERED power lvl is over %d", functionInferedPower)

	inferedPower, nameVar := 8001, "A silly name"
	fmt.Printf("\nthe new INFERED power lvl is %d and has a %s",inferedPower,nameVar)

}

func getPowerLvl() int {
	return 7000
}