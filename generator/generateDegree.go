package generator

import (
	"math"
	"math/rand"
	"time"
)

func generateNaturalPatternStream(min, max float64) <-chan float64 {
	amplitude := (max - min) / 2
	offset := (max + min) / 2
	currentValue := offset
	velocity := 0.0

	ch := make(chan float64)

	go func() {
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			randomFactor := (rand.Float64() - 0.5) * amplitude * 0.5

			velocity += randomFactor
			velocity *= 0.1

			currentValue += velocity
			currentValue = math.Max(min, math.Min(max, currentValue))

			ch <- currentValue
		}
	}()

	return ch
}

// func main() {
// 	stream := generateNaturalPatternStream(0, 100)

// 	// Example usage: print the first 20 values
// 	for i := 0; i < 20; i++ {
// 		value := <-stream
// 		println(value)
// 	}
// }