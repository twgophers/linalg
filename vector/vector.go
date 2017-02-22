package vector

func Add(v1, v2 []float64) []float64 {
  result := make([]float64, len(v1))

  for i, x := range v1 {
    result[i] = x + v2[i]
  }
  return result
}
