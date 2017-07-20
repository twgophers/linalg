package linalg

import (
	"math"

	"github.com/twgophers/collections"
)

//Add subtracts two collections.Vectors element wise
func Add(x, y collections.Vector) collections.Vector {
	pairs := collections.Zip(x, y)

	result := make(collections.Vector, len(pairs))

	for i, value := range pairs {
		result[i] = value.A + value.B
	}
	return result
}

//Subtract subtracts two collections.collections.Vectors elementwise
func Subtract(x, y collections.Vector) collections.Vector {
	pairs := collections.Zip(x, y)

	result := make(collections.Vector, len(pairs))

	for i, value := range pairs {
		result[i] = value.A - value.B
	}
	return result
}

func Dot(x, y collections.Vector) float64 {
	pairs := collections.Zip(x, y)
	var sum float64
	for _, tuple := range pairs {
		sum += tuple.A * tuple.B
	}
	return sum
}

func SumOfSquares(sample collections.Vector) float64 {
	return Dot(sample, sample)
}

func SquaredDistance(x, y collections.Vector) float64 {
	result := Subtract(x, y)
	return SumOfSquares(result)
}

func Distance(x, y collections.Vector) float64 {
	return math.Sqrt(SquaredDistance(x, y))
}
