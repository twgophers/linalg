package matrix

import "fmt"

//Row in a matrix
type Row []float64

//Col in a matrix
type Col []float64

//Matrix Represents a matrix
type Matrix []Row

type operator func(i, j int) float64

//Shape Return the number of lines and columns
func (matrix Matrix) Shape() (int, int) {
	matrix.validate()
	return len(matrix), len(matrix[0])
}

func (matrix Matrix) validate() {
	colSize := len(matrix[0])
	for _, row := range matrix {
		if colSize != len(row) {
			panic(fmt.Sprintf("The dimentions are differents: %d != %d",
				colSize, len(row)))
		}
	}
}

//RowAt return a row in positio i
func (matrix Matrix) RowAt(i int) Row {
	matrix.validate()
	if i > len(matrix) {
		panic(fmt.Sprintf("Does not exist row in position %d", i))
	}
	return matrix[i-1]
}

//ColumnAt column in the position
func (matrix Matrix) ColumnAt(col int) Col {
	matrix.validate()
	column := make(Col, len(matrix))
	for i, row := range matrix {
		if col > len(row) {
			panic(fmt.Sprintf("Does not exist column in position %d", col))
		}
		column[i] = row[col-1]
	}
	return column
}

func makeMatrix(op operator, numRows, numCols int) Matrix {
	matrix := make(Matrix, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make(Row, numCols)
		for j := 0; j < numCols; j++ {
			matrix[i][j] = op(i, j)
		}
	}
	return matrix
}

//MakeMatrixDiagonal create a diagonal Matrix
func MakeMatrixDiagonal(numRows, numCols int) Matrix {
	if numRows != numCols {
		panic(fmt.Sprintf("The dimentions are differents: %v, %v", numRows, numCols))
	}
	return makeMatrix(
		func(i, j int) float64 {
			if i == j {
				return 1.0
			}
			return 0.0
		}, numRows, numCols)
}
