package vector


func Add(v, w []float64) []float64 {

  if (len(v) != len(w)) {
    panic("The vectors have different sizes.")
  }

  result := make([]float64, len(v))

  for i, x := range v {
    result[i] = x + w[i]
  }

  return result
}

func Subtract(v, w []float64) []float64 {
  if (len(v) != len(w)) {
    panic("The vectors have different sizes.")
  }
  result := make([]float64, len(v))

  for i, x := range v {
    result[i] = x - w[i]
  }

  return result
}
