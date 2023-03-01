package main

import (
	"fmt"
)

type Vehicle struct {
	VehicleName string
	VehicleType string
}

func main() {

	vech := []Vehicle{
		{VehicleName: "Innova", VehicleType: "Four_wheel"},
		{VehicleName: "Activa", VehicleType: "Two_wheel"},
		{VehicleName: "splender", VehicleType: "Two_wheel"},
		{VehicleName: "swift", VehicleType: "Four_wheel"},
	}

	// Print the slice of vechile structs
	fmt.Println("vechiles:")
	for _, val := range vech {
		fmt.Println("  VehicleName:", val.VehicleName)
		fmt.Println("  VehicleName:", val.VehicleType)
	}
	fmt.Println()

	// creating map of struct
	VehicleMap := map[string]Vehicle{
		"Maruti_suzuki": {VehicleName: "ignis", VehicleType: "Four_wheel"},
		"hero":          {VehicleName: "shine", VehicleType: "Two_wheel"},
		"Royal_Enfield": {VehicleName: "Interceptor", VehicleType: "Two_wheel"},
	}

	// Print the map of Person structs
	fmt.Println("vechile_Map:")
	for _, p := range VehicleMap {
		fmt.Println("  vechile Name:", p.VehicleName)
		fmt.Println("  vechile Type:", p.VehicleType)
	}
}
