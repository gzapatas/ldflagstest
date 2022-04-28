package main

import (
	"encoding/json"
	"fmt"

	"github.com/gzapatas/ldflagstest/build"
)

var Version string = "deployment"

func Suma(a, b int) int {
	return a + b
}

func Resta(a, b int) int {
	return a - b
}

func Multiplicacion(a, b int) int {
	return a * b
}

type EventMeter struct {
	accounting  interface{}
	eventTypeId uint8
}

type Meters struct {
	Turnover  int64 `json:"turnover"`
	Totalwins int64 `json:"totalwins"`
}

func main() {
	var smibs []string = []string{
		"Smib1",
	}

	//253 -> RESET , 254 -> INITIAL, 255 -> FINAL, 252 -> ROLLOVER

	for range smibs {
		//Llamar al endpoint de events
		var events []EventMeter = []EventMeter{
			{
				accounting: map[string]interface{}{
					"turnover":  5000,
					"totalwins": 6000,
				},
				eventTypeId: uint8(254),
			},
			{
				accounting: map[string]interface{}{
					"turnover":  8000,
					"totalwins": 10000,
				},
				eventTypeId: uint8(255),
			},
		}

		var deltaTurnover int64 = 0
		var deltaTotalwins int64 = 0

		for _, event := range events {
			var meter Meters
			data, _ := json.Marshal(event.accounting)

			json.Unmarshal(data, &meter)

			if event.eventTypeId == 254 {
				deltaTurnover -= meter.Turnover
				deltaTotalwins -= meter.Totalwins
			}

			if event.eventTypeId == 255 {
				deltaTurnover += meter.Turnover
				deltaTotalwins += meter.Totalwins
			}
		}

		var production = deltaTotalwins - deltaTurnover
		//ACA DEBES GUARDAR EN BASE DE DATOS
		fmt.Printf("deltaTotalwins %v\n", deltaTotalwins)
		fmt.Printf("deltaTurnover %v\n", deltaTurnover)
		fmt.Printf("PRODUCCION %v\n", production)
	}

	fmt.Printf("Version = %v\n", Version)
	fmt.Printf("Time = %v\n", build.Time)
	fmt.Printf("User = %v\n", build.User)

	fmt.Printf("Suma = %v\n", Suma(2, 5))
	fmt.Printf("Resta = %v\n", Resta(2, 5))
	fmt.Printf("Multiplicacion = %v\n", Multiplicacion(2, 5))
}
