package main

import(
	"fmt"
	"strings"
)

func main() {

	plants := []PowerPlant{
		PowerPlant{hydro,30,active},
		PowerPlant{hydro,30,inactive},
		PowerPlant{solar,20,active},
		PowerPlant{wind,60,unavailable},
	}

	grid := PowerGrid{300, plants}

	activePlants := []int{0 , 1}

	gridLoad := 75.

	//selectReport(plantCapacities , activePlants, gridLoad)

}



func generatePlantCapacityReport(plantCapacities []float64) {
	for idx, cap := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
	}
}

func generatePlantPowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64){
	capacity := 0.
	for _,plantId := range activePlants {
		capacity += plantCapacities[plantId]
	}

	fmt.Printf("%-20s%.0f\n","Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load: ", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utitlization: ", gridLoad/capacity*100)
}

func selectReport(plantCapacities []float64, activePlants []int, gridLoad float64) {
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Println("Please choose an option:")

	var option string

	fmt.Scanln(&option)

	switch option {
	case "1":
		generatePlantCapacityReport(plantCapacities)
	case "2":
		generatePlantPowerGridReport(activePlants, plantCapacities, gridLoad)
	default:
		fmt.Println("Unknowne option, exiting")

	}

}

type PlantType string

const (
	hydro PlantType = "hydro"
	wind PlantType = "wind"
	solar PlantType = "solar"
)

type PlantStatus string

const (
	active PlantStatus = "Active"
	inactive PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	plantType PlantType
	capacity float64
	status PlantStatus
}

type PowerGrid struct {
	load float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
	for idx, p := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #",idx)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-",len(label)))
		fmt.Printf("%-20s%s\n","Type: ", p.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity: ", p.capacity)
		fmt.Printf("%-20s%s\n", "Status: ", p.status)
		fmt.Println("")
	}
}