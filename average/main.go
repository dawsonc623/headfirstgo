package main

import (
	"fmt"
	"log"

	"cyledawson.net/datafile"
)

func main() {
	weights, err := datafile.GetFloats("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, weight := range weights {
		sum += weight
	}

	sampleCount := float64(len(weights))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
