package main

import "rand"
import "time"
import "fmt"
import "flag"

var N int
var R int
func init() {
  flag.IntVar(&N, "size", 1000000, "Size of input arrays")
  flag.IntVar(&R, "runs", 10, "Number of times to repeat each test")
  flag.Parse()
}

func Bench(d []int) []float {
  times := make([]float, 5)

  t := NewTree(func(a,b int) bool { return a < b })

  start := time.Nanoseconds()
  for _,v := range d {
    t.Insert(v)
  }
  times[0] = float(time.Nanoseconds() - start) / 1000000000.0

  start = time.Nanoseconds()
  for _,v := range d {
    t.Insert(v)
  }
  times[1] = float(time.Nanoseconds() - start) / 1000000000.0

  start = time.Nanoseconds()
  for i := 0; i < len(d)/2; i++ {
    t.Remove(d[i])
  }
  times[2] = float(time.Nanoseconds() - start) / 1000000000.0
  
  start = time.Nanoseconds()
  for i := 0; i < len(d)/2; i++ {
    t.Remove(d[i])
  }
  times[3] = float(time.Nanoseconds() - start) / 1000000000.0

  start = time.Nanoseconds()
  for v := range d {
    t.Contains(v)
  }
  times[4] = float(time.Nanoseconds() - start) / 1000000000.0

  return times
}

func main() {
  names := []string {
    "Unique Inserts",
    "Repeated Inserts",
    "Unique Deletes",
    "Repeated Deletes",
    "Queries",
  }

  d := rand.Perm(N)
  total := Bench(d)
  for i := 1; i < R; i++ {
    times := Bench(d)
    for j := range times {
      total[j] += times[j]
    }
  }
  for i := range total {
    total[i] /= float(R)
  }

  fmt.Printf("Using input size %d and averaged over %d runs.\n", N, R)
  fmt.Printf("%3.3f:\t%d\t%s\n", total[0], N, names[0])
  fmt.Printf("%3.3f:\t%d\t%s\n", total[1], N, names[1])
  fmt.Printf("%3.3f:\t%d\t%s\n", total[2], N/2, names[2])
  fmt.Printf("%3.3f:\t%d\t%s\n", total[3], N/2, names[3])
  fmt.Printf("%3.3f:\t%d\t%s\n", total[4], N, names[4])
}
