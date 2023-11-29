package main

import (
	"fmt"
	waterStation "powerPlant/waterTreatmentPlant"
	chemicalStation "powerPlant/waterTreatmentPlant/chemicalDosingPlant"
)

func main() {
	w := new(waterStation.WaterTreatmentPlant)
	fmt.Println(w.Data("Wartsila"))

	c := new(chemicalStation.ChemicalDosingPlant)
	fmt.Println(c.Data("Nalco"))
}
