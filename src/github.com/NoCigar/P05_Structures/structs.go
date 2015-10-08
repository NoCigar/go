package main

import (
	"fmt"
)

func main() {

	unknownTank := Tank{}

	m1a1 := Tank{
		Name: "Abrams",
		Caliber: 120,
	}

	panther := Tank{ Name:"PzKpwV" }
	panther.Caliber = 88

	t34 := Tank {"T34", 72}

	fmt.Printf("unknown: %s %d", unknownTank.Name , unknownTank.Caliber)
	fmt.Printf("\nabrams: %s %d", m1a1.Name , m1a1.Caliber)
	fmt.Printf("\npanther: %s %d", panther.Name , panther.Caliber)
	fmt.Printf("\nt34: %s %d", t34.Name , t34.Caliber)

	fmt.Println()

	upgradeCaliberNoPointer(t34)
	fmt.Printf("\nt34 after noPointer upgrade: %s %d", t34.Name , t34.Caliber)
	upgradeCaliber(&t34)
	fmt.Printf("\nt34 after pointer upgrade: %s %d", t34.Name , t34.Caliber)

	fmt.Println()

	panther.upgradeCaliberMethod()
	fmt.Printf("\npanther after upgradeCaliber method: %s %d", panther.Name , panther.Caliber)

	emptyTank := Tank{Name: "empty"}
	fmt.Printf("\nemptyTank pre upgradeCaliber method: %s %d", emptyTank.Name , emptyTank.Caliber)
	emptyTank.upgradeCaliberMethod()
	fmt.Printf("\nemptyTank after upgradeCaliber method: %s %d", emptyTank.Name , emptyTank.Caliber)


}


type Tank struct {
	Name string
	Caliber int
}

func upgradeCaliberNoPointer (tank Tank) {
	tank.Name+="-upgraded"
	tank.Caliber+=10
}

func upgradeCaliber (tank *Tank){
	tank.Name+="-upgraded"
	tank.Caliber+=10
}

func (tank *Tank) upgradeCaliberMethod() {
	tank.Name+="-upgradedX2"
	tank.Caliber+=20
}