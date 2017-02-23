package vector

import(
  "testing"
  "reflect"
)

func TestAdd(t *testing.T) {
  cases := []struct {
      v1        []float64
      v2        []float64
      want      []float64
    }{
      {[]float64{1.0, 1.0}, []float64{0.0, 0.0}, []float64{1.0, 1.0}},
      {[]float64{1.0, 2.0}, []float64{2.0, 1.0}, []float64{3.0, 3.0}},
      {[]float64{}, []float64{}, []float64{}},
    }
    for _, c := range cases {
      got := Add(c.v1, c.v2)
      if !reflect.DeepEqual(got, c.want) {
        t.Errorf("Add(%v, %v) want: %v; got: %v",
          c.v1, c.v2, c.want, got)
      }
    }
}

func TestAdd_WhenLenVectorsAreDifferent(t *testing.T) {
  v1 := []float64{1.0,2.0}
  v2 := []float64{1.0}

  defer func() {
        rec := recover()
        if rec == nil {
            t.Errorf("The vectors (%v, %v) have different sizes.", v1, v2)
        }
  }()

  Add(v1, v2)
}

func TestSubtract(t *testing.T) {
  cases := []struct {
      v1        []float64
      v2        []float64
      want      []float64
    }{
      {[]float64{1.0, 1.0}, []float64{0.0, 0.0}, []float64{1.0, 1.0}},
      {[]float64{0.0, 1.0}, []float64{1.0, 2.0}, []float64{-1.0, -1.0}},
    }
    for _, c := range cases {
      got := Subtract(c.v1, c.v2)
      if !reflect.DeepEqual(got, c.want) {
        t.Errorf("Add(%v, %v) want: %v; got: %v",
          c.v1, c.v2, c.want, got)
      }
    }
}

func TestSubtract_WhenLenVectorAreDifferent(t *testing.T) {
  v := []float64{1.0, 2.0}
  w := []float64{1.0}

  defer func() {
    if (recover() == nil) {
      t.Errorf("The vector (%v, %v) have different sizes.", v, w)
    }
  }()

  Subtract(v, w)
}
