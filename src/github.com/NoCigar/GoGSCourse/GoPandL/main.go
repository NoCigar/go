package main

import(
	"fmt"
)

func main() {
	var plantCapacities []float64

	plantCapacities = []float64{30, 30, 30, 60, 60, 100}

	activePlants := []int{0 , 1}

	gridLoad := 75.

	selectReport(plantCapacities , activePlants, gridLoad)

	/*fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Println("Please choose an option:")

	var option string

	fmt.Scanln(&option)

	var option string

	option = selectReport()

	switch option {
	case "1":
		generatePlantCapacityReport(plantCapacities...)
	case "2":
		generatePlantPowerGridReport(activePlants, plantCapacities, gridLoad)
	default:
		fmt.Println("Unknowne option, exiting")

	}
*/
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