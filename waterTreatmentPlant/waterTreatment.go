package waterTreatment

import "fmt"

func init() {
	fmt.Println("Water treatment plant initialized")
}

type WaterTreatmentPlant struct {
	Name string
}

func (w WaterTreatmentPlant) Data(name string) string {
	w.Name = "Water Treatment Plant: " + name
	return w.Name
}
