package matrix

import "fmt"

//Row in a matrix
type Row []float64

//Matrix Represents a matrix
type Matrix []Row

//Shape Return the number of lines and columns
func (matrix Matrix) Shape() (int, int) {
	colSize := len(matrix[0])
	for _, row := range matrix {
		if colSize != len(row) {
			panic(fmt.Sprintf("The dimentions are differents: %d != %d",
				colSize, len(row)))
		}
	}
	return len(matrix), len(matrix[0])
}

//RowAt return a row in positio i
func (matrix Matrix) RowAt(i int) Row {
	if i > len(matrix) {
		panic(fmt.Sprintf("Does not exist row in position %d", i))
	}
	return matrix[i-1]
}

//ColumnAt column in the position
func (matrix Matrix) ColumnAt(col int) []float64 {

	column := make([]float64, len(matrix))
	for i, row := range matrix {
		if col > len(row) {
			panic(fmt.Sprintf("Does not exist column in position %d", col))
		}
		column[i] = row[col-1]
	}
	return column
}
