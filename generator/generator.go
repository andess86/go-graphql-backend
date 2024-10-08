package generator

import (
	"log"
	"math/rand"
	"time"
)

// Propeller represents the propeller data structure
type Propeller struct {
	ID     string
	Pitch  float64
	RPM    int
	Degree float64
}

// PropellerGenerator handles the generation of propeller data.
type PropellerGenerator struct{}

// NewPropellerGenerator creates a new instance of the PropellerGenerator
func NewPropellerGenerator() *PropellerGenerator {
	return &PropellerGenerator{}
}

// StartGenerating returns a channel that emits propeller data periodically.
func (pg *PropellerGenerator) StartGenerating() <-chan Propeller {
	log.Println("Just inside StartGenerating()")
    propellerChannel := make(chan Propeller)


    go func() {
        log.Println("Starting to send mock propeller data...")
        DegreeStream := generateNaturalPatternStream(0, 100)
        defer close(propellerChannel)
        for {
            
            propellerData := Propeller{
                ID:     "1",                   // Static ID, but you can change this to be dynamic if needed
                Pitch:  rand.Float64() * 10,   // Random pitch value between 0 and 10
                RPM:    rand.Intn(5000),       // Random RPM value between 0 and 5000
                Degree: <-DegreeStream,        // Random degree between 0 and 359
            }
            propellerChannel <- propellerData
            log.Println("Generator: Sent new propeller data:", propellerData) // Log each send

            time.Sleep(100 * time.Millisecond) // Simulate data update every 500ms
        }
    }()

    return propellerChannel
}

