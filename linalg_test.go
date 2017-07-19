package linalg

import (
	"reflect"
	"testing"

	"github.com/twgophers/collections"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		x, y collections.Vector
		want collections.Vector
	}{
		{collections.Vector{1.0, 1.0}, collections.Vector{0.0, 0.0},
			collections.Vector{1.0, 1.0}},
		{collections.Vector{1.0, 2.0}, collections.Vector{2.0, 1.0}, collections.Vector{3.0, 3.0}},
		{collections.Vector{1.0, 2.0}, collections.Vector{2.0}, collections.Vector{3.0}},
		{collections.Vector{2.0}, collections.Vector{1.0, 2.0}, collections.Vector{3.0}},
		{collections.Vector{}, collections.Vector{}, collections.Vector{}},
		{collections.Vector{2.0}, collections.Vector{}, collections.Vector{}},
		{collections.Vector{}, collections.Vector{2.0}, collections.Vector{}},
	}
	for _, c := range cases {
		got := Add(c.x, c.y)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Add(%v, %v) want: %v; got: %v",
				c.x, c.y, c.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	cases := []struct {
		x, y collections.Vector

		want collections.Vector
	}{
		{
			collections.Vector{1.0, 1.0},
			collections.Vector{0.0, 0.0},
			collections.Vector{1.0, 1.0},
		},
		{
			collections.Vector{0.0, 1.0},
			collections.Vector{1.0, 2.0},
			collections.Vector{-1.0, -1.0},
		},
		{
			collections.Vector{1.0, 2.0},
			collections.Vector{1.0},
			collections.Vector{0.0},
		},
		{
			collections.Vector{1.0},
			collections.Vector{2.0, 3.0},
			collections.Vector{-1.0},
		},
		{
			collections.Vector{},
			collections.Vector{},
			collections.Vector{},
		},
		{
			collections.Vector{2.0},
			collections.Vector{},
			collections.Vector{},
		},
		{
			collections.Vector{},
			collections.Vector{2.0},
			collections.Vector{},
		},
	}
	for _, c := range cases {
		got := Subtract(c.x, c.y)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Subtract(%v, %v) want: %v; got: %v",
				c.x, c.y, c.want, got)
		}
	}
}

func TestDot(t *testing.T) {
	cases := []struct {
		x, y collections.Vector
		want float64
	}{
		{collections.Vector{1.0, 1.0}, collections.Vector{1.0, 1.0}, 2.0},
		{collections.Vector{1.0, 2.0, 3.0}, collections.Vector{4.0, 5.0}, 14.0},
		{collections.Vector{1.0}, collections.Vector{2.0, 5.0}, 2.0},
		{collections.Vector{}, collections.Vector{2.0, 5.0}, 0.0},
	}

	for _, c := range cases {
		gotDot := Dot(c.x, c.y)
		if gotDot != c.want {
			t.Errorf("%v.Dot(%v) want: %v but got: %v", c.x, c.y, c.want, gotDot)
		}
	}
}

func TestSumOfSquares(t *testing.T) {
	cases := []struct {
		collections.Vector
		want float64
	}{
		{collections.Vector{1.0, 1.0}, 2},
		{collections.Vector{1.0, 2.0, 3.0}, 14.0},
		{collections.Vector{0.0, 0.0, 0.0}, 0.0},
	}

	for _, c := range cases {
		gotSumOfSquares := SumOfSquares(c.Vector)
		if c.want != gotSumOfSquares {
			t.Errorf("SumOfSquares(%v) want: %v but got: %v", c.Vector, c.want,
				gotSumOfSquares)
		}
	}
}

// def squared_distance(v, w):
//     return sum_of_squares(collections.Vector_subtract(v, w))
func TestSquaredDistance(t *testing.T) {
	cases := []struct {
		x    collections.Vector
		y    collections.Vector
		want float64
	}{
		{collections.Vector{1, 1}, collections.Vector{2, 2}, 2},
	}

	for _, c := range cases {
		gotSquaredDistance := SquaredDistance(c.x, c.y)
		if c.want != gotSquaredDistance {
			t.Errorf("(%v).SquaredDistance(%v) want: %v but expected: %v",
				c.x, c.y, c.want, gotSquaredDistance)
		}
	}
}

func TestDistance(t *testing.T) {
	cases := []struct {
		x, y collections.Vector
		want float64
	}{
		{collections.Vector{1.0, 1.0}, collections.Vector{2.0, 2.0}, 1.4142135623730951},
		{collections.Vector{1.0}, collections.Vector{2.0, 2.0}, 1.0},
	}

	for _, c := range cases {
		gotDistance := Distance(c.x, c.y)
		if c.want != gotDistance {
			t.Errorf("(%v).Distance(%v) want: %v but got %v", c.x, c.y, c.want,
				gotDistance)
		}
	}
}
