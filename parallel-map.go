// Parallel Map
// ------------
// Parallel Map - The "Hello, World" of Erlang, implemented in Go.
// Is it doable? Yes, it is.

package main

import (
  "fmt"
)

type f func(int) int
type result struct {
  idx    int
  result int
}

// pmap applies `f` to each element in `slice` in a separate go routine.
func pmap(slice []int, fn f) []int {
  resultsChannel := make(chan result)

  for i, val := range slice {
    go func(idx int, val int) {
      resultsChannel <- result{idx, fn(val)}
    } (i, val)
  }

  results := make([]int, len(slice))
  
  for _ = range results {
    result := <- resultsChannel
    results[result.idx] = result.result
  }

  return results
}

func main() {
  // Some test data
  data := []int {2, 3, 5, 7, 11}
  
  // Apply some funtions to the test data
  fmt.Println("Verdoppelung")
  res := pmap(data, func(x int) int {
    return 2*x
  })
  
  for i, x := range data {
    fmt.Printf("%d -> %d\n", x, res[i])
  }

  fmt.Println("Quadrate")
  res = pmap(data, func(x int) int {
    return x*x
  })
  
  for i, x := range data {
    fmt.Printf("%d -> %d\n", x, res[i])
  }

  fmt.Println("Fakultät")
  res = pmap(data, func(x int) int {
    res := 1
    for i := 2; i <= x; i++ {
      res *= i
    }
    return res
  })
  
  for i, x := range data {
    fmt.Printf("%d -> %d\n", x, res[i])
  }
}