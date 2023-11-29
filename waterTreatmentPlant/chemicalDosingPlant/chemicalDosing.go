package chemicalDosing

import "fmt"

func init() {
	fmt.Println("Chemical Dosing Plant initialized")
}

type ChemicalDosingPlant struct {
	Name string
}

func (c ChemicalDosingPlant) Data(name string) string {
	c.Name = "Chemical Dosing : " + name
	return c.Name
}
