package matrix

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestShape(t *testing.T) {
	cases := []struct {
		matrix             Matrix
		wantRows, wantCols int
	}{
		{
			Matrix{Row{1.0, 1.0, 1.0}, Row{0.0, 2.0, 1.0}, Row{-1.0, 0.0, 1.0}},
			3,
			3,
		},
		{
			Matrix{Row{1.0, 1.0}, Row{0.0, 2.0}, Row{-1.0, 0.0}},
			3,
			2,
		},
	}
	for _, c := range cases {
		gotRows, gotCols := c.matrix.Shape()
		if gotRows != c.wantRows {
			t.Errorf("Shape(%v) want_rows: %v; got_rows: %v",
				c.matrix, c.wantRows, gotRows)
		}
		if gotCols != c.wantCols {
			t.Errorf("Shape(%v) want_cols: %v; got_cols: %v",
				c.matrix, c.wantCols, gotCols)
		}
	}
}

func TestShape_panicWhenTheSizeOfRowsAreDifferents(t *testing.T) {
	matrix := Matrix{Row{1.0, 2.0, 3.0}, Row{1.0, 2.0, 3.0}, Row{1.0, 2.0}}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Shape(%v) did not panic.", matrix)
		} else if !strings.HasPrefix(fmt.Sprint(r), "The dimentions are differents:") {
			t.Errorf("Shape(%v): unexpected panic %q", matrix, r)
		}
	}()
	matrix.Shape()
}

func TestRowAt(t *testing.T) {
	cases := []struct {
		matrix   Matrix
		position int
		want     Row
	}{
		{
			Matrix{
				Row{1.0, 1.0, 1.0},
				Row{0.0, 2.0, 1.0},
				Row{-1.0, 0.0, 1.0},
			},
			2,
			Row{0.0, 2.0, 1.0},
		},
		{
			Matrix{Row{1.0, 2.0, 3.0}},
			1,
			Row{1.0, 2.0, 3.0},
		},
	}
	for _, c := range cases {
		got := c.matrix.RowAt(c.position)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("%v.RowAt(%v) want: %v; got: %v", c.matrix, c.position, c.want,
				got)
		}
	}
}

func TestRowAt_panicWhenTheRowDoesNotExists(t *testing.T) {
	matrix := Matrix{Row{1.0, 2.0, 3.0}}
	position := 2
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%v.RowAt(%v) dit not panic", matrix, position)
		} else if !strings.HasPrefix(fmt.Sprint(r), "Does not exist row in position") {
			t.Errorf("%v.RowAt(%v): unexpected panic %q", matrix, position, r)
		}
	}()
	matrix.RowAt(position)
}

func TestRow_panicWhenTheMatrixIsWrong(t *testing.T) {
	matrix := Matrix{Row{1.0}, Row{1.0, 1.0}}
	column := 0
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%v.RowAt(%v) did not panic.", matrix, column)
		} else if !strings.HasPrefix(fmt.Sprint(r), "The dimentions are differents:") {
			t.Errorf("%v.RowAt(%v): unexpected panic %q", matrix, column, r)
		}
	}()
	matrix.RowAt(column)
}

func TestColumnAt(t *testing.T) {
	cases := []struct {
		matrix Matrix
		col    int
		want   Col
	}{
		{Matrix{Row{1.0, 1.0}, {0.0, 0.0}}, 2, Col{1.0, 0.0}},
		{Matrix{Row{1.0, 1.0}, {0.0, 0.0}, {3.0, 2.0}}, 1, Col{1.0, 0.0, 3.0}},
	}
	for _, c := range cases {
		got := c.matrix.ColumnAt(c.col)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("%v.ColumnAt(%v) want: %v; got: %v", c.matrix, c.col, c.want, got)
		}
	}
}

func TestColumnAt_panicWhenTheColumnDoesNotExist(t *testing.T) {
	matrix := Matrix{Row{1.0, 1.0}}
	column := 3
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%v.ColumnAt(%v) did not panic.", matrix, column)
		} else if !strings.HasPrefix(fmt.Sprint(r), "Does not exist column in position") {
			t.Errorf("%v.ColumnAt(%v): unexpected panic %q", matrix, column, r)
		}
	}()
	matrix.ColumnAt(column)
}

func TestColumn_panicWhenTheMatrixIsWrong(t *testing.T) {
	matrix := Matrix{Row{1.0}, Row{1.0, 1.0}}
	column := 0
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%v.ColumnAt(%v) did not panic.", matrix, column)
		} else if !strings.HasPrefix(fmt.Sprint(r), "The dimentions are differents:") {
			t.Errorf("%v.ColumnAt(%v): unexpected panic %q", matrix, column, r)
		}
	}()
	matrix.ColumnAt(column)
}

func TestMakeMatrixDiagonal(t *testing.T) {
	cases := []struct {
		numRows int
		numCols int
		want    Matrix
	}{
		{2, 2, Matrix{Row{1.0, 0.0}, Row{0.0, 1.0}}},
		{3, 3, Matrix{Row{1.0, 0.0, 0.0}, Row{0.0, 1.0, 0.0}, Row{0.0, 0.0, 1.0}}},
	}
	for _, c := range cases {
		got := MakeMatrixDiagonal(c.numRows, c.numCols)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MakeMatrixDiagonal(%v, %v) want: %v; got: %v",
				c.numRows, c.numCols, c.want, got)
		}
	}
}

func TestMakeMatrixDiagonal_panicWhenNumRowAndNumColsAreDifferents(t *testing.T) {
	numRows, numCols := 2, 3
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MakeMatrixDiagonal(%v, %v) did not panic.", numRows, numCols)
		} else if !strings.HasPrefix(fmt.Sprint(r), "The dimentions are differents:") {
			t.Errorf("MakeMatrixDiagonal(%v, %v): unexpected panic %q", numRows, numCols, r)
		}
	}()
	MakeMatrixDiagonal(numRows, numCols)
}
