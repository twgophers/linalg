package linalg

import (
	"math"

	"github.com/rpinheiroalmeida/collections"
)

//Add subtracts two collections.Vectors element wise
func Add(x, y collections.Vector) collections.Vector {
	tuples := collections.Zip(x, y)

	result := make(collections.Vector, len(tuples))

	for i, value := range tuples {
		result[i] = value.A + value.B
	}
	return result
}

//Subtract subtracts two collections.collections.Vectors elementwise
func Subtract(x, y collections.Vector) collections.Vector {
	tuples := collections.Zip(x, y)

	result := make(collections.Vector, len(tuples))

	for i, value := range tuples {
		result[i] = value.A - value.B
	}
	return result
}

func Dot(x, y collections.Vector) float64 {
	tuples := collections.Zip(x, y)
	var sum float64
	for _, tuple := range tuples {
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
